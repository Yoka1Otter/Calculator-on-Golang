/*
--Made by Yoka1Otter
https://github.com/Yoka1Otter
*/

package main

import "fmt"

func main() {
	var a, b int
	var choice int
	fmt.Print("Welcome to the calculator" + "\n")
	fmt.Print("Enter the data (First Number): ")
	fmt.Scanln(&a)
	fmt.Print("Enter the data (Second Number): ")
	fmt.Scanln(&b)

	fmt.Println("Select an operation: " + "\n" + "1. Addition" + "\n" + "2. Subtraction" + "\n" + "3. Multiplication" + "\n" + "4. Division" + "\n")
	fmt.Println("Enter choice: ")
	fmt.Scanln(&choice)

	if choice > 4 {
		fmt.Print("Error")
	} else if choice < 1 {
		fmt.Print("Error")
	}

	switch choice {
	case 1:
		fmt.Println(a + b)
	case 2:
		fmt.Println(a - b)
	case 3:
		fmt.Println(a * b)
	case 4:
		fmt.Println(a / b)
	}

}
