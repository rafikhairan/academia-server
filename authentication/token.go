package authentication

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rafikhairan/academia/config"
	"github.com/rafikhairan/academia/model"
	"time"
)

type UserClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user model.User) string {
	JWTConfig := config.AppConfig.JWT

	claims := UserClaims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.AppConfig.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * JWTConfig.Expiration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(JWTConfig.SignatureKey))

	return signedToken
}
