package main

import (
	"bufio"
	"fmt"
	"os"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type point struct {
	x int
	y int
}

type grid struct {
	points map[int]map[int]int
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func getPointsFromLine(line line) []point {
	var output []point
	var startX, endX, startY, endY int
	if line.x1 < line.x2 {
		startX = line.x1
		endX = line.x2
	} else if line.x1 > line.x2 {
		startX = line.x2
		endX = line.x1
	} else {
		startX = line.x1
		endX = line.x1
	}
	if line.y1 < line.y2 {
		startY = line.y1
		endY = line.y2
	} else if line.y1 > line.y2 {
		startY = line.y2
		endY = line.y1
	} else {
		startY = line.y1
		endY = line.y1
	}
	for i := startX; i <= endX; i++ {
		for j := startY; j <= endY; j++ {
			output = append(output, point{i, j})
		}
	}
	return output
}

func getPointsFromDiagLine(line line) []point {
	var output []point
	if line.x1 < line.x2 {
		if line.y1 < line.y2 {
			y := line.y1
			for i := line.x1; i <= line.x2; i++ {
				output = append(output, point{i, y})
				y++
			}
		} else {
			y := line.y1
			for i := line.x1; i <= line.x2; i++ {
				output = append(output, point{i, y})
				y--
			}

		}
	} else if line.x2 < line.x1 {
		if line.y1 < line.y2 {
			y := line.y2
			for i := line.x2; i <= line.x1; i++ {
				output = append(output, point{i, y})
				y--
			}
		} else {
			y := line.y2
			for i := line.x2; i <= line.x1; i++ {
				output = append(output, point{i, y})
				y++
			}

		}
	}
	return output
}

func (grid *grid) drawLine(points []point) {
	for _, point := range points {
		grid.points[point.x][point.y]++
	}

}

func newGrid(size int) grid {
	points := map[int]map[int]int{}
	for i := 0; i < size; i++ {
		points[i] = map[int]int{}
	}
	return grid{points: points}
}

func main() {
	gridA := newGrid(1000)
	gridB := newGrid(1000)
	var lines []line
	f, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var x1, y1, x2, y2 int
		_, err = fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		lines = append(lines, line{x1, y1, x2, y2})
	}
	for _, line := range lines {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			gridA.drawLine(getPointsFromLine(line))
			gridB.drawLine(getPointsFromLine(line))
		} else if abs(line.y1-line.y2) == abs(line.x1-line.x2) {
			gridB.drawLine(getPointsFromDiagLine(line))
		}
	}
	var numMoreThanOneA int
	var numMoreThanOneB int
	for _, line := range gridA.points {
		for _, value := range line {
			if value > 1 {
				numMoreThanOneA++
			}
		}
	}
	for _, line := range gridB.points {
		for _, value := range line {
			if value > 1 {
				numMoreThanOneB++
			}
		}
	}
	fmt.Println("Part A:", numMoreThanOneA)
	fmt.Println("Part B:", numMoreThanOneB)
}
