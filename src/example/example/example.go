package test

import (
	"fmt"
)

func forExample1() {
	// for example1
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func forExample2() {
	// for example2
	sum, i := 0, 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}

func forExample3() {
	sum, i := 0, 0
	for {
		if i >= 10 {
			break
		}
		sum += i
		i++
	}
	fmt.Println(sum)
}
