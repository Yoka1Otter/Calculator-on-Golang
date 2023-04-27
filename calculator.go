/*
--Made by Yoka1Otter
https://github.com/Yoka1Otter
*/

package main

import (
	"fmt"
	"strconv"
)

var romanNumbers = map[string]int{
	"M": 1000, "CM": 900, "D": 500, "CD": 400,
	"C": 100, "XC": 90, "L": 50, "XL": 40,
	"X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1,
}

func toRoman(number int) string {
	roman := ""
	for letter, value := range romanNumbers {
		for number >= value {
			roman += letter
			number -= value
		}
	}
	return roman
}

func main() {
	var x, y string
	var a, b int

	fmt.Println(a, b)

	var itRim int
	var choice string

	fmt.Println("Welcome to the calculator" + "\n")
	fmt.Scanln(&x, &choice, &y)

	if x == "I" {
		x = "1"
		itRim++
	} else if x == "II" {
		x = "2"
		itRim++
	} else if x == "III" {
		x = "3"
		itRim++
	} else if x == "IV" {
		x = "4"
		itRim++
	} else if x == "V" {
		x = "5"
		itRim++
	} else if x == "VI" {
		x = "6"
		itRim++
	} else if x == "VII" {
		x = "7"
		itRim++
	} else if x == "VIII" {
		x = "8"
		itRim++
	} else if x == "IX" {
		x = "9"
		itRim++
	} else if x == "X" {
		x = "10"
		itRim++
	}

	if y == "I" {
		y = "1"
		itRim++
	} else if y == "II" {
		y = "2"
		itRim++
	} else if y == "III" {
		y = "3"
		itRim++
	} else if y == "IV" {
		y = "4"
		itRim++
	} else if y == "V" {
		y = "5"
		itRim++
	} else if x == "VI" {
		y = "6"
		itRim++
	} else if y == "VII" {
		y = "7"
		itRim++
	} else if y == "VIII" {
		y = "8"
		itRim++
	} else if y == "IX" {
		y = "9"
		itRim++
	} else if y == "X" {
		y = "10"
		itRim++
	}

	if itRim == 1 {
		fmt.Println("Error")
	} else if itRim == 2 || itRim == 0 {
		/*for i, key := range roman {
			fmt.Println(i, key)
		}*/

		a, error1 := strconv.Atoi(x)
		if error1 == nil {
			//fmt.Println(x)
		}
		b, e := strconv.Atoi(y)
		if e == nil {
			//fmt.Println(y)
		}

		//fmt.Scanln(&a, &choice, &b)

		if a > 10 || b > 10 || a < 1 || b < 1 {
			fmt.Print("Error")
		} else {

			var result int
			if itRim == 2 && result < 0 {
				fmt.Print("Error")
			}
			switch choice {
			case "+":
				result = a + b
				if itRim == 2 && result < 0 {
					fmt.Print("Error")
				} else {
					if itRim == 2 {
						fmt.Println(toRoman(result))
					} else {
						fmt.Println(result)
					}
				}
			case "-":
				result = a - b
				if itRim == 2 && result < 0 {
					fmt.Print("Error")
				} else {
					if itRim == 2 {
						fmt.Println(toRoman(result))
					} else {
						fmt.Println(result)
					}
				}
			case "*":
				result = a * b
				if itRim == 2 && result < 0 {
					fmt.Print("Error")
				} else {
					if itRim == 2 {
						fmt.Println(toRoman(result))
					} else {
						fmt.Println(result)
					}
				}
			case "/":
				result = a / b
				if itRim == 2 && result < 0 {
					fmt.Print("Error")
				} else {
					if itRim == 2 {
						fmt.Println(toRoman(result))
					} else {
						fmt.Println(result)
					}
				}
			default:
				fmt.Println("Error")
			}
		}
	}
}
