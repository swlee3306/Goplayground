package main

/*
OOD 5가지 원칙중 SOLID중  방식

아래 구조 형식을 점검해보면 ocp 방식의 프로그램 구조 이다.
만약 Report 를 만들어 보내는 프로그램을 개발한다고 가정 했을 때

MakeReport 와 SendReport 라는 함수 두가지를 이용해서 Report 를 만드는 함수와 Report 를 전송하는 함수 두가지로만 한다고 했을때
확장성에 있어서 ocp 방식의 설계가 적용되지 않은 방식이다.

만약 SendReport 안에 여러가지 Switch case 문을 이용해서 각 Method 별로 하나의 함수에 여러 기능을 넣어 수정하는것 보다.
go 에서는 각 함수의 기능 확장성을 고려해 인터페이스 문을 이용해서 원하는 기능을 계속 추가해서 함수를 생성할 수 있도록 하는것이 더 괜찮은 설계라고 이야기 한다.

아래는 ocp 방식의 설계 예시이다.
변경에는 닫혀있고, 확장에는 열려있다.
*/ 

type Report struct {

}

type EmailReportSender struct {

}

type FileReportSender struct {

}

type ReportSender interface {
	SendReport(*Report)
}

func (r *Report) MakeReport() *Report {
	return r
}

func (s *EmailReportSender) SendReport(r *Report) {

}

func (s *FileReportSender) SendReport(r * Report){

}