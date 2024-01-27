package main

import "fmt"

func pop(c chan int) {
	fmt.Println("Pop function")
	v := <-c
	fmt.Println(v)
}

func main() {
	var c chan int

	//channel 의 기본 사용

	//c = make(chan int, 1)

	// c <- 10
	// v := <-c

	// fmt.Println(v)

	/*
		아래 코드를 보면 0개의 사이즈를 가지는 채널 변수를 선언 한다.
		해당 변수는 channal 값이 0 이기 떄문에 값을 넣어주는 함수를 별도로 제작한다.
		길이가 0개인 channel은 다른 쪽에서 값을 빼주기 전까지 기다린다.

		순서
		1. channel 초기화
		2. pop 함수를 go routin 으로 실행
		3. Pop function 이라는 문구를 출력하고 c 안에 값이 들어오기를 별도 스레드로 기다리고 있는다.
		4. 별도 스레드로 실행중인 pop 함수에 매개변수로 들어가는 c 에 10 이라는 값을 넣어준다.
		5. 값이 들어온 것을 확인한 pop 함수는 해당 값을 출력하고 함수 실행을 종료 한다.
		6. 마지막 줄인 end of program 을 종료 시키고 최종적으로 프로그램을 마무리 한다.
	*/

	c = make(chan int)
	go pop(c)
	c <- 10

	fmt.Println("end of program")
}
