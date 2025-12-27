package walls

import (
	"context"

	"thedekk/WWT/internal/domains/walls/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type WallService struct {
	db   DBTX
	repo *repository.Queries
}

func NewWallService(db DBTX) *WallService {
	repo := repository.New(db)

	return &WallService{
		db:   db,
		repo: repo,
	}
}

func (s *WallService) WithTx(db pgx.Tx) *WallService {
	return NewWallService(db)
}

func (s *WallService) NewWall(ctx context.Context, userID uuid.UUID) (error) {
	if err := s.repo.SetWall(ctx, userID); err != nil {
		return err
	}

	return nil
} 

func (s *WallService) GetWallByID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	return s.repo.GetWallIDByUserID(ctx, userID)
}