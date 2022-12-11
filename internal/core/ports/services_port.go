package ports

type IUuidService interface {
	Random() string
}

type ICrypto interface {
	Md5(plaintext string) string
	Bcrypt(plaintext string) (string, error)
	ValidateBcrypt(plaintext string, encrypt string) bool
}
