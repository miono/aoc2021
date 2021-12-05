package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var numbers []string

	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}

	fmt.Println("Part A:", partA(numbers))
	fmt.Println("Part B:", partB(numbers))
}

func getStringFromBitSums(length int, bitSums map[int]int) string {
	var outString string
	for i := 0; i < len(bitSums); i++ {
		if bitSums[i] > (length / 2) {
			outString = outString + "1"
		} else {
			outString = outString + "0"
		}
	}
	return outString
}
func inverseBitString(bitString string) string {
	var outString string
	for _, k := range bitString {
		if k == '1' {
			outString = outString + "0"
		} else {
			outString = outString + "1"
		}
	}
	return outString
}

func multiplyBitStrings(stringA, stringB string) int {
	a, _ := strconv.ParseInt(stringA, 2, 0)
	b, _ := strconv.ParseInt(stringB, 2, 0)
	return int(a * b)
}

func partA(numbers []string) int {
	bitSums := make(map[int]int)
	totalLength := len(numbers)
	for _, j := range numbers {
		for k, l := range j {
			if l == '1' {
				bitSums[k]++
			}
		}
	}
	bitString := getStringFromBitSums(totalLength, bitSums)

	return multiplyBitStrings(bitString, inverseBitString(bitString))
}

func findNumber(pos int, searchMode string, numbers []string) string {
	if len(numbers) == 1 {
		return numbers[0]
	} else {
		// Find most common bit
		var commonBit byte
		var matching int
		var nonMatching int
		if searchMode == "oxygen" {
			for _, j := range numbers {
				if j[pos] == '1' {
					matching++
				} else {
					nonMatching++
				}
			}
			if matching >= nonMatching {
				commonBit = '1'
			} else {
				commonBit = '0'
			}
		}
		if searchMode == "co2" {
			for _, j := range numbers {
				if j[pos] == '0' {
					matching++
				} else {
					nonMatching++
				}
			}
			if matching <= nonMatching {
				commonBit = '0'
			} else {
				commonBit = '1'
			}
		}
		// Create new numbers-slice
		var newNumberSlice []string
		for _, j := range numbers {
			if j[pos] == commonBit {
				newNumberSlice = append(newNumberSlice, j)
			}
		}
		return findNumber(pos+1, searchMode, newNumberSlice)
	}
}

func partB(numbers []string) int {
	oxygenGeneratorRating := findNumber(0, "oxygen", numbers)
	co2ScrubberRating := findNumber(0, "co2", numbers)

	return multiplyBitStrings(oxygenGeneratorRating, co2ScrubberRating)
}
