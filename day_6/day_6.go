package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var commSlice []string
	// var checkSlice string
	f, _ := os.Open("inputs_day6.txt")
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		commSlice = strings.Split(line, "")
	}
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
}
