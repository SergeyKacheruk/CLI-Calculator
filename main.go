package main

import "fmt"

func main() {
	var a, b float64
	var operator string

	fmt.Println("Go-calculator")

	fmt.Print("We enter the first number:")
	fmt.Scan(&a)

	fmt.Print("We enter the operator: (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("We enter the second number:")
	fmt.Scan(&b)

	var result float64

	switch operator {
	case "+":
		result = a + b

	case "-":
		result = a - b

	case "*":
		result = a * b

	case "/":
		result = a / b

	}

	fmt.Printf("%.2f %s %.2f = %.2f\n", a, operator, b, result)

}
