package cryptic

import "unsafe"

type Secret string

func XORString(message string, sec Secret) string {
	var l = len(sec)
	var b = make([]byte, 0, len(message))
	for i := 0; i < len(message); i++ {
		b = append(b, message[i]^sec[i%l])

	}
	return *(*string)(unsafe.Pointer(&b))
}
