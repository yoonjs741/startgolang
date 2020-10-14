package main

import "fmt"

func main() {
	println("START")
	a := -77879
	var b int
	if a < 0 {
		b = -a
	} else {
		b = a
	}
	odd(a, b)
	println("END")
}

func odd(i, j int) {
	resultodd := fmt.Sprintf("%d는 홀수입니다.", i)
	resultadd := fmt.Sprintf("%d는 짝수입니다.", i)
	if j%2 == 1 {
		println(resultodd)
	} else {
		println(resultadd)
	}
}
