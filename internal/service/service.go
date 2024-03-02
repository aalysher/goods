package service

import (
	"database/sql"
	"fmt"

	"github.com/aalysher/goods/config"
)

type Service struct {
	cfg Config
	db  *sql.DB
}

func New() (*Service, error) {
	s := &Service{}

	if err := s.init(); err != nil {
		return nil, fmt.Errorf("init service %w", err)
	}

	return s, nil
}

func (s *Service) init() (err error) {
	s.cfg, err = config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if err = s.initDB(); err != nil {
		return fmt.Errorf("init db: %w", err)
	}

	return nil
}

func (s *Service) initDB() (err error) {
	s.db, err = NewPostgres(s.cfg.Postgres)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}

	return nil
}
