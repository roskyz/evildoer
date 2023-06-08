package cryptic

import (
	"crypto/aes"
	"unsafe"
)

func AESEncrypt(plainMessage string, sec Secret) (string, error) {
	cipher, err := aes.NewCipher([]byte(sec))
	if err != nil {
		return "", err
	}
	var b = make([]byte, len(plainMessage))
	cipher.Encrypt(b, []byte(plainMessage))
	return *(*string)(unsafe.Pointer(&b)), nil
}

func AESDecrypt(encodedMessage string, sec Secret) (string, error) {
	cipher, err := aes.NewCipher([]byte(sec))
	if err != nil {
		return "", err
	}
	var b = make([]byte, len(encodedMessage))
	cipher.Decrypt(b, []byte(encodedMessage))
	return *(*string)(unsafe.Pointer(&b)), nil
}
