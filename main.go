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

	NormalEchoWords := []string{}

	for i, word := range os.Args[1:] {
		if len(os.Args) <= 1 {
			fmt.Println("error: ", "No input provided")
		}

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
			return
		} else if word == prefix1 {
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
			return
		} else {
			NormalEchoWords = append(NormalEchoWords, word)
		}
	}
	completeNormalEcho := strings.Join(NormalEchoWords, " ")
	fmt.Println(completeNormalEcho)
}
func main() {
	readInput()
}
