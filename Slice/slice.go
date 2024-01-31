package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5} //정적 배열 생성
	b := make([]int, 0, 0)    //동적 배열 생성

	//len 길이, cap 남은 공간

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	fmt.Printf("len(b) = %d\n", len(b))
	fmt.Printf("cap(b) = %d\n", cap(b))

	a = append(a, 1)

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	/*
		만약 배열에 길이가 n개 이고 값이 들어가있는 값이 n-1 개만 들어가 있으면 append 를 하더라도
		새로운 배열을 복제해서 다른 주소에 넣는것이 아닌 같은 주소의 빈 공간에 값을 넣어준다.
	*/

	//예제
	/*
		a_addr 공간을 2개에 배열로 생성하고 해당 공간에 3이란 값을 append 시켜서 b_addr 이라는 새로운 공간에 저장한다.
		결과는 두개가 다른 주소를 가진다.
	*/

	a_addr := []int{1, 2}

	b_addr := append(a_addr, 3)

	fmt.Printf("%p %p\n", a_addr, b_addr)

	for i := 0; i < len(a_addr); i++ {
		fmt.Printf("%d, ", a_addr[i])
	}

	fmt.Println()

	for i := 0; i < len(b_addr); i++ {
		fmt.Printf("%d, ", b_addr[i])
	}
	fmt.Println()

	fmt.Println(cap(a_addr), cap(b_addr))

	//예제
	/*
		c_addr 공간을 2개인데 남은 공간 은 4개로 사용해서 사용할 공간을 2개로만 초기화 해서 append 해서 값을 넣을시 주소가 변경되지 않고 같은 위치에 값만 추가되는것을 표현한 예제이다.
	*/

	c_addr := make([]int, 3, 4)

	d_addr := append(c_addr, 3)

	fmt.Printf("%p %p\n", c_addr, d_addr)

	for i := 0; i < len(c_addr); i++ {
		fmt.Printf("%d, ", c_addr[i])
	}

	fmt.Println()

	for i := 0; i < len(d_addr); i++ {
		fmt.Printf("%d, ", d_addr[i])
	}
	fmt.Println()

	fmt.Println(cap(c_addr), cap(d_addr))

	// 같은 주소의 c_addr 배열이 있기 때문에 결과가 같게 변경된다.
	d_addr[0] = 1
	d_addr[1] = 2

	for i := 0; i < len(c_addr); i++ {
		fmt.Printf("%d, ", c_addr[i])
	}

	fmt.Println()

	for i := 0; i < len(d_addr); i++ {
		fmt.Printf("%d, ", d_addr[i])
	}
	fmt.Println()

}
