package repository

import (
	"github.com/gocql/gocql"
	"github.com/pooladkhay/tinier-auth-service/domain"
	"github.com/pooladkhay/tinier-auth-service/helper/errs"
)

type User interface {
	Create(*domain.User) *errs.Err
	GetByEmail(email string) (*domain.User, *errs.Err)
	Delete(email string) *errs.Err
}

type user struct {
	client *gocql.Session
}

func NewUser(s *gocql.Session) User {
	return &user{client: s}
}

func (c *user) Create(u *domain.User) *errs.Err {
	q := "INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?, ?);"
	err := c.client.Query(
		q,
		u.Name,
		u.Email,
		u.Password,
		u.CreatedAt,
	).Exec()
	if err != nil {
		return errs.NewBadRequestError(err.Error())
	}
	return nil
}

func (c *user) GetByEmail(email string) (*domain.User, *errs.Err) {
	var result domain.User
	var q = "SELECT name, email, password, created_at FROM users WHERE email=?;"

	err := c.client.Query(q, email).Scan(
		&result.Name,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			return nil, errs.NewNotFoundError("invalid credentials")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &result, nil
}

func (c *user) Delete(email string) *errs.Err {
	q := "DELETE FROM users WHERE email=? IF EXISTS;"

	err := c.client.Query(q, email).Exec()
	if err != nil {
		if err == gocql.ErrNotFound {
			return errs.NewNotFoundError("no user found for given email")
		}
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}
