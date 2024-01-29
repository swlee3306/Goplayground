package main

/*
OOD 5가지 원칙중 SOLID중 Interface 방식
- 여러 관계를 모아놓은 인터페이스 보다 관계 하나씩 정의 하는게 더 좋다.
*/

//하나의 인터페이스에서 여러 기능을 가지고 있는 예제
//SRP 에 어긋나는 선언이다.
type Actor interface {
	Move()
	Actor()
	Talk()
}


//하나의 인터페이스에 한개의 기능씩만 정의 하도록 가지고 있는 예제
type Talkable interface {
	Talk()
}

type Attackable interface {
	Attack()
}

type Movable interface{
	Move()
}


//SOLID를 한줄로 정리하면 OCP 로 가기 위한 프로그램 패턴이다. 의존성을 내리고 응집선 올리기 위한