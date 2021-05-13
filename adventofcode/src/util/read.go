package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const file = "input.txt"
const url = "https://adventofcode.com/2020/day/1/input"

func ReadData() []int {
	// _ = readUrl(url)
	return readFile(file)
}

// ReadData from file
func readFile(f string) (intList []int) {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("Unable to read file")
		fmt.Println(err)
	}
	for _, s := range strings.Split(string(data), "\n") {
		if s == "" {
			continue
		} else {
			newInt, err := strconv.Atoi(s)
			if err != nil {
				// fmt.Println("Error parsing:", s, reflect.TypeOf(s).Kind())
				fmt.Println(err)
				break
			}
			// fmt.Println(s, reflect.TypeOf(s).Kind(), newInt, reflect.TypeOf(newInt).Kind())
			intList = append(intList, newInt)
		}
	}
	return intList
}

// TODO learn how to auth/cookie handle adventofcode.com
// Get data from url
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

func readUrl(u string) (intList []int) {
	secret := os.Getenv("ADVENT_SECRET")
	fmt.Println("ADVENT_SECRET: ", secret)

	return intList
}
