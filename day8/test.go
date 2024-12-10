package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	Row int
	Col int
}

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

func checkBounds(grid [][]rune, p Point) bool {
	return p.Row >= 0 && p.Row < len(grid) && p.Col >= 0 && p.Col < len(grid[0])
}

func generatePairs(options []Point, current []Point, results *[][]Point) {
	if len(current) == 2 {
		*results = append(*results, current)
		return
	}
	for _, opt := range options {
		// deep copy
		newList := make([]Point, 0)
		newList = append(newList, current...)
		newList = append(newList, opt)

		generatePairs(options, newList, results)
	}
}

func buildFreqs(grid [][]rune) map[rune][]Point {
	freqs := map[rune][]Point{}
	for r := range grid {
		for c := range grid[r] {
			freq := grid[r][c]
			if freq != '.' {
				_, ok := freqs[freq]
				if !ok {
					freqs[freq] = make([]Point, 0)
				}
				freqs[freq] = append(freqs[freq], Point{Row: r, Col: c})
			}
		}
	}
	return freqs
}

func part1(filename string) int {
	grid, err := readGrid(filename)
	if err != nil {
		return 0
	}
	count := 0

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	freqs := buildFreqs(grid)
	for _, points := range freqs {
		results := make([][]Point, 0)
		generatePairs(points, make([]Point, 0), &results)
		for _, pairs := range results {
			// p1 3, 4; p2 5, 6
			// n1 1, 3; n2 7, 8
			// 6 + (6 - 4), 4 + (4 - 6)
			prevP := pairs[0]
			p := pairs[1]
			// don't process duplicates
			if prevP.Row == p.Row && prevP.Col == p.Col {
				continue
			}
			n1, _, _ := diff(prevP, p)
			if checkBounds(grid, n1) && !visited[n1.Row][n1.Col] {
				visited[n1.Row][n1.Col] = true
				count += 1
			}
			n2, _, _ := diff(p, prevP)
			if checkBounds(grid, n2) && !visited[n2.Row][n2.Col] {
				visited[n2.Row][n2.Col] = true
				count += 1
			}
		}
	}
	for _, row := range visited {
		line := []string{}
		for _, val := range row {
			if val {
				line = append(line, "#")
			} else {
				line = append(line, ".")

			}
		}
		fmt.Println(line)
	}
	return count
}

func part2(filename string) int {
	grid, err := readGrid(filename)
	if err != nil {
		return 0
	}
	count := 0

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	freqs := buildFreqs(grid)
	for _, points := range freqs {
		for _, point := range points {
			if !visited[point.Row][point.Col] {
				visited[point.Row][point.Col] = true
				count += 1

			}
		}
		results := make([][]Point, 0)
		generatePairs(points, make([]Point, 0), &results)
		for _, pairs := range results {
			prevP := pairs[0]
			p := pairs[1]
			// don't process duplicates
			if prevP.Row == p.Row && prevP.Col == p.Col {
				continue
			}

			n1, A, B := diff(prevP, p)
			inBounds := true
			for inBounds {
				inBounds = checkBounds(grid, n1)
				if inBounds && !visited[n1.Row][n1.Col] {
					visited[n1.Row][n1.Col] = true
					count += 1
				}
				n1 = Point{Row: n1.Row + A, Col: n1.Col + B}
			}
			n2, A, B := diff(p, prevP)
			inBounds = true
			for inBounds {
				inBounds = checkBounds(grid, n2)
				if inBounds && !visited[n2.Row][n2.Col] {
					visited[n2.Row][n2.Col] = true
					count += 1
				}
				n2 = Point{Row: n2.Row + A, Col: n2.Col + B}
			}
		}
	}
	for _, row := range visited {
		line := []string{}
		for _, val := range row {
			if val {
				line = append(line, "#")
			} else {
				line = append(line, ".")

			}
		}
		fmt.Println(line)
	}
	return count
}

func diff(p1 Point, p2 Point) (Point, int, int) {
	B := (p1.Col - p2.Col)
	A := (p1.Row - p2.Row)
	c := p1.Col + B
	r := p1.Row + A
	return Point{Row: r, Col: c}, A, B
}

func main() {
	example := part2("./input1.txt")
	if example != 34 {
		fmt.Println("assertion failed :(", example)
		return
	}
	fmt.Println("assertion passed!!")
	output := part2("./input2.txt")
	fmt.Println("output", output)
}
