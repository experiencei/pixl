package main

import "fmt"


func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1
	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println(shoppingList)
}