package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Test implementation of "Approximate Counting Algorithm".
// https://en.wikipedia.org/wiki/Approximate_counting_algorithm
// Note that algorithm might not be same as official (standard?) one.

func main() {
	fmt.Println("vim-go")

	aca(1000)
	aca(5000)
	aca(10000)
	aca(20000)

}

func aca(nloop int) {
	var count float64
	var r int

	//count = 1
	for i := 0; i < nloop; i++ {
		//		fmt.Println(rand.Intn(10), int(math.Pow(2, count)))
		r = rand.Intn(int(math.Pow(2, count)))
		if r == 0 {
			count++
			fmt.Println()
			fmt.Printf("%v:%v (%v) | ", i, math.Pow(2, count), r)
		}
		//fmt.Printf("%v:%v (%v) | ", i, math.Pow(2, count), r)
	}
	fmt.Println(count, math.Pow(2, count))
}
