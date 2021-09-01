//go:generate mockgen -destination=../../mocks/platform/jwt/mock.go -package=mock_jwt github.com/muhsatrio/golang-boilerplate/platform/jwt JwtAdapter
package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/muhsatrio/golang-boilerplate/domain"
)

// UserPersistence contains the list of functions for database table users
type JwtAdapter interface {
	GenerateToken(claim domain.TokenClaim) (token string, err error)
}

type jwtRepo struct {
	SigningKey    string
	Expiry        int
	SigningMethod *jwt.SigningMethodHMAC
}

// Init is to init the jwt persistence that contains data accounts
func Init(signingKey string, expiry int) JwtAdapter {
	return jwtRepo{
		SigningKey:    signingKey,
		Expiry:        expiry,
		SigningMethod: jwt.SigningMethodHS256,
	}
}

type CustomClaims struct {
	*jwt.StandardClaims
	Email string `json:"email"`
	Name  string `json:"name"`
}

// JWT platform function

func (r jwtRepo) GenerateToken(claim domain.TokenClaim) (token string, err error) {
	claims := CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(r.Expiry)).Unix(),
		},
		claim.Email,
		claim.Name,
	}

	t := jwt.NewWithClaims(r.SigningMethod, claims)

	token, err = t.SignedString([]byte(r.SigningKey))
	if err != nil {
		return
	}

	return
}
