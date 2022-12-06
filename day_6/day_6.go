package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var commSlice []string
	part2 := map[string]struct{}{}
	rmdStr := "null"
	fmt.Println(rmdStr)
	f, _ := os.Open("inputs_day6.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		commSlice = strings.Split(line, "")
	}
	// part 1
	for i := 3; i < len(commSlice); i++ {
		if commSlice[i-3] != commSlice[i-2] && commSlice[i-3] != commSlice[i-1] && commSlice[i-3] != commSlice[i] && commSlice[i-2] != commSlice[i-1] && commSlice[i-2] != commSlice[i] && commSlice[i-1] != commSlice[i] {
			fmt.Printf("%s", commSlice[i-3])
			fmt.Printf("%s", commSlice[i-2])
			fmt.Printf("%s", commSlice[i-1])
			fmt.Printf("%s", commSlice[i])
			fmt.Println()
			fmt.Println("First marker found at:")
			fmt.Println(i + 1)
			i = len(commSlice)
		}
	}

	// Solve part 2
	for i := 0; i < len(commSlice)-14; i++ {
		for j := 0; j < 14; j++ {
			part2[commSlice[i+j]] = struct{}{}
		}
		uniqueCount := 0
		for key, _ := range part2 {
			uniqueCount = uniqueCount + 1
			rmdStr = key
		}
		if uniqueCount == 14 {
			fmt.Println("WE MADE IT!")
			fmt.Println(i + 14)
			i = len(commSlice)
		} else {
			for key, _ := range part2 {
				delete(part2, key)
			}
		}

	}
}
