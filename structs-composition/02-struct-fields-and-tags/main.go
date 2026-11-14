package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Salary     float64 `json:"salary"`
	Department string  `json:"department"`
}

func main() {

	emp := Employee{
		ID:         01,
		Salary:     400000,
		Name:       "John",
		Email:      "john@gmail.com",
		Department: "Accounts",
	}

	fmt.Printf("%+v", emp)

	// struct to json
	data, err := json.Marshal(emp)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	empJson = {
  "id": 2,
  "name": "Ankit",
  "email": "ankit@gmail.com",
  "salary": 500000,
  "department": "Engineering"
}




}
