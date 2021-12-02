package main

import (
	"bufio"
	"fmt"
	"os"
)

type movement struct {
	direction string
	amount    int
}

type position struct {
	depth   int
	forward int
	aim     int
}

func main() {
	var movements []movement
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var direction string
		var amount int
		_, err = fmt.Sscanf(scanner.Text(), "%s %d", &direction, &amount)
		if err != nil {
			panic(err)
		}
		movements = append(movements, movement{direction, amount})
	}

	posA := new(position)
	posB := new(position)
	for _, move := range movements {
		if move.direction == "forward" {
			posA.forward = posA.forward + move.amount
			posB.forward = posB.forward + move.amount
			posB.depth = posB.depth + move.amount*posB.aim
		}
		if move.direction == "up" {
			posA.depth = posA.depth - move.amount
			posB.aim = posB.aim - move.amount
		}
		if move.direction == "down" {
			posA.depth = posA.depth + move.amount
			posB.aim = posB.aim + move.amount
		}
	}
	fmt.Println("Part A:", posA.depth*posA.forward)
	fmt.Println("Part B:", posB.depth*posB.forward)
}
