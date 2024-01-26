package main

import "fmt"

func main() {
	a := 3
	var b float64 = 3.5

	// var c int = b
	// d := a * b //둘의 타입이 달라서 연산 불가능

	// var e int64 = 7
	// f := a * e //둘이 타입이 같아도 사이즈가 다르면 연산 불가능

	var c int = int(b)
	d := float64(a) * b

	var e int64 = 7
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)

	//각 타입을 서로 맞춰서 연산 해준다
	//출력값 : 3 3.5 3 10.5 7 21
}
