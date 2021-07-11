package jwt

import "github.com/muhsatrio/golang-boilerplate/domain"

// UserPersistence contains the list of functions for database table users
type JwtAdapter interface {
	GenerateToken(claim domain.TokenClaim) (token string, err error)
}

type jwtRepo struct {
	SigningKey string
	Expiry     int
}

// Init is to init the jwt persistence that contains data accounts
func Init(signingKey string, expiry int) JwtAdapter {
	return jwtRepo{
		SigningKey: signingKey,
		Expiry:     expiry,
	}
}

// JWT platform function

func (r jwtRepo) GenerateToken(claim domain.TokenClaim) (token string, err error) {
	panic("implement me")
}
