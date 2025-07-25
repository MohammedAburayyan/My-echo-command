package main

import (
	"fmt"
	"os"
	"strings"
)

func readInput() {

	if len(os.Args) <= 1 {
		fmt.Println("error: ", "No input provided")
	}
	restOfInput := strings.Join(os.Args[1:], " ")
	fmt.Println(restOfInput)

}

func main() {
	readInput()
}
