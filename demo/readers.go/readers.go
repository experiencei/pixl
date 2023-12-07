package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	sum := 0

	for {
		input, inputErr := scanner.ReadString(' ')

		num, convErr := strconv.Atoi(strings.TrimSpace(input))
		if convErr != nil {
			fmt.Println(convErr)
		} else {
			sum += num
		}

		