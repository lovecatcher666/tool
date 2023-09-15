package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
)

/**
md5加密
*/
func getMd5(str string) string {
	res := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", res)
}

/*
AES对称加密 - 密码反馈模式（Cipher FeedBack (CFB)）
token 长度需要满足
	16位 => AES-128加密
	24位 => AES-192加密
	32位 => AES-256加密
*/
func getAesEncryptCFB(str string, token string) string {

	origData := []byte(str)
	key := []byte(token)
	block, _ := aes.NewCipher(key)

	encrypted := make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		//panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return string(encrypted)
}
func getAesDecryptCFB(str string, token string) string {

	encrypted := []byte(str)
	key := []byte(token)
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return string(encrypted)
}
func main() {

	a := "asadfsd哈哈哈哈ha"
	token := "qwsdsdfsavdsavds"
	res := getAesEncryptCFB(a, token)

	fmt.Println(res)

	res = getAesDecryptCFB(res, token)

	fmt.Println(res)

}
