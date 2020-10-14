package main

import "fmt"

func main() {
	println("START")
	a := 77879
	odd(a)
	println("END")
}

func odd(i int) {
	resultodd := fmt.Sprintf("%d는 홀수입니다.", i)
	resultadd := fmt.Sprintf("%d는 짝수입니다.", i)
	resultpositive := fmt.Sprintf("%d는 양수입니다.", i)
	resultnegative := fmt.Sprintf("%d는 음수입니다.", i)
	if i > 0 {
		println(resultpositive)
		switch {
		case i%2 == 1:
			println(resultodd)
		case i%2 == 0:
			println(resultadd)
		}
	} else {
		println(resultnegative)
		j := -i
		switch {
		case j%2 == 1:
			println(resultodd)
		case j%2 == 0:
			println(resultadd)
		}
	}

}
