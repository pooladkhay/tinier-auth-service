package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pooladkhay/tinier-auth-service/helper/errs"
)

func generateToken(email string) (*string, *errs.Err) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = email
	claims["exp"] = time.Now().UTC().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("USER_JWT_SECRET")))
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &t, nil
}
