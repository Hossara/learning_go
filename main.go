package main

import (
	"errors"
	"fmt"
	"log"
)

type User struct {
	firstName string
	lastName  string
}

type car struct {
	// inherit all user fields
	User

	make   string
	model  string
	height string
	width  string
	wheel  struct {
		radius   int
		material string
	}
}

func (car car) detail() string {
	return car.make + " " + car.model + " " + car.height + " " + car.width
}

type UserDetail struct {
}

func sayHello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func sum(x int, y int) (int, int, int) {
	return x, y, x + y
}

func concatName(firstname, lastname string) string {
	return fmt.Sprintf("%s %s", firstname, lastname)
}

func devider(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("divisor is zero")
	}

	return dividend / divisor, nil
}

func main() {
	myCar := car{
		User: User{
			firstName: "James",
		},
		make:   "od",
		model:  "tesla",
		height: "12",
		width:  "12",
		wheel: struct {
			radius   int
			material string
		}{radius: 12, material: "test"},
	}

	fmt.Printf(myCar.detail())
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

	//var userData User = User{firstName, lastName}

	user := User{firstName, lastName}

	log.Println(user)

	if lenght := len(firstName); lenght < 30 {
		fmt.Printf("Hello %d\n", lenght)
	}

	x := []string{
		"hi",
		"hi2",
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}

	for idx, val := range x {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	for _, val := range x {
		fmt.Printf("%v\n", val)
	}

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}

	fmt.Println(myslice2)

	arr1 := [5]int{
		1: 10,
		2: 40,
	}

	fmt.Println(arr1)

	/*
		In the example below myslice is a slice with length 2. It is made from arr1 which is an array with length 6.

		The slice starts from the third element of the array which has value 12 (remember that array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.). The slice can grow to the end of the array. This means that the capacity of the slice is 4.

		If myslice started from element 0, the slice capacity would be 6.
	*/
	arr2 := [6]int{10, 11, 12, 13, 14, 15}
	myslice1 := arr2[2:4]

	fmt.Printf("myslice = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}
