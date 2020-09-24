package main

import "fmt"

func main() {
	println("Hello world!")
	var a, c int
	a = 3
	b := 5
	c = a + b
	d := "Bye world.."
	var k int = 3
	const sangsu = "Punto"
	const (
		sangsu1 = iota
		sangsu2
		sangsu3
	)

	msg := "IBI"
	saysomething("Hi", "I", "am", "Punto")

	println(c, d, sangsu3)

	if k == 1 {
		println("One")
	} else if k == 2 {
		println("Two")
	} else {
		println("Other")
	}

	var name string
	var category = 3
	x := category << 1
	switch x {
	case 1:
		name = "Punto"
	case 2:
		name = "DonJon"
	case 3, 4:
		name = "Ramge"
	default:
		name = "Asson"
	}

	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	println(name, x)
	var inpNum = x
	fmt.Println(dec2bin(inpNum))
	println(sum)
	say(&msg)
	println(msg)
}

func saysomething(txt ...string) {
	for _, s := range txt {
		println(s)
	}
}

func say(msg *string) {
	println(*msg)
	*msg = "Changed"
}

func dec2bin(dec int) string {
	bin := ""
	if dec >= 1 {
		bin = fmt.Sprintf("%d", dec%2) + bin
		bin = dec2bin(dec/2) + bin
	}
	return bin
}

func calc() {
	k := 1
	if k == 1 {
		println("One")
	}
}
