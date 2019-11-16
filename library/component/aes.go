package component

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"strconv"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func Encrypt(key, msg string) (string, error) {
	plaintext := []byte(msg)
	// Create the aes encryption algorithm
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "加密异常", errors.New(fmt.Sprintf("Error: NewCipher(%d bytes) = %s", len(key), err))
	}
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	cipherText := make([]byte, len(plaintext))
	cfb.XORKeyStream(cipherText, plaintext)
	return fmt.Sprintf("%x", cipherText), nil
}

func Decrypt(key, msg string) (string, error) {
	// new add
	var text []byte
	tmp := ""
	for k, v := range msg {
		if k%2 == 0 {
			tmp = ""
		}
		tmp += string(v)
		if len(tmp) == 2 {
			tmp, _ := strconv.ParseInt(tmp, 16, 10)
			text = append(text, byte(tmp))
		}

	}
	// Decrypt strings
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "解密异常", errors.New(fmt.Sprintf("Error: NewCipher(%d bytes) = %s", len(key), err))
	}
	cfbDec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(text))
	cfbDec.XORKeyStream(plaintextCopy, text)
	return fmt.Sprintf("%s", plaintextCopy), nil
}
