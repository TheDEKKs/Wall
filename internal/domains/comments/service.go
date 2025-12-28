package comments

import (
	"context"
	"time"

	"thedekk/WWT/internal/domains/comments/repository"
	"thedekk/WWT/internal/domains/users"
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
	db          DBTX
	repo        *repository.Queries
	userService *users.UserService
}

func NewCommentService(db DBTX, userService *users.UserService) *CommentsService {
	repo := repository.New(db)

	return &CommentsService{
		db:          db,
		repo:        repo,
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
		Text:   text,
	}); err != nil {
		return err
	}

	return nil
}

type Comments struct {
	UserName string    `json:"user_name"`
	Comment  string    `json:"comment"`
	Time     time.Time `json:"time"`
}

func (s *CommentsService) GetCommentsWall(ctx context.Context, wallName string) (*[]Comments, error) {
	wallID, err := s.userService.GetWallIDByUserName(ctx, wallName)
	if err != nil {
		return nil, err
	}

	comments, err := s.repo.GetCommentsByWallID(ctx, *wallID)
	if err != nil {
		return nil, err
	}

	commentsOut := []Comments{}

	for _, c := range comments {
		user, err := s.userService.GetUserByUserID(ctx, c.UserID)
		if err != nil {
			return nil, err
		}

		commentsOut = append(commentsOut, Comments{UserName: user.UserName, Comment: c.Text, Time: *c.CreatedAt})

	}

	return &commentsOut, nil
}
