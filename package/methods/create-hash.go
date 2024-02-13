package methods

import "golang.org/x/crypto/bcrypt"

func HashToken(refresh string) (string, error) {
	token := []byte(refresh)
	res, err := bcrypt.GenerateFromPassword(token, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
