package main

import (
	"fmt"
	"sort"
)

// type expList []int

// // swap
// func (l *expList) swap(x, y int) {
// 	l[x], l[y] = l[y], l[x]
// }

// Sort
func Sort(input *[]int) {

	fmt.Println(input)
	fmt.Println(len(*input))

	sort.Ints(*input)
	fmt.Println(*input)

	// list := &expList{}
	// *list = input

}
