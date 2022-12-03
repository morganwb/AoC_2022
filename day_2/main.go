package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "sort"
)

func main() {
	var totalScore int
	var tempScore int
	var pt2 bool
	var opponentChoice, ourChoice int
	var scores = map[string]int {
		"A": 1,
		"B": 2,
		"C": 3,
		"X": 1,
		"Y": 2,
		"Z": 3,
	}
	var outcomes = map[string]int {
		"A Z": 0,
		"A Y": 6,
		"B Z": 6,
		"B X": 0,
		"C X": 6,
		"C Y": 0,
	}
	pt2 = true

	totalScore = 0
	f, _ := os.Open("day2_inputs.txt")
	scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		line := scanner.Text()
		opponentChoice = scores[string(line[0])]
		// Now we can make our own choice for outcome!
		ourChoice = scores[string(line[2])]
		fmt.Println("---------------------------------------------")
		if pt2 == true {
			if ourChoice == 2 {
				fmt.Println("we want to tie!")
				ourChoice = opponentChoice
				tempScore = 3
			} else if ourChoice == 1 {
				fmt.Println("we want to lose!")
				for key, val := range outcomes {
					if (string(key[0]) == string(line[0]) && val == 0) {
						fmt.Println("they picked:")
						fmt.Println(string(key[0]))
						fmt.Println("so we should pick:")
						fmt.Println(string(key[2]))
						ourChoice = scores[string(key[2])]
						tempScore = val
					}
				}
			} else if ourChoice == 3{
				fmt.Println("we want to win!")
				for key, val := range outcomes {
					if (string(key[0]) == string(line[0]) && val == 6) {
						fmt.Println("they picked:")
						fmt.Println(string(key[0]))
						fmt.Println("so we should pick:")
						fmt.Println(string(key[2]))
						ourChoice = scores[string(key[2])]
						tempScore = val
					}
				}
			}
		}
		if pt2 == false {
			if ourChoice == opponentChoice {
				tempScore = 3
			} else {
				tempScore = outcomes[line]
			}
		}
		fmt.Println(line)
		fmt.Println(ourChoice)
		fmt.Println(tempScore)
		totalScore = totalScore + tempScore + ourChoice
		fmt.Println(totalScore)
		tempScore = 0
    }
	fmt.Println(totalScore)
}