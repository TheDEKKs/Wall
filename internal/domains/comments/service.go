package comments


import (
	"context"

	"thedekk/WWT/internal/domains/users/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type CommentsService struct {
	db   DBTX
	repo *repository.Queries
}

func NewCommentService(db DBTX) *CommentsService {
	repo := repository.New(db)

	return &CommentsService{
		db:   db,
		repo: repo,
	}
}