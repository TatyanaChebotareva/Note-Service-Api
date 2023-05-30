package app

import (
	"context"
	"log"

	"github.com/TatyanaChebotareva/Note-Service-Api/internal/config"
	"github.com/TatyanaChebotareva/Note-Service-Api/internal/pkg/db"
	noteRepo "github.com/TatyanaChebotareva/Note-Service-Api/internal/repository/note"
	note "github.com/TatyanaChebotareva/Note-Service-Api/internal/service"
)

type serviceProvider struct {
	db         db.Client
	configPath string
	config     *config.Config

	// repositories
	noteRepository noteRepo.Repository

	// services
	noteService *note.Service
}

func newServiceProvider(configPath string) *serviceProvider {
	return &serviceProvider{
		configPath: configPath,
	}
}

func (s *serviceProvider) GetConfig() *config.Config {
	if s.config == nil {
		cfg, err := config.NewConfig(s.configPath)
		if err != nil {
			log.Printf("failed to get config: %s", err.Error())
		}
		s.config = cfg
	}

	return s.config
}

func (s *serviceProvider) GetDB(ctx context.Context) db.Client {
	if s.db == nil {
		cfg, err := s.GetConfig().GetDBConfig()
		if err != nil {
			log.Fatalf("failed to get db config: %s", err.Error())
		}

		dbc, err := db.NewClient(ctx, cfg)
		if err != nil {
			log.Fatalf("can't connect to db err: %s", err.Error())
		}
		s.db = dbc
	}

	return s.db
}

func (s *serviceProvider) GetNoteRepository(ctx context.Context) noteRepo.Repository {
	if s.noteRepository == nil {
		s.noteRepository = noteRepo.NewNoteRepository(s.GetDB(ctx))
	}
	return s.noteRepository
}

func (s *serviceProvider) GetNoteService(ctx context.Context) *note.Service {
	if s.noteService == nil {
		s.noteService = note.NewService(s.GetNoteRepository(ctx))
	}

	return s.noteService
}
