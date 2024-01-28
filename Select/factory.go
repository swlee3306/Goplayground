package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
 아래 프로그램은 select 문을 이용해서 프로그램이 멈출 때 까지 무한적으로 차 혹은 비행기를 생산하는 프로그램이다.
*/

// 자동차의 최종 형태를 담을 구조체
type Car struct {
	val string
}

// 비행기의 최종 형태를 담을 구조체
type Plane struct {
	val string
}

// 들어온 채널을 통해서 타이어를 추가 하고 그 추가한 채널 변수를 반환한다. * 무한 루프
// 각 들어온 채널중에 비행기는 비행이 타이어를 추가하고 자동차는 자동차 타이어를 추가해준다 * select 문 이용
func MakeTire(carChan chan Car, outCarChan chan Car, planeChan chan Plane, outChanPlane chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Tire(car), "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Tire(plane), "
			outChanPlane <- plane
		}
	}
}

// 들어온 채널을 통해서 엔진을 추가 하고 그 추가한 채널 변수를 반환한다. * 무한루프
// 각 들어온 채널중에 비행기는 비행이 엔진을 추가하고 자동차는 자동차 엔진을 추가해준다 * select 문 이용
func MakeEngine(carChan chan Car, outCarChan chan Car, planeChan chan Plane, outChanPlane chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Engine(car), "
			outCarChan <- car
		case plane := <-planeChan:
			plane.val += "Engine(plane), "
			outChanPlane <- plane
		}
	}
}

// 각 생산할 차에 처음 시작 채널을 생성하는 함수
func StarCarWork(chan1 chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Car{val: "Car" + strconv.Itoa(i) + ": "}
		i++
	}
}

// 각 생산할 비행기에 처음 시작 채널을 생성하는 함수
func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane" + strconv.Itoa(i) + ": "}
		i++
	}
}

func main() {

	//3개의 채널을 생성한다. * 자동차
	carChan1 := make(chan Car)
	carChan2 := make(chan Car)
	carChan3 := make(chan Car)

	//3개의 채널을 생성한다. * 비행기
	planeChan1 := make(chan Plane)
	planeChan2 := make(chan Plane)
	planeChan3 := make(chan Plane)

	/*
		총 2개의 스레드를 실행 시킨다.

		1. 2개 채널을 입력 받아 타이어를 추가해서 다른 채널의 아웃풋으로 반환한다.
		   ㄴ chan1 을 입력받아 타이어를 붙이고 chan2 로 반환 한다.

		2. 2개 채널을 입력 받아 엔진을 추가해서 다른 채널의 아웃풋으로 반환한다.
		   ㄴ chan2 을 입력받아 타이어를 붙이고 chan3 로 반환 한다.

		즉 이렇게 받은 상태에서 최종적으로 chan3 에 타이어랑 엔진을 추가한 car 구조체가 저장된다.

		위 방식으로 동일하게 비행기도 생산 할수 있도록 같은 방식의 채널을 생성해서 차와 비행기를 동시에 생산한다.
	*/
	go StarCarWork(carChan1)
	go StartPlaneWork(planeChan1)
	go MakeTire(carChan1, carChan2, planeChan1, planeChan2)
	go MakeEngine(carChan2, carChan3, planeChan2, planeChan3)

	//go routine 으로 돌아가고있는 타이어 함수 엔진 함수 두개에 체널 값을 넣어주기 전 Car 구조체를 하나 초기화 해서 chan1 에 넣어준다.(push 과정)
	//plane 채널에도 위 구성과 동일하게 구조가 확립됨

	for {
		//chan1 이 생성되었을때 모든 go routine 함수를 수행 하고 최종 결과 chan3번을 결과에 담는다.
		select {
		case result := <-carChan3:
			fmt.Println("car make" + result.val)
		case result := <-planeChan3:
			fmt.Println("plane make" + result.val)
		}

	}
}

// 위 프로그램에서 중요한 부분은 각 채널 마다 push,pop 을 명확하게 해주어 deadrock 이 발생하지 않도록 주의해야 한다.
