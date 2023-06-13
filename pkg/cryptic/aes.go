package cryptic

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"unsafe"
)

func AESEncrypt(plainMessage string, sec Secret) (string, error) {
	block, err := aes.NewCipher([]byte(sec))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	plainBytes := *(*[]byte)(unsafe.Pointer(&plainMessage))
	plainBytes = pkcs7Padding(plainBytes, blockSize)

	var dst = make([]byte, len(plainBytes))
	blockMode := cipher.NewCBCEncrypter(block, []byte(sec)[:blockSize])
	blockMode.CryptBlocks(dst, plainBytes)
	return hex.EncodeToString(dst), nil
}

func AESDecrypt(encodedMessage string, sec Secret) (string, error) {
	cipherText, err := hex.DecodeString(encodedMessage)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(sec))
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()

	var dst = make([]byte, len(cipherText))
	blockMode := cipher.NewCBCDecrypter(block, []byte(sec)[:blockSize])
	blockMode.CryptBlocks(dst, cipherText)

	dst = pkcs7UnPadding(dst)
	//(*reflect.StringHeader)((*reflect.SliceHeader)(unsafe.Pointer(&dst)))
	return *(*string)(unsafe.Pointer(&dst)), nil
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func pkcs7UnPadding(data []byte) []byte {
	length := len(data)
	unPadding := int(data[length-1])
	return data[:(length - unPadding)]
}
