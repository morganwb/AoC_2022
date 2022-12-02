package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

func main() {
	var currentVal int
	var calPerElf []int
	// var sortedElves []int
	f, _ := os.Open("day1_inputs.txt")
	scanner := bufio.NewScanner(f)
	var oldHighest int
    for scanner.Scan() {
      line := scanner.Text()
	  if (line == "") {
		if currentVal > oldHighest {
			oldHighest = currentVal
		}
		calPerElf = append(calPerElf, currentVal)
		currentVal = 0
	  } else{
		strInt, err := strconv.Atoi(line)
		if err != nil{
			panic(err)
		}
		currentVal = currentVal + strInt
	  }
    }
	fmt.Println("Elf with the highest calorie count:")
	fmt.Println(oldHighest)
	fmt.Println("Calorie total of the highest three elves:")
	sort.Slice(calPerElf, func(p, q int) bool { 
		return calPerElf[p] > calPerElf[q] })
	fmt.Println(calPerElf[0] + calPerElf[1] + calPerElf[2])

}