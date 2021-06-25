package service

import (
	"fmt"
	"time"

	"github.com/pooladkhay/tinier-auth-service/domain"
	"github.com/pooladkhay/tinier-auth-service/helper/errs"
	"github.com/pooladkhay/tinier-auth-service/repository"
	"golang.org/x/crypto/bcrypt"
)

type Auth interface {
	Login(d *domain.LoginRequest) (*domain.LoginResponse, *errs.Err)
	Register(d *domain.RegisterRequest) *errs.Err
	DeleteUser(email string) *errs.Err
}

type auth struct {
	userRepo repository.User
}

func NewAuth(u repository.User) Auth {
	return &auth{userRepo: u}
}

func (a *auth) Login(d *domain.LoginRequest) (*domain.LoginResponse, *errs.Err) {
	user, err := a.userRepo.GetByEmail(d.Email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(d.Password)); err != nil {
		return nil, errs.NewUnauthorizedError("invalid credentials")
	}

	resp := new(domain.LoginResponse)
	resp.Name = user.Name
	resp.Email = user.Email

	token, err := generateToken(user.Email)
	if err != nil {
		return nil, err
	}
	resp.AccessToken = fmt.Sprintf("Bearer %s", *token)

	return resp, nil
}

func (a *auth) Register(d *domain.RegisterRequest) *errs.Err {

	if user, _ := a.userRepo.GetByEmail(d.Email); user != nil {
		return errs.NewConflictError("user already exists")
	}

	user := new(domain.User)
	user.Name = d.Name
	user.Email = d.Email
	user.CreatedAt = time.Now().UTC()

	// encrypting passwords with bcrypt
	encryptedPass, err := bcrypt.GenerateFromPassword([]byte(d.Password), bcrypt.DefaultCost)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	user.Password = string(encryptedPass)

	if err := a.userRepo.Create(user); err != nil {
		return err
	}

	return nil
}

func (a *auth) DeleteUser(email string) *errs.Err {
	err := a.userRepo.Delete(email)
	if err != nil {
		return err
	}
	return nil
}
