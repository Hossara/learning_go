package main

import (
	"errors"
	"fmt"
)

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func sum(x int, y int) (int, int, int) {
	return x, y, x + y
}

func concatName(firstname, lastname string) string {
	return fmt.Sprintf("%s %s", firstname, lastname)
}

func devide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("divisor is zero")
	}

	return dividend / divisor, nil
}

func main() {
	var age = 12

	firstName, lastName := "Hossein", "Araghi"

	if lenght := len(firstName); lenght < 30 {
		fmt.Printf("Good username\n")
	}

	fmt.Printf("Enter your age: %d \n", age)

	fmt.Printf("Fisrt name is %s and Last name is %s\n", firstName, lastName)

	fmt.Printf("Hi, %s!\n", concatName(firstName, lastName))

	_, _, sum_value := sum(2, 4)

	fmt.Printf("%d\n", sum_value)
}
