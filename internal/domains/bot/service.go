package bot

import (
	"context"

	"thedekk/WWT/internal/domains/bot/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type BotService struct {
	db   DBTX
	repo *repository.Queries
}

func NewBotService(db DBTX) *BotService {
	repo := repository.New(db)

	return &BotService{
		db:   db,
		repo: repo,
	}
}
