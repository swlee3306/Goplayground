package main

import (
	"fmt"
	"time"
)

// 2초 간격으로 채널에 정수형 i 값을 넣어준다.
func push(c chan int) {
	i := 0
	for {
		time.Sleep(2 * time.Second)
		c <- i
		i++
	}
}

func main() {
	//채널 생성
	c := make(chan int)

	//go routine 으로 해당 함수가 연속적으로 수행된다.
	go push(c)

	//select 문을 이용해서 channel c 안에 값이 있으면 그것을 출력하고 없으면 default 값인 "Idle" 문자열을 1초 간격으로 출력한다.
	// 해당 함수들은 연속성을 확인하기 위해서 무한루프로 생성된다.
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		default:
			fmt.Println("Idle")
			time.Sleep(1 * time.Second)

		}
	}

}
