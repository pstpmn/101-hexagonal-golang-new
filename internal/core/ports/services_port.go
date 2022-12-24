package ports

import "time"

type IUuidService interface {
	Random() string
}

type ICryptoService interface {
	Md5(plaintext string) string
	Bcrypt(plaintext string) (string, error)
	ValidateBcrypt(plaintext string, encrypt string) bool
}

type IJwtService interface {
	Generate(data map[string]interface{}, key string, exp time.Time) (string, error)
	Extract(token string, key string) (map[string]string, error)
}
