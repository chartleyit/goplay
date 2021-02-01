package main

import (
	"fmt"
)

func main() {
	customers := GetCustomers()

	for _, customer := range customers {
		fmt.Println(customer)
	}
}
