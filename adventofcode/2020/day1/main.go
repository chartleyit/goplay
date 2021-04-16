package main

import (
	"fmt"
	"io/ioutil"
)

const url = "https://adventofcode.com/2020/day/1/input"
const file = "data.txt"

// TODO learn how to auth/cookie handle adventofcode.com
// Get data from url
// url: https://adventofcode.com/2020/day/1/input
// func GetData(u string) string {
// 	resp, err := http.Get(u)
// 	if err != nil {
// 		fmt.Println("Unable to access website")
// 		fmt.Println(err)
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println("Unable to read response body")
// 		fmt.Println(err)
// 	}
// 	return string(body)
// }

// ReadData from file
func ReadData(f string) string {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("Unable to read file")
		fmt.Println(err)
	}
	return string(data)
}

// Insert data
// O(log n) or O(height)
// Search data

func main() {
	// fmt.Println(GetData(url))
	input := ReadData(file)
	fmt.Println(input)
}
