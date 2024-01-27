package main

import "fmt"

func main() {
	var a float32 = 1234.523
	var b float32 = 3456.123
	var c float32 = a * b
	var d float32 = c * 3

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	fmt.Println(d)

	/*
		1234.523
		3456.123
		4.266663e+06
		1.2799989e+07
	*/

	// 연산이 커질수록 실수 연산이 문제가 생기기 때문에 관련 오차 범위 부분을 고려해서 프로그래밍 해야 한다.
}
