package main

import (
	"fmt"
	"strconv"
)

// here I'm going to try to solve the task of finding
// two values whos sum is 2020 from the input slice
// list is unsorted

// chop - returns 2 arrays cut in the middle

// Split uses 2 arrays at once to find value
func Split(i *[]int) {
	// // spent all day working on fucking iterating over shit and fucking sorted right fucking there
	// sorted := *i
	// sort.Ints(sorted)
	// if len(i)%2 != 0 {
	// 	fmt.Println("this doesn't work with odd length arrays")
	// 	return
	// }
	// a := *i[:len(i)/2]
	// b := *i[len(i)/2:]

	// fmt.Println(len(a))
	// fmt.Println(len(b))
}

// Hardway recusirvely searches array for 2 values that equal 2020
func HardWay(input []string) (int, int) {
	a, _ := strconv.Atoi(input[0])
	var b int
	if len(input) > 1 {
		for j := 0; j <= len(input)-1; j++ {
			b, _ = strconv.Atoi(input[j])
			if a+b == 2020 {
				break
			} else if j == len(input)-1 {
				a, b = HardWay(input[1:])
			}
		}
	}

	return a, b
}

// FindThree returns three numbers that equal target
func FindThree(numbers []int, target int) (int, int, int) {
	var a, b, c int
	l := len(numbers) - 1
	n := numbers[l/2]

	if 2020-n > numbers[l] {
		fmt.Println(n, "Number doesn't exist")
	} else if 2020-n < numbers[0] {
		fmt.Println(n, "Number doesn't exist")
	}

	return a, b, c
}
