package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/app/api/note_v1"
	noteRepo "github.com/TatyanaChebotareva/Note-Service-Api/internal/repository/note"
	note "github.com/TatyanaChebotareva/Note-Service-Api/internal/service"
	desc "github.com/TatyanaChebotareva/Note-Service-Api/pkg/note_v1"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

const (
	hostGrpc = ":50051"
	hostHttp = ":8090"
)

const (
	host       = "localhost"
	port       = "54321"
	dbUser     = "note-service-user"
	dbPassword = "note-service-password"
	dbName     = "note-service"
	sslMode    = "disable"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		startGrpc()
	}(&wg)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func(wg *sync.WaitGroup, ctx context.Context) {
		defer wg.Done()
		startHttp(ctx)
	}(&wg, ctx)

	wg.Wait()
}

func startGrpc() error {
	list, err := net.Listen("tcp", hostGrpc)
	if err != nil {
		log.Fatalf("failed to mappint port: %s", err.Error())
	}

	dbDsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, dbUser, dbPassword, dbName, sslMode,
	)

	db, err := sqlx.Open("pgx", dbDsn)
	if err != nil {
		return err
	}
	defer db.Close()

	noteRepository := noteRepo.NewNoteRepository(db)
	noteService := note.NewService(noteRepository)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpcValidator.UnaryServerInterceptor()),
	)
	desc.RegisterNoteV1Server(s, note_v1.NewNote(noteService))

	fmt.Println("GRPC Server is listening on port ", hostGrpc)

	if err = s.Serve(list); err != nil {
		log.Printf("failed to serve: %s", err.Error())
		return err
	}

	return nil
}

func startHttp(ctx context.Context) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := desc.RegisterNoteV1HandlerFromEndpoint(ctx, mux, hostGrpc, opts)
	if err != nil {
		return err
	}

	fmt.Println("HTTP Server is listening on port", hostHttp)

	return http.ListenAndServe(hostHttp, mux)
}
