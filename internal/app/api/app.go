package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/app/api/note_v1"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type App struct {
	noteImpl        *note_v1.Note
	serviceProvider *serviceProvider

	pathConfig string

	grpcServer *grpc.Server
	mux        *runtime.ServeMux
}

func NewApp(ctx context.Context, pathConfig string) (*App, error) {
	a := &App{
		pathConfig: pathConfig,
	}

	err := a.initDeps(ctx)

	return a, err
}

func (a *App) Run() error {
	defer func() {
		a.serviceProvider.db.Close()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := a.runGRPC()
		if err != nil {
			log.Fatalf("Failed to process gRPC server: %s", err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		err := a.runPublicHTTP()
		if err != nil {
			log.Fatalf("Failed to process muxer: %s", err.Error())
		}
	}()

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServiceProvider,
		a.initServer,
		a.initGRPCServer,
		a.initPublicHTTPHandlers,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider(a.pathConfig)
	return nil
}

func (a *App) initServer(ctx context.Context) error {
	a.noteImpl = note_v1.NewNote(a.serviceProvider.GetNoteService(ctx))

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)

	desc.RegisterNoteV1Server(a.grpcServer, a.noteImpl)

	return nil
}

func (a *App) initPublicHTTPHandlers(ctx context.Context) error {
	a.mux = runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, a.mux, a.serviceProvider.GetConfig().GRPC.GetAddress(), opts)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) runGRPC() error {
	list, err := net.Listen("tcp", a.serviceProvider.GetConfig().GRPC.GetAddress())

	if err != nil {
		log.Printf("Failed to listen TCP on port %s", a.serviceProvider.GetConfig().GRPC.GetAddress())
		return err
	}

	log.Printf("GRPC Server is listening on host: %s", a.serviceProvider.GetConfig().GRPC.GetAddress())

	if err = a.grpcServer.Serve(list); err != nil {
		return err
	}

	return nil
}

func (a *App) runPublicHTTP() error {
	log.Printf("HTTP Server is listening on host: %s ", a.serviceProvider.GetConfig().HTTP.GetAddress())

	if err := http.ListenAndServe(a.serviceProvider.GetConfig().HTTP.GetAddress(), a.mux); err != nil {
		return err
	}

	return nil
}
