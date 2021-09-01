package bcrypt

import "golang.org/x/crypto/bcrypt"

func Compare(reqPassword, password string, saltHash int) (err error) {

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(reqPassword))

	return
}
