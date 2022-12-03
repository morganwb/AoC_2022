package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "sort"
)

func main() {
	var alphaBetVals = make(map[string]int)
	var i, j, k, strLen int
	var uniqueStr, firstSack, secondSack string
	var finalResult int = 0
	var lineStorage [3]string
	var SRstorage int
	var secondResult int = 0
	var lineCounter int = 0
	alphaBet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i = 0; i<len(alphaBet); i++ {
		alphaBetVals[string(alphaBet[i])] = i+1
	}
	f, _ := os.Open("inputs_day3.txt")
	scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		line := scanner.Text()
		strLen = len(line)
		firstSack = string(line[0:strLen/2])
		secondSack = string(line[strLen/2:strLen])
		for i = 0; i<strLen/2; i++ {
			for j = 0; j<strLen/2; j++ {
				if firstSack[i] == secondSack[j]{
					uniqueStr = string(firstSack[i])
				}
			}
		}
		finalResult = finalResult + alphaBetVals[uniqueStr]
		lineCounter = lineCounter + 1
		// fmt.Println(lineCounter % 3)
		lineStorage[lineCounter % 3] = line
		if (lineCounter)%3 == 0 {
			fmt.Println("full! contents are:")
			fmt.Println(lineStorage)
			for i = 0; i<len(lineStorage[0]); i++ {
				for j = 0; j<len(lineStorage[1]); j++ {
					for k = 0; k<len(lineStorage[2]); k++ {
						if lineStorage[0][i] == lineStorage[1][j] && lineStorage[1][j] == lineStorage[2][k]{
							SRstorage = alphaBetVals[string(lineStorage[1][j])]
							fmt.Println(string(lineStorage[1][j]))
						}
					}
				}
			}
			fmt.Println(secondResult)
			fmt.Println(SRstorage)
			secondResult = secondResult + SRstorage
		}
		
    }
	fmt.Println(finalResult)
	fmt.Println(secondResult)
}