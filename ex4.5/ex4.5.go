package main

import "fmt"

func main() {
	var a int16 = 3456   //a는 int16타입 - 2 바이트 정수
	var b int8 = int8(a) // int8 로 변환

	fmt.Println(b)
}

// 타입 변환시 들어갈 공간이 부족해서 데이터가 온전히 들어가지 못하는 상황이 벌어졌다.
