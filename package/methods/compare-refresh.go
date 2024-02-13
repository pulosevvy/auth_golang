package methods

import "golang.org/x/crypto/bcrypt"

func CompareRefresh(hash string, input string) error {
	hashB := []byte(hash)
	inputB := []byte(input)

	err := bcrypt.CompareHashAndPassword(hashB, inputB)
	return err
}
