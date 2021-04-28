package main

import (
	"fmt"
	"math"
)

func numberList() {
	number := 50.0
	guess := 30.0
	sq := math.Pow(guess, 2)
	root := math.Sqrt2
	newRoot := math.Sqrt(sq)

	if guess < number {
		fmt.Printf("sq: %v, root of two:%v, root of sq: %v \n", sq, root, newRoot)
	}
	if guess >= number {
		fmt.Println("Too high")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	slc := []int{2, 4, 8}
	for k, v := range slc {
		fmt.Println(k, v)
	}

}
