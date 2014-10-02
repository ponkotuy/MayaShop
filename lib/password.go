package lib

import (
	"math/rand"
	"crypto/sha256"
)

var stretchCount = 2057

func CreatePasswordHash(passwd string, salt []byte) []byte {
	var h256 [32]byte
	hashed := make([]byte, 32)
	hashed = []byte(passwd)
	for i := 0; i < stretchCount; i += 1 {
		x := append(hashed, salt...)
		h256 = sha256.Sum256(x)
		hashed = h256[:]
	}
	return hashed
}

func PasswordSalt(passwd string) ([]byte, []byte) {
	salt := RandomBytes(32)
	return CreatePasswordHash(passwd, salt), salt
}

func RandomBytes(num int) []byte {
	result := make([]byte, num)
	for i := 0; i < num; i += 1 {
		result[i] = byte(rand.Intn(256))
	}
	return result
}
