package main

import "fmt"

func main() {
	var a int = 10                    //변수 선언 (정수형)
	var msg string = "hello Variable" //변수 선언 (문자열)

	a = 20                //데이터 교체
	msg = "Good morining" //데이터 교체
	fmt.Println(msg, a)   //출력
}
