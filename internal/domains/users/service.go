package users

import (
	"context"
	"errors"

	"thedekk/WWT/internal/domains/users/repository"
	"thedekk/WWT/internal/domains/walls"
	"thedekk/WWT/internal/security"

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

type UserService struct {
	db          DBTX
	repo        *repository.Queries
	wallService *walls.WallService
}

func NewUserService(db DBTX, wallService *walls.WallService) *UserService {
	repo := repository.New(db)

	return &UserService{
		db:          db,
		repo:        repo,
		wallService: wallService,
	}
}

func IsUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505"
	}
	return false
}

var (
	NonUnique    = errors.New("Non-unique")
	DoesNotExist = errors.New("Account does not exist")
)

func (s *UserService) RegistrationUser(ctx context.Context, userName, password string) (*map[string]string, error) {
	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	repoTx := s.repo.WithTx(tx)

	hash, err := security.HashPasswordCreate(password)
	if err != nil {
		return nil, err
	}

	user, err := repoTx.RegistrationUser(ctx, repository.RegistrationUserParams{
		UserName:     userName,
		PasswordHash: hash,
	})
	if err != nil {
		if IsUniqueViolation(err) {
			return nil, NonUnique
		}

		return nil, err
	}

	token, err := security.JwtCreate(hash, user.ID)
	if err != nil {
		return nil, err
	}

	wallTx := s.wallService.WithTx(tx)

	if err := wallTx.NewWall(ctx, user.ID); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	userCookie := map[string]string{
		"Token":  token,
		"UserID": user.ID.String(),
	}

	return &userCookie, nil
}

func (s *UserService) LoginUser(ctx context.Context, userName, password string) (*map[string]string, error) {
	user, err := s.repo.LoginUser(ctx, userName)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, DoesNotExist
	}

	if err != nil {
		return nil, err
	}

	if err := security.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, err
	}

	token, err := security.JwtCreate(user.PasswordHash, user.ID)
	if err != nil {
		return nil, err
	}

	userCookie := map[string]string{
		"Token":  token,
		"UserID": user.ID.String(),
	}

	return &userCookie, nil
}

func (s *UserService) GetWallIDByUserName(ctx context.Context, userName string) (*uuid.UUID, error) {
	userID, err := s.repo.GetUserIDByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}

	wallID, err := s.wallService.GetWallByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &wallID, nil
}

func (s *UserService) GetUserByUserID(ctx context.Context, userID uuid.UUID) (repository.User, error) {
	return s.repo.GetUserByUserID(ctx, userID)
}
