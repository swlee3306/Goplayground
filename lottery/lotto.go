package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// lottery.exe. filename count
	// 실행 파일명, 읽을 텍스트 파일 이름, 뽑을 숫자 를 3가지 매개변수를 받는다. ex) ./lottery randomtxt.txt 4
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Invalid arguments!\nlottery filename count")
		return
	}

	// 2번 쨰로 읽은 문자열 txt 파일 명
	filename := os.Args[1]

	// 3번 째로 읽은 문자열 뽑을 인원
	// ㄴ 해당 데이터가 문자열 이기 때문에 숫자로 인식 할수 있는 정수형 타입으로 변경 필요
	count, err := strconv.Atoi(os.Args[2])

	//뽑을 인원 문자를 숫자로 변경하지 못하는 경우 에러 발생 예외 처리
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot convert count to integer! count:", count)
		return
	}

	// 해당 파일을 읽어서 문자열 배열 변수에 저징한다. 아래는 실패시 예외 처리 코드도 함께 작성되어 있다.
	candidates, err := readCandidates(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read candidates file! :", err)
		return
	}

	//현재 시간 기준으로 랜덥인 시드를 생성 한다.
	rand.Seed(time.Now().UnixNano())

	//뽑을 인원 만큼 문자열 배열을 생성 한다.
	winners := make([]string, count)

	//뽑을 인원 만큼 읽어온 파일에서 랜덤으로 돌려서 나온 인원 한 문자씩 해당 winners 문자열 배열에 append 시켜준다
	// * 이떄 같은 인원이 들어가지 않도록 해당 인원 제외한 문자열을 다시 candidates 문자열 배열에 다시 append 해서 저장한다.
	for i := 0; i < count; i++ {
		n := rand.Intn(len(candidates))
		winners[i] = candidates[n]
		candidates = append(candidates[:n], candidates[n+1:]...)
	}

	//winners 문자열 배열을 읽어와 이긴 인원을 한줄씩 출력
	fmt.Println("Winners are !!!")
	for _, winner := range winners {
		fmt.Println(winner)
	}
}

// 파일을 한줄씩 읽어서 문자열 배열로 반환하는 함수
func readCandidates(filename string) ([]string, error) {
	// 파일 오픈 함수
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	// 오픈후 문제가 없으면 해당 파일을 클로즈 한다.
	// * defer 함수 위 함수 수행 후 나중에 처리 하는 go 함수
	defer file.Close()

	//파일에서 한줄씩 읽어서 scanner 배열에 저장
	scanner := bufio.NewScanner(file)

	//읽어온 파일을 문자열 배열로 저장 하기 위한 candidates 문자열 배열을 생성
	var candidates []string

	//해당 문자열을 읽어와서 한줄씩 canditates 문자열 배열로 저장
	for scanner.Scan() {
		candidates = append(candidates, scanner.Text())

	}

	//읽어온 함수가 문제가 없을시 해당 문자열을 리턴
	return candidates, nil

}
