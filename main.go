package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const prefix1 string = "-t"

const prefix2 string = "-r"

func readInput() {

	reverse := false

	for i, word := range os.Args[1:] {
		if len(os.Args) <= 1 {
			fmt.Println("error: ", "No input provided")
		}
		// NOT WORKING
		// if word != prefix1 && word != prefix2 {
		// 	fmt.Println(os.Args[1:])
		// 	return
		// }

		if word == prefix2 {
			reverse = true
		}

		if reverse {
			completeString := strings.Join(os.Args[1:i+1], " ")
			runeList := []rune(completeString)
			pointer1 := 0
			pointer2 := len(runeList)

			for pointer1 < pointer2 {
				pointer2--
				runeList[pointer1], runeList[pointer2] = runeList[pointer2], runeList[pointer1]
				pointer1++
			}
			fmt.Println(string(runeList))
		}

		if word == prefix1 {
			completeString := strings.Join(os.Args[1:i+1], " ")
			convertString := os.Args[i+2]
			timesRepeated, err := strconv.Atoi(convertString)
			if err != nil {
				fmt.Println("error: ", "Invalid number for repetition")
				return
			} else {
				for i := 0; i < timesRepeated; i++ {
					fmt.Println(completeString)
				}
			}
		}
	}
}

func main() {
	readInput()
}
