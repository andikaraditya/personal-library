package user

import (
	"context"
	"errors"
	"time"

	"github.com/andikaraditya/personal-library/internal/api"
	"github.com/andikaraditya/personal-library/internal/db"
	"github.com/andikaraditya/personal-library/internal/helper"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserService interface {
	CreateUser(user User) error
	Login(user User) (string, error)
}

type srv struct {
	db db.DBService
}

var (
	Service UserService
)

func init() {
	Service = New(db.Service)
}

func New(db db.DBService) UserService {
	return &srv{db}
}

func (s *srv) CreateUser(user User) error {
	var err error
	if user.ID == "" {
		user.ID = uuid.NewString()
	}

	user.Password, err = helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	if err := s.db.Commit(nil, func(tx pgx.Tx) error {
		_, err := tx.Exec(
			context.Background(),
			`INSERT INTO "user" (
				id,
				name,
				email,
				password
			) VALUES ($1, $2, $3, $4)`,
			user.ID,
			user.Name,
			user.Email,
			user.Password,
		)

		if err != nil {
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) {
				if pgErr.Code == "42P05" {
					return api.ErrPayload
				}
			}
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *srv) Login(user User) (string, error) {
	var hash string
	var userId string

	if err := s.db.QueryRow(
		`SELECT password, id
			FROM "user"
			WHERE email = $1`,
		user.Email,
	).Scan(&hash, &userId); err != nil {
		return "", err
	}

	err := helper.ComparePassword(hash, user.Password)
	if err != nil {
		return "", api.ErrPayload
	}

	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return t, nil
}
