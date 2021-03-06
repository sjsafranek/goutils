package hashers

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"

	"github.com/sjsafranek/goutils/transformers"
)

func MD5HashString(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha512HashString(text string) string {
	hasher := sha512.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func MD5HashByte(text string) []byte {
	return transformers.ToByte(MD5HashString(text))
}

func Sha512HashByte(text string) []byte {
	return transformers.ToByte(Sha512HashString(text))
}

// func ToByte(text string) []byte {
// 	return []byte(text)
// }
