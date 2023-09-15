package main

import "fmt"

func checkEqualSlice(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true

}

func main() {

	a := []string{
		"a", "b", "c",
	}

	b := []string{
		"a", "b", "c",
	}

	res := checkEqualSlice(a, b)

	fmt.Println(res)

}
