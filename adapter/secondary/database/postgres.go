package database

import (
	"context"
	"errors"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

const pathMigration = "../script/migrations"

type Repository interface {
	ListClient(ctx context.Context, page, limit int) (*[]entity.Client, *entity.Pagination, error)
	GetClient(ctx context.Context, clientID uuid.UUID) (clients entity.Client, err error)
	DeleteClient(ctx context.Context, clientId uuid.UUID) error

	Close()
}

type repository struct {
	db *sqlx.DB
}

func NewReaderConnection(URI string) (Repository, error) {
	db, err := sqlx.Connect("postgres", URI)
	if err != nil {
		return &repository{}, errors.New("failed to connect to read database")
	}

	err = runMigrations(db, pathMigration)
	if err != nil {
		return &repository{}, errors.New("failed to run migrations")
	}

	return &repository{db}, nil
}

func runMigrations(db *sqlx.DB, pathMigration string) error {
	err := goose.Up(db.DB, pathMigration)
	if err != nil {
		logger.WithError(err).Error("failed to run migrations")
		return err
	}

	return nil
}

func DownMigrations(db *sqlx.DB, pathMigration string) error {
	err := goose.Down(db.DB, pathMigration)
	if err != nil {
		logger.WithError(err).Error("failed to run migrations")
		return err
	}
	return err
}

func (r *repository) Close() {
	err := r.db.Close()
	if err != nil {
		logger.WithError(err).Error("error closing database connection")
	}
}
