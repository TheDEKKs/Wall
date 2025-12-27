package comments

import (
	"context"

	"thedekk/WWT/internal/domains/users"
	"thedekk/WWT/internal/domains/comments/repository"
	"thedekk/WWT/internal/transport/middlewares"

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
	userService *users.UserService

}

func NewCommentService(db DBTX, userService *users.UserService) *CommentsService {
	repo := repository.New(db)

	return &CommentsService{
		db:   db,
		repo: repo,
		userService: userService,
	}
}

func (s *CommentsService) NewComment(ctx context.Context, wallName, text string) error {
	cookie := ctx.Value("cookie").(middlewares.CookieCtx)

	wallID, err := s.userService.GetWallIDByUserName(ctx, wallName)
	if err != nil {
		return err
	}

	if err := s.repo.NewComment(ctx, repository.NewCommentParams{
		UserID: cookie.UserID,
		WallID: *wallID,
		Text: text,
	}); err != nil {
		return err
	}


	return nil
}