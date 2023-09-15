package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
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

/*
AES对称解密 - 密码反馈模式（Cipher FeedBack (CFB)）
token 长度需要满足
	16位 => AES-128加密
	24位 => AES-192加密
	32位 => AES-256加密
*/
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

/**
sha1 方式加密
*/
func getSha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

/**
base64加密
baseTable：一个64位的字符串
首先使用Base64编码需要一个含有64个字符的表，这个表由大小写字母、数字、+和/组成
采用Base64编码处理数据时，会把每三个字节共24位作为一个处理单元，再分为四组，每组6位，查表后获得相应的字符即编码后的字符串
*/
func getBase64Encode(str string, baseTable string) string {
	code := base64.NewEncoding(baseTable)
	return code.EncodeToString([]byte(str))
}

/**
base64解密
baseTable：一个64位的字符串
首先使用Base64编码需要一个含有64个字符的表，这个表由大小写字母、数字、+和/组成
采用Base64编码处理数据时，会把每三个字节共24位作为一个处理单元，再分为四组，每组6位，查表后获得相应的字符即编码后的字符串
*/
func getBase64Decode(str string, baseTable string) string {
	code := base64.NewEncoding(baseTable)
	res, _ := code.DecodeString(str)
	return string(res)
}

func main() {

	a := "asadfsd哈哈哈哈ha"
	token := "qwsdsdfsavdsavds"

	res := getMd5(a)
	fmt.Println(res)

	res = getAesEncryptCFB(a, token)
	fmt.Println(res)

	res = getAesDecryptCFB(res, token)
	fmt.Println(res)

	res = getSha1(a)
	fmt.Println(res)

	base64Table := "IJjkKLMNO567PQX12RVW3YZaDEFGbcdefghiABCHlSTUmnopqrxyz04stuvw89+/"

	res = getBase64Encode(a, base64Table)
	fmt.Println(res)

	res = getBase64Decode(res, base64Table)
	fmt.Println(res)

}
