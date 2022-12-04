package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	// "sort"
)

func main() {
	var part1Result int
	var part2Result int
	var aB [4]int
	var err [4]error
	f, _ := os.Open("inputs_day4.txt")
	scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, ",")
		pair1 := strings.Split(pair[0], "-")
		pair2 := strings.Split(pair[1], "-")
		aB[0], err[0] = strconv.Atoi(pair1[0])
		aB[1], err[1] = strconv.Atoi(pair1[1])
		aB[2], err[2] = strconv.Atoi(pair2[0])
		aB[3], err[3] = strconv.Atoi(pair2[1])
		fmt.Println(aB)
		// Check if sets are supersets
		if (aB[0] <= aB[2] && aB[1] >= aB[3]) || (aB[2] <= aB[0] && aB[3] >= aB[1]) {
			part1Result = part1Result + 1
			fmt.Println("super set!")
		}
		fmt.Println(aB[0] <= aB[3] && aB[0] >= aB[2])
		fmt.Println(aB[1] <= aB[3] && aB[1] >= aB[2])
		fmt.Println(aB[2] <= aB[1] && aB[2] >= aB[0])
		fmt.Println(aB[3] <= aB[1] && aB[3] >= aB[0])
		if (aB[0] <= aB[3] && aB[0] >= aB[2]) || (aB[1] <= aB[3] && aB[1] >= aB[2]) ||
		(aB[2] <= aB[1] && aB[2] >= aB[0]) || (aB[3] <= aB[1] && aB[3] >= aB[0]) {
			part2Result = part2Result + 1
			fmt.Println("overlap!")
		}
	}
	fmt.Println(part1Result)
	fmt.Println(part2Result)
}
