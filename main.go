package main

import "fmt"

func main() {
	fmt.Println("This is calculator")
	var Number1, Number2 int
	var sign string
	fmt.Scan(&Number1)
	fmt.Scan(&sign)
	fmt.Scan(&Number2)
	if sign == "+" {
		fmt.Println(Add(Number1, Number2))
	} else if sign == "-" {
		fmt.Println(Sub(Number1, Number2))
	} else if sign == "*" {
		fmt.Println(Mul(Number1, Number2))
	} else if sign == "/" {
		fmt.Println(Div(Number1, Number2))
	} else {
		fmt.Println("No operation exist")
	}

}
func Add(Number1, Number2 int) int {
	return Number1 + Number2
}
func Sub(Number1, Number2 int) int {
	return Number1 - Number2
}
func Mul(Number1, Number2 int) int {
	return Number1 * Number2
}
func Div(Number1, Number2 int) int {
	return Number1 / Number2
}
