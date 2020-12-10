package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ar [5]int
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 0; i < len(ar); i++ {
		ar[i] = r1.Intn(100)
	}
	fmt.Println("rand :", ar)

	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				nTmp := ar[i]
				ar[i] = ar[j]
				ar[j] = nTmp
			}
		}
	}
	fmt.Println("sort :", ar)
	/*fmt.Println(len(ar))*/
}
