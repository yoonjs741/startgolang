package main

import "fmt"

func main() {
	println("START")
	var a = 10
	for i := 1; i <= a; i++ {
		sum(i)
	}
	println("END")
}

func sum(a int) {
	var factorial int = 1
	//
	// a의 팩토리얼 구하는 함수 자리
	//
	for i := 1; i <= a; i++ {
		factorial *= i
	}
	result := fmt.Sprintf("%d! = %d", a, factorial)
	println(result)
}
