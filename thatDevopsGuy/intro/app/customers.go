package main

type (
	Customer struct {
		FirstName string
		LastName  string
		FullName  string
	}
)

func GetCustomers() (customers []Customer) {

	chris := Customer{FirstName: "Chris", LastName: "Hartley"}

	customers = append(customers, chris,
		Customer{FirstName: "Bob", LastName: "Smith"},
	)

	return customers
}
