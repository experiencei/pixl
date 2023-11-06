package main

import "fmt"

func main() {

	var myName = "Ibrahim"
	fmt.Println("my name is" , myName)

	var nameOfMe string = "Semiat"
	fmt.println("name" , nameOfMe)

	userName := "experiencei"
	fmt.Println("username is" , userName)

	var sum int
	fmt.println("the sum of null define is" , sum)

	part1 , others := 1 , 4
	fmt.println("part1 of other is" , part1 , "other of part1 is" , others) 

	part2 , others := 1 , 4
	fmt.println("part2 of other is" , part2 , "other of part2 is" , others) 

	sum = part1 + part2
	fmt.println("the sum of part1 and part2 is" , sum)

  var (
		lessonName = "Variables"
		lessonType = "Demo"
	)	
  fmt.println("The lesson Name is" , lessonName)
  fmt.println("The lesson Type is" , lessonType)

	word1, word2, _ := "hello", "there", "!"
	fmt.Println(word1, word2)
}


