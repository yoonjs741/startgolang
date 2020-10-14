package main

import "fmt"

func main() {
	println("START")
	var price = 39900
	var coupon = 70
	calc(price, coupon)
	println("END")
}

func calc(a int, b int) {
	c := float32(b) / 100
	d := float32(a) - (float32(a) * float32(c))
	result := fmt.Sprintf("%d원 입니다.", int(d))
	println(result)
}
