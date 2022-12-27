package pkg

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type ICrypto interface {
	Md5(plaintext string) string
	Bcrypt(plaintext string) (string, error)
	ValidateBcrypt(plaintext string, encrypt string) bool
}

type crypto struct {
}

func NewCrypto() ICrypto {
	return &crypto{}
}

func (c crypto) ValidateBcrypt(plaintext string, encrypt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(plaintext))
	if err != nil {
		return false
	}
	return true
}

func (c crypto) Md5(plaintext string) string {
	data := []byte(plaintext)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (c crypto) Bcrypt(plaintext string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plaintext), 14)
	return string(bytes), err
}
