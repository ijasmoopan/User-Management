package useCases

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateMD5HashPassword(password string) string {
	bytes := md5.Sum([]byte(password))
	return hex.EncodeToString(bytes[:])
}