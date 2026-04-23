package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

type AESEncryptor struct {
	key []byte
}

// NewAESEncryptor 初始化加密器，金鑰長度需為 32 位元組
func NewAESEncryptor(key string) (*AESEncryptor, error) {
	if len(key) != 32 {
		return nil, errors.New("AES-256-CBC 需要一個 32 位元組的金鑰")
	}
	return &AESEncryptor{key: []byte(key)}, nil
}

// PKCS7Padding 填充
func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去填充
func pkcs7UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("解密數據長度錯誤")
	}
	unpadding := int(origData[length-1])
	if unpadding > length {
		return nil, errors.New("解密數據填充錯誤")
	}
	return origData[:(length - unpadding)], nil
}

// Encrypt 加密文字，返回 Base64URL 編碼字串
func (a *AESEncryptor) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	originData := pkcs7Padding([]byte(plaintext), blockSize)

	// 建立 IV
	ciphertext := make([]byte, blockSize+len(originData))
	iv := ciphertext[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[blockSize:], originData)

	// Base64URL 編碼 (對應 PHP 的 strtr + rtrim)
	return base64.RawURLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt 解密文字
func (a *AESEncryptor) Decrypt(encryptedData string) (string, error) {
	// Base64URL 解碼
	data, err := base64.RawURLEncoding.DecodeString(encryptedData)
	if err != nil {
		return "", errors.New("解密失敗：無效的格式[0]")
	}

	block, err := aes.NewCipher(a.key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	if len(data) < blockSize {
		return "", errors.New("解密失敗：無效的格式[1]")
	}

	iv := data[:blockSize]
	ciphertext := data[blockSize:]

	if len(ciphertext)%blockSize != 0 {
		return "", errors.New("解密失敗：密文長度非區塊倍數")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// 去除 Padding
	result, err := pkcs7UnPadding(ciphertext)
	if err != nil {
		return "", errors.New("解密失敗：無效的格式[2]")
	}

	return string(result), nil
}
