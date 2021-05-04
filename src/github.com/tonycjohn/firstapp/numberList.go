package main

import (
	"fmt"
	"math"
	"strings"
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

func codingBat() bool {
	numberArray := []int{1, 1, 2, 4}
	for i, num := range numberArray {
		if i >= 2 && num == 3 && numberArray[i-1] == 2 && numberArray[i-2] == 1 {
			return true
		}
	}
	return false
}

func factorial(num int) int {
	f := num
	i := num - 1
	for i >= 1 {
		f = f * i
		i = i - 1
	}
	return f
}

//reversing and splosion
func stringReverse(originalString string) string {
	var reverseString string
	var splosionString string
	for i, s := range originalString {
		fmt.Println(i, s)
		fmt.Println(originalString[0:i])
		splosionString += originalString[0:i]
	}
	fmt.Println(splosionString)

	for i := len(originalString) - 1; i >= 0; i-- {
		reverseString += string(originalString[i])
	}
	return reverseString
}

func longestWord(sentence string) string {
	longWord, length := "", 0
	for _, word := range strings.Split(sentence, " ") {
		if len(word) > length {
			longWord, length = word, len(word)
		}
	}
	return longWord
}

func sortyByParity(input []int) []int {
	var oddArray []int
	var evenArray []int
	for _, v := range input {
		if v%2 == 0 {
			evenArray = append(evenArray, v)
		} else {
			oddArray = append(oddArray, v)
		}
	}
	evenArray = append(evenArray, oddArray...)
	return evenArray
}

func fourSum(input []int, target int) map[string]int {
	mapOfTwo := make(map[string]int)

	for i, v := range input {
		if i < len(input)-1 {
			subInput := input[i+1 : len(input)-1]
			for j, w := range subInput {
				key := string(i) + string(j)
				value := v + w
				mapOfTwo[key] = value
			}
		}
	}
	var keys []string
	mapOfFour := make(map[string]int)
	for k, _ := range mapOfTwo {
		keys = append(keys, k)
	}

	for i, v := range keys {
		if i < len(keys)-1 {
			subKeys := keys[i+1 : len(keys)-1]
			for _, w := range subKeys {
				if strings.Contains(v, string(w[0])) || strings.Contains(v, string(w[1])) {
					continue
				} else {
					mapOfFour[v+w] = mapOfTwo[v] + mapOfTwo[w]
				}
			}
		}
	}

	for k, v := range mapOfFour {
		if v != target {
			delete(mapOfFour, k)
		}
	}

	fmt.Println(mapOfFour)

	return mapOfFour
}
