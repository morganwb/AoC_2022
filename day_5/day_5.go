package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
	// "strconv"
	// "sort"
)

func main() {
	// var part1Result int
	// part1Result = 0
	// var part2Result int
	var alphaBetVals = make(map[string]int)
	var i,j int
	var cargoShip = make(map[string][]string)
	var loadInitial bool = true
	cargoShipRead := make([]string, 0)
	// var stackAssembler []string
	lineCounter := 0
	alphaBet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i = 0; i<len(alphaBet); i++ {
		alphaBetVals[string(alphaBet[i])] = i+1
	}
	f, _ := os.Open("inputs_day5.txt")
	scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		line := scanner.Text()
		if loadInitial{
			cargoShipRead = append(cargoShipRead, line)
			lineCounter = lineCounter + 1
			if string(line[1]) == string('1'){
				lineLength := len(line)
				for i = 0; i<lineLength; i++ {
					if string(line[i]) != " "{
						index := string(line[i])
						for j = len(cargoShipRead)-1; j>=0; j-- {
							if string(cargoShipRead[j][i]) != " " {
								cargoShip[index] = append(cargoShip[index], string(cargoShipRead[j][i]))
							}
						}
						fmt.Println("============================")
						fmt.Println("here are the entries for:")
						fmt.Println(index)
						fmt.Println(cargoShip[index])
					}
				}
				loadInitial = false
			}
		}
	}
	// fmt.Println(cargoShip)
	// fmt.Println(part2Result)
}