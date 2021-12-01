package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numbers []int
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		newnum, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, newnum)
	}

	cur := 0
	resultA := 0
	resultB := 0
	for i, j := range numbers {
		if j > cur {
			resultA++
		}
		cur = j

		if i+3 < len(numbers) {
			if numbers[i]+numbers[i+1]+numbers[i+2] < numbers[i+1]+numbers[i+2]+numbers[i+3] {
				resultB++
			}
		}

	}
	fmt.Println("Part 1:", resultA-1)
	fmt.Println("Part 2:", resultB)

}
