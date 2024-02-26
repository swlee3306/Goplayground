package main

import "fmt"

// Node is a linked list node
//T는 제네릭 타입입니다.
//제네릭 타입은 다양한 유형의 데이터를 처리하는 데 사용됩니다.
//제네릭 타입은 함수 또는 구조체를 정의 할 때 사용됩니다.
//제네릭 타입은 코드의 재사용성을 높이고 코드의 중복을 줄이는 데 사용됩니다.
//제네릭 타입은 Go 1.18 버전에서 추가되었습니다.

type Node[T any] struct {
	next *Node[T]
	val T
}

// ArrayLinkedList is a linked list
func main() {
	root := &Node[int]{nil, 10}
	root.next = &Node[int]{nil, 20}
	root.next.next = &Node[int]{nil, 30}
	
	// print linked list
	for p := root; p != nil; p = p.next {
		fmt.Printf("%d -> ", p.val)
	}

	// print linked list
	fmt.Printf("ArrayLinkedList\n")
}

// 출력 : 10 -> 20 -> 30 -> ArrayLinkedList
// 시간 복잡도 : O(n)
// 공간 복잡도 : O(1)
// n은 연결 목록의 노드 수입니다.
// 불연속적인 메모리 공간에 데이터를 저장하므로 연결 목록은 데이터를 삽입하거나 삭제하는 데 유용합니다.
// 위의 코드는 Go에서 연결 목록의 간단한 구현입니다.
// 연결 목록은 각 요소가 다음 요소에 연결된 시퀀스의 요소로 구성된 데이터 구조입니다.
// 연결 목록에는 루트 노드가 있으며이는 목록의 첫 번째 요소를 가리킵니다.
// 목록의 각 요소는 값과 목록의 다음 요소를 가리키는 포인터가있는 노드입니다.
// 위의 코드는 세 개의 노드가있는 연결 목록을 만들고 목록의 요소를 인쇄합니다.

