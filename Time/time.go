package main

import (
	"fmt"
	"time"
)

// 2초 간격으로 채널에 정수형 i 값을 넣어준다.
func push(c chan int) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		c <- i
		i++
	}
}

func main() {
	//채널 생성
	c := make(chan int)

	//go routine 으로 해당 함수가 연속적으로 수행된다.
	go push(c)

	timer := time.After(10 * time.Second)
	tickTimer := time.Tick(2 * time.Second)

	//select 문을 이용해서 channel c 안에 값이 있으면 그것을 출력하고 없으면 대기하고 있다가 10초 후에 타임 아웃이라는 문자열을 출력하고 리턴된다.
	//tickTimer 채널을 통해서 현재 2초마다 Tick 이라는 문자열을 반복하도록 설정 했다.
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-timer:
			fmt.Println("timeout")
			return
		case <-tickTimer:
			fmt.Println("Tick")
		}
	}

}
