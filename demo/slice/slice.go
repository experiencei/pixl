package main

import "fmt"


func printSlice(title string , slice []string) {
	fmt.Println()
	fmt.Println("---" , title , "---")

	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)

	}
}


func main() {

}