package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 获取随机字符串
 num => 返回字符串的个数
 t => 返回字符串包括的字符类型
*/
func getRandom(num int, t string) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	switch t {
	case "num":
		str = "0123456789"
	case "str":
		str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "low":
		str = "abcdefghijklmnopqrstuvwxyz"
	case "upper":
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	res := make([]byte, num)
	strArr := []byte(str)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		index := rand.Intn(len(str))
		res[i] = strArr[index]
	}
	return string(res)

}

func main() {
	res := getRandom(6, "str")

	fmt.Println(res)

}
