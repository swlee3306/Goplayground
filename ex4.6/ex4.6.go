package main

import "fmt"

var g int = 10 //g 는 이 패키지 안에서 사용 가능한 전역 변수 이다.

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	//m = s + 20 s 값이 중괄호 안에 까지만 유효함
}
