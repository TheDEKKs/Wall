package bot

import (
	"context"

	"thedekk/WWT/internal/domains/telegram/repository"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DBTX interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

type TelegramService struct {
	db   DBTX
	repo *repository.Queries
}

func NewTelegramService(db DBTX) *TelegramService {
	repo := repository.New(db)

	return &TelegramService{
		db:   db,
		repo: repo,
	}
}

func (s *TelegramService) GetOrCreateUserCode(ctx context.Context, userName, firstName, lastName string, telegramID int64) (*repository.Telegram, error) {
	if user, err := s.repo.GetUserByTelegramID(ctx, telegramID); err == nil {
		return &user, nil
	}

	user, err := s.repo.NewTelegramRecord(ctx, repository.NewTelegramRecordParams{
		TelegramID: telegramID,
		Username:   userName,
		FirstName:  firstName,
		LastName:   lastName,
	})
	if err != nil {
		return nil, err
	}

	return &user, nil
}
