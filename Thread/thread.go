package main

import (
	"fmt"
	"time"
)

//go thread 아래 예제를 살펴보면 fun1 을 동시에 main 함수와 동일하게 수행된다.(섞여서 수행)
//이유는 해당 멀티 스레드를 지원하는 go 언어의 특징이다.
//스레스 시간은 랜덤이다. os 에 따라서 결정됨

func main() {
	go func1()
	for i := 0; i< 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("main", i)
	}
	fmt.Scanln()
}

func func1(){
	for i := 0; i< 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("func1: ",i)
	}
}