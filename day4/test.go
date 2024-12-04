package main

import (
	"bufio"
	"fmt"
	"os"
)

func readGrid(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := [][]rune{}

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	return grid, nil
}

type Point struct {
	Row int
	Col int
}

func checkPoint(grid [][]rune, p Point, search rune) bool {
	return grid[p.Row][p.Col] == search
}

func checkBounds(grid [][]rune, p Point) bool {
	return p.Row >= 0 && p.Row < len(grid) && p.Col >= 0 && p.Col < len(grid[0])
}

func checkLine(grid [][]rune, initial Point, i int, j int) bool {
	p := Point{Row: initial.Row, Col: initial.Col}
	for _, letter := range []rune{'M', 'A', 'S'} {
		p.Row += i
		p.Col += j
		if checkBounds(grid, p) {
			if !checkPoint(grid, p, letter) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func checkX(grid [][]rune, initial Point) int {
	output := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if checkLine(grid, initial, i, j) {
				output++
			}
		}
	}
	return output
}

func part1(filename string) int {
	grid, err := readGrid(filename)
	if err != nil {
		return -1
	}
	counter := 0
	for r, row := range grid {
		for c, val := range row {
			if val == 'X' {
				p := Point{Row: r, Col: c}
				counter += checkX(grid, p)
			}
		}
	}
	return counter
}

func checkA(grid [][]rune, initial Point) bool {
	topLeft := Point{Col: initial.Col - 1, Row: initial.Row - 1}
	topRight := Point{Col: initial.Col - 1, Row: initial.Row + 1}
	bottomLeft := Point{Col: initial.Col + 1, Row: initial.Row - 1}
	bottomRight := Point{Col: initial.Col + 1, Row: initial.Row + 1}
	if !checkBounds(grid, topLeft) || !checkBounds(grid, topRight) || !checkBounds(grid, bottomLeft) || !checkBounds(grid, bottomRight) {
		return false
	}
	if !((grid[topLeft.Row][topLeft.Col] == 'S' &&
		grid[bottomRight.Row][bottomRight.Col] == 'M') ||
		(grid[topLeft.Row][topLeft.Col] == 'M' &&
			grid[bottomRight.Row][bottomRight.Col] == 'S')) {
		return false
	}
	if !((grid[topRight.Row][topRight.Col] == 'S' &&
		grid[bottomLeft.Row][bottomLeft.Col] == 'M') ||
		(grid[topRight.Row][topRight.Col] == 'M' &&
			grid[bottomLeft.Row][bottomLeft.Col] == 'S')) {
		return false
	}
	return true
}

func part2(filename string) int {
	grid, err := readGrid(filename)
	if err != nil {
		return -1
	}
	counter := 0
	for r, row := range grid {
		for c, val := range row {
			if val == 'A' {
				p := Point{Row: r, Col: c}
				if checkA(grid, p) {
					counter++
				}
			}
		}
	}
	return counter
}

func main() {
	example := part2("input2.txt")
	if example != 9 {
		fmt.Println("assertion failed :(", example)
		return
	}
	fmt.Println("assertion passed!!")
	output := part2("input1.txt")
	fmt.Println("output", output)
}
