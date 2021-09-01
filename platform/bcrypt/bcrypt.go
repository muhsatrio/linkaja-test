package bcrypt

import "golang.org/x/crypto/bcrypt"

func Compare(reqPassword, password string, saltHash int) (err error) {
	hashedReqPass, err := bcrypt.GenerateFromPassword([]byte(reqPassword), saltHash)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(hashedReqPass, []byte(password))

	return
}
