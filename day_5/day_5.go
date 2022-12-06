package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var alphaBetVals = make(map[string]int)
	var i, j, k int
	var cS = make(map[string][]string)
	var cS9 = make(map[string][]string)
	var loadInitial bool = true
	cSRead := make([]string, 0)
	lineCounter := 0
	alphaBet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i = 0; i < len(alphaBet); i++ {
		alphaBetVals[string(alphaBet[i])] = i + 1
	}
	f, _ := os.Open("inputs_day5.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if loadInitial {
			cSRead = append(cSRead, line)
			lineCounter = lineCounter + 1
			if string(line[1]) == string('1') {
				lineLength := len(line)
				for i = 0; i < lineLength; i++ {
					if string(line[i]) != " " {
						index := string(line[i])
						for j = len(cSRead) - 1; j >= 0; j-- {
							if string(cSRead[j][i]) != " " {
								cS[index] = append(cS[index], string(cSRead[j][i]))
								cS9[index] = append(cS9[index], string(cSRead[j][i]))
							}
						}
						cS[index] = cS[index][1:len(cS[index])]
						cS9[index] = cS9[index][1:len(cS9[index])]
					}
				}
				loadInitial = false
			}
		} else if len(line) > 0 {
			cmdProc := strings.Split(line, " ")
			var moveCount, _ = strconv.Atoi(cmdProc[1])
			moveFrom := cmdProc[3]
			moveTo := cmdProc[5]
			for i = 0; i < moveCount; i++ {
				cS[moveTo] = append(cS[moveTo], cS[moveFrom][len(cS[moveFrom])-1])
				cS[moveFrom] = cS[moveFrom][0 : len(cS[moveFrom])-1]
			}
			cS9mF_len := len(cS9[moveFrom])
			bottomEnd := cS9mF_len - moveCount
			cS9[moveTo] = append(cS9[moveTo], cS9[moveFrom][bottomEnd:cS9mF_len]...)
			cS9[moveFrom] = cS9[moveFrom][0:bottomEnd]
		}
	}
	// I'm just gonna say it, the keys in key-value pair objects should be stored
	// in the order in which they were added, so when the keys are returned
	// by a generator-ish routine, they reappear in the order they were added.
	// Why isn't this a thing?
	keyCount := 1
	for key, _ := range cS {
		keyCount = keyCount + 1
		key = ""
		fmt.Printf(key)
	}
	fmt.Println("")
	for k = 1; k < keyCount; k++ {
		fmt.Printf(cS[strconv.Itoa(k)][len(cS[strconv.Itoa(k)])-1])
	}
	fmt.Printf("\n")
	fmt.Println("FINAL cS SOLUTION:")
	fmt.Println(cS)

	fmt.Println("#####################################################################")
	keyCount = 1
	for key, _ := range cS9 {
		keyCount = keyCount + 1
		key = ""
		fmt.Printf(key)
	}
	fmt.Println("")
	for k = 1; k < keyCount; k++ {
		fmt.Printf(cS9[strconv.Itoa(k)][len(cS9[strconv.Itoa(k)])-1])
	}
	fmt.Printf("\n")
	fmt.Println("FINAL cS9 SOLUTION:")
	fmt.Println(cS9)
}
