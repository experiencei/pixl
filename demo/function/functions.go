package main

import (
	"fmt"
	"strings"
)

func greet() {
	fmt.Println("Hello from greet function")
}


func add(lhs , rhs int) int  {
	return lhs + rhs
}


func double(x int) int  {
	return x + x
}


// Variadic functions are functions that take a variable number of arguments.
// Ellipsis prefix (three-dots) in front of the parameter type makes a function variadic.
// The function may be called with zero or more arguments for that parameter.
// If the function takes parameters of different types, only the last parameter of a function can be variadic.


// creating a variadic function
func f1(a ...int) {
    fmt.Printf("%T\n", a) // => []int, slice of int
    fmt.Printf("%#v\n", a)
}

// variadicfunction that modifies one of the arguments passed.
func f2(a ...int) {
    a[0] = 50
}

// creating a variadic function that calculates and returns the sum and product of its arguments
func sumAndProduct(a ...float64) (float64, float64) {
    sum := 0.
    product := 1.

    for _, v := range a {
        sum += v
        product *= v
    }

    return sum, product
}


// mixing variadic and non-variadic parameters is allowed
// non-variadic parameters are always before the variadic parameter
func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
	return returnString
}

func main() {
	greet()
 
	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Have another", anotherBakersDozen)

}

