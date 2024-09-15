package commons

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// PKCS5Padding 使用 PKCS5/PKCS7 方式填充数据
func pKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

// PKCS5UnPadding 去除填充数据
func pKCS5UnPadding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func EncryptAES(plaintext string, key string) (string, error) {
	// create cipher
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// allocate space for ciphered data
	out := make([]byte, len(plaintext))

	// encrypt
	c.Encrypt(out, []byte(plaintext))
	// return hex string
	return base64.StdEncoding.EncodeToString(out), nil
}

// AESEncrypt 加密方法
func AESEncrypt(plaintext string, key string) (string, error) {

	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		return "", err
	}

	plaintextBytes := pKCS5Padding([]byte(plaintext), block.BlockSize())
	ciphertext := make([]byte, aes.BlockSize+len(plaintextBytes))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintextBytes)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AESDecrypt 解密方法
func AESDecrypt(ciphertext string, key string) (string, error) {
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	keybyte := []byte(key)
	block, err := aes.NewCipher(keybyte)
	if err != nil {
		return "", err
	}

	if len(ciphertextBytes) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := ciphertextBytes[:aes.BlockSize]
	ciphertextBytes = ciphertextBytes[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertextBytes, ciphertextBytes)

	plaintextBytes := pKCS5UnPadding(ciphertextBytes)
	return string(plaintextBytes), nil
}
