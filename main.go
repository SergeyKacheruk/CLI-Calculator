package main

import "fmt"

func main() {

	var a, b float64
	var operator string

	fmt.Println("Калькулятор на Go")

	fmt.Print("Вводим первое число:")
	fmt.Scan(&a)

	fmt.Print("Вводим оператор: (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Вводим второе число:")
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
