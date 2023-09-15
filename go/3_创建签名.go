package main

import (
	"crypto/md5"
	"fmt"
	"sort"
	"strings"
)

func genSign(data map[string]string, token string) string {
	dataArr := []string{}

	for k, _ := range data {
		if k == "sign" {
			continue
		}
		dataArr = append(dataArr, k)
	}

	sort.Strings(dataArr)

	str := ""

	for _, key := range dataArr {
		str += key + "=" + data[key] + "&"
	}

	str = strings.Trim(str, "&")

	str += token // 完成加密字符串的拼接

	sign := md5.Sum([]byte(str))

	return fmt.Sprintf("%x", sign)

}

func main() {

	data := make(map[string]string)

	data["name"] = "张三"
	data["sex"] = "男"
	data["age"] = "18"
	data["school"] = "桂林电子科技大学"

	token := "ashasjdakjsxkamc"

	res := genSign(data, token)

	fmt.Println(res)
}
