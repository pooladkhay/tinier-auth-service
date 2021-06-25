package domain

import "github.com/pooladkhay/tinier-auth-service/helper/errs"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (lr *LoginRequest) Validate() *errs.Err {
	if len(lr.Email) == 0 {
		return errs.NewBadRequestError("email is required")
	}
	if len(lr.Password) == 0 {
		return errs.NewBadRequestError("password is required")
	}
	return nil
}

func (rr *RegisterRequest) Validate() *errs.Err {
	if len(rr.Name) == 0 {
		return errs.NewBadRequestError("name is required")
	}
	if len(rr.Email) == 0 {
		return errs.NewBadRequestError("email is required")
	}
	if len(rr.Password) == 0 {
		return errs.NewBadRequestError("password is required")
	}
	return nil
}
