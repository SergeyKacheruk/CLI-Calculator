package main

import "fmt"

func main() {

	for {
		var a, b, result float64
		var operator string
		var choice int

		fmt.Println("Go-calculator")

		getInput("The first number: ", &a)

		getInput("Operator (+, -, *, /): ", &operator)
		getInput("The second number: ", &b)

		switch operator {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			if b == 0 {
				fmt.Println("Error! Division by zero is impossible!")
				continue
			}
			result = a / b
		default:
			fmt.Println("Error! Invalid mark!")
			continue
		}
		fmt.Printf("%.2f %s %.2f = %.2f\n\n", a, operator, b, result)
		fmt.Print("Do you want to continue? 1 - Yes, 2 - No: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			continue
		case 2:
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("Error!! Invalid number!")
		}

	}
}
func getInput(msg string, inp any) error {
	for {
		fmt.Print(msg)

		_, err := fmt.Scan(inp)

		return err
	}
}
