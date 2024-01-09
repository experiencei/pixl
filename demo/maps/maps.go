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

	fmt.Println("need", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]
	fmt.Println("Need cereal?")

	if !found {
		fmt.Println("nope")
	} else {
		fmt.Println("yup:", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}
	fmt.Println("There are", totalItems, "on the shopping list")


	//** COMPARING MAPS **//

    // Maps cannot be compared using == operator. A map can be compared only to nil.
    a := map[string]string{"A": "X"}
    b := map[string]string{"B": "X"}

    // fmt.Println(a == b) // error -> invalid operation: a == b (map can only be compared to nil)

    // to compare 2 maps that have the keys and values of type string
    // we get a string representation of the maps and compare those strings.

    // getting a string representation of maps called a and b
    s1 := fmt.Sprintf("%s", a)
    s2 := fmt.Sprintf("%s", b)

    if s1 == s2 {
        fmt.Println("Maps are equal")
    } else {
        fmt.Println("Maps are not equal")
    }
}