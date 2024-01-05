package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}


// Must specify the number of array elements in the function parameters
func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is dirty")
		}
	}
}


func main() {
    // the ellipsis operator (...) finds out automatically the length of the array
	rooms := [...]Room{
		{name: "Office"},
		{name: "Warehouse"},
		{name: "Reception"},
		{name: "Ops"},
	}

	checkCleanliness(rooms)

	fmt.Println("Performing cleaning...")
	// Elements start at index 0
	rooms[2].cleaned = true // element 3
	rooms[3].cleaned = true // element 4

	checkCleanliness(rooms)


	//creating multi dimensional array

	carMatrix := [2][3]int{
		{1 , 22 },
		{0 , 4},
	}

	_ = carMatrix

	// using keyed and unkeyed element

	weekend := [7]bool{
		5 : true,
		6 : true,
	}

	fmt.Println("The weekend is only true" , weekend)

	// un unkeyed element gets its index from the last keyed element
	cities := [...]string{
		5:        "Paris",
		"London", // this is at index 6
		1:        "NYC",
}
 _ = cities
}