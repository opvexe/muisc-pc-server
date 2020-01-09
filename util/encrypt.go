package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

/*
	获取MD5
*/
func MD5(s string) string {
	m := md5.New()
	io.WriteString(m, s)
	return strings.ToLower(fmt.Sprintf("%x", m.Sum(nil)))
}

/*
	对称加密
*/
func AesEncrypt(origData, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	result := base64.StdEncoding.EncodeToString(crypted)
	return result, nil
}

/*
	对称解密
*/
func AesDecrypt(crypted, key []byte) (string, error) {
	ciphertext := strings.Replace(string(crypted), " ", "", -1)
	cryptedOri, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cryptedOri))
	blockMode.CryptBlocks(origData, cryptedOri)
	origData = PKCS7UnPadding(origData)
	return string(origData), nil
}

/*
	加密填充
*/
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

/*
	解密填充
*/
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
