//--Summary:
//  Create a program that can perform dice rolls. The program should
//  report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must handle any number of dice, rolls, and sides
//
//--Notes:
//* Use packages from the standard library to complete the project

package main

import (
	"fmt"
	"math/rand"
	"time"
)
 
func roll(sides int) int {
	return rand.Intn(sides) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//* The program must handle any number of dice, rolls, and sides
	dice, sides := 2, 12
	rolls := 1

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, ", die #", d, ":", rolled)
		}
		//* Print the sum of the dice roll
		fmt.Println("Total rolled:", sum)
		evenNums := sum % 2 == 0;
	  oddNums :=  sum % 2 == 1;
	
		switch sum := sum; {
		case sum == 2 && dice == 2:
			//  - "Snake eyes": when the total roll is 2, and total dice is 2
			fmt.Println("Snake Eyes!")
		case sum == 7:
			//  - "Lucky 7": when the total roll is 7
			fmt.Println("Lucky Seven!")
		case evenNums:
			//  - "Even": when the total roll is even
			fmt.Println("Even")
		case oddNums:
			//  - "Odd": when the total roll is odd
			fmt.Println("Odd")
		}
	}
}
