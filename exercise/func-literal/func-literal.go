//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import "fmt"
import "unicode"

type LineCallback func(line string)

//* Create a single function to iterate over each line of text that is
//  provided in main().
func lineIterator(lines []string, callback LineCallback) {
	for i := 0; i < len(lines); i++ {
		//  - The function must return nothing and must execute a closure
		callback(lines[i])
	}
}