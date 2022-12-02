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


		if pt2 == true {
			




		}



		if ourChoice == opponentChoice {
			tempScore = 3
		} else {
			tempScore = outcomes[line]
		}
		fmt.Println("---------------------------------------------")
		fmt.Println(line)
		fmt.Println(ourChoice)
		fmt.Println(tempScore)
		totalScore = totalScore + tempScore + ourChoice
		fmt.Println(totalScore)
		tempScore = 0
    }
	fmt.Println(totalScore)
}