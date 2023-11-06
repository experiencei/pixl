package main

import "fmt"

func greet() {
	fmt.println("Hello from greet function")
}


func add(lhs , rhs int) int  {
	return lhs + rhs
}


func double(x int) int  {
	return x + x
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

