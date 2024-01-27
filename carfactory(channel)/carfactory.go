package main

import "fmt"

// 자동차의 최종 형태를 담을 구조체
type Car struct {
	val string
}

// 들어온 채널을 통해서 타이어를 추가 하고 그 추가한 채널 변수를 반환한다.
func MakeTire(carChan chan Car, outChan chan Car) {
	car := <-carChan
	car.val += "Tire, "

	outChan <- car
}

// 들어온 채널을 통해서 엔진을 추가 하고 그 추가한 채널 변수를 반환한다.
func MakeEngine(carChan chan Car, outChan chan Car) {
	car := <-carChan
	car.val += "Engine, "

	outChan <- car
}

func main() {

	//3개의 채널을 생성한다.
	chan1 := make(chan Car)
	chan2 := make(chan Car)
	chan3 := make(chan Car)

	/*
		총 2개의 스레드를 실행 시킨다.

		1. 2개 채널을 입력 받아 타이어를 추가해서 다른 채널의 아웃풋으로 반환한다.
		   ㄴ chan1 을 입력받아 타이어를 붙이고 chan2 로 반환 한다.

		2. 2개 채널을 입력 받아 엔진을 추가해서 다른 채널의 아웃풋으로 반환한다.
		   ㄴ chan2 을 입력받아 타이어를 붙이고 chan3 로 반환 한다.

		즉 이렇게 받은 상태에서 최종적으로 chan3 에 타이어랑 엔진을 추가한 car 구조체가 저장된다.
	*/
	go MakeTire(chan1, chan2)
	go MakeEngine(chan2, chan3)

	//go routine 으로 돌아가고있는 타이어 함수 엔진 함수 두개에 체널 값을 넣어주기 전 Car 구조체를 하나 초기화 해서 chan1 에 넣어준다.(push 과정)
	chan1 <- Car{val: "Car1: "}

	//chan1 이 생성되었을때 모든 go routine 함수를 수행 하고 최종 결과 chan3번을 결과에 담는다.
	result := <-chan3

	//해당 결과를 출력한다.
	fmt.Println("success make " + result.val)

}
