package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
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

func getOrientation(val rune) (int, int, int, bool) {
	switch val {
	case '<':
		return 0, -1, 3, true
	case '>':
		return 0, 1, 1, true
	case '^':
		return -1, 0, 0, true
	case 'v':
		return 1, 0, 2, true
	}
	return 0, 0, 'x', false
}

func getStarting(grid [][]rune) (Point, int, int, int, bool) {
	for r, row := range grid {
		for c, col := range row {
			p := Point{Row: r, Col: c}
			i, j, dir, found := getOrientation(col)
			if found {
				return p, i, j, dir, true
			}
		}
	}
	return Point{Row: 0, Col: 0}, 0, 0, 'x', false
}

func checkBounds(grid [][]rune, p Point) bool {
	return p.Row >= 0 && p.Row < len(grid) && p.Col >= 0 && p.Col < len(grid[0])
}

// 1 4
// 1 8
// 6 8
// 6 4
func canMakeBox(grid [][]rune, dir int, points []Point) (Point, bool) {
	if len(points) != 3 {
		fmt.Println("must be 3 points")
		return Point{Row: 0, Col: 0}, false
	}
	rows := map[int]int{}
	cols := map[int]int{}
	for _, p := range points {
		_, ok := rows[p.Row]
		if ok {
			rows[p.Row] += 1
		} else {
			rows[p.Row] = 1
		}

		_, ok = cols[p.Col]
		if ok {
			cols[p.Col] += 1
		} else {
			cols[p.Col] = 1
		}
	}
	needRow, passed := extract(rows)
	if !passed {
		return Point{Row: 0, Col: 0}, false
	}
	needCol, passed := extract(cols)
	if !passed {
		return Point{Row: 0, Col: 0}, false
	}

	inBounds := true
	p := Point{Row: points[2].Row, Col: points[2].Col}
	// _, _, dir = turn90(dir)
	i, j, _ := turn90(dir)
	obstacle := false
	found := false
	for inBounds && !obstacle && !found {
		point := Point{Row: p.Row + i, Col: p.Col + j}
		inBounds = checkBounds(grid, point)
		if needRow == 6 && needCol == 4 {
			fmt.Println("walking", point.Row, point.Col)
		}
		if inBounds && grid[point.Row][point.Col] == '#' {
			obstacle = true
		} else if inBounds && point.Row == needRow && point.Col == needCol {
			found = true
		} else {
			p = point
		}
	}
	if found {
		return Point{Row: needRow, Col: needCol}, true
	}
	return Point{Row: 0, Col: 0}, false
}

func extract(data map[int]int) (int, bool) {
	hasOne := false
	var need int
	hasTwo := false
	for key, val := range data {
		if val == 1 {
			hasOne = true
			need = key
		} else if val == 2 {
			hasTwo = true
		} else {
			break
		}
	}
	return need, hasOne && hasTwo
}

var rotate = []rune{'^', '>', 'v', '<'}

func turn90(dir int) (int, int, int) {
	if dir < 3 {
		dir += 1
	} else {
		dir = 0
	}
	i, j, _, _ := getOrientation(rotate[dir])
	return i, j, dir
}

func part1(filename string) int {
	grid, err := readGrid(filename)
	if err != nil {
		return -1
	}
	p, i, j, dir, found := getStarting(grid)
	if !found {
		return -1
	}

	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	inBounds := true
	for inBounds {
		visited[p.Row][p.Col] = true
		point := Point{Row: p.Row + i, Col: p.Col + j}
		inBounds = checkBounds(grid, point)
		if !inBounds {
			break
		}
		if grid[point.Row][point.Col] == '#' {
			// fmt.Println("turn", p.Row, p.Col, dir)
			i, j, dir = turn90(dir)
		} else {
			p = point
		}
	}

	count := 0
	for _, row := range visited {
		for _, value := range row {
			if value {
				count++
			}
		}
	}
	return count
}

func checkCircular(grid [][]rune, dir int, start Point) bool {
	inBounds := true
	p := Point{Row: start.Row, Col: start.Col}
	i, j, _ := turn90(dir)
	for inBounds {
		point := Point{Row: p.Row + i, Col: p.Col + j}
		inBounds = checkBounds(grid, point)
		if inBounds && grid[point.Row][point.Col] == '#' {
			i, j, dir = turn90(dir)
		} else if inBounds && point.Row == start.Row && point.Col == start.Col {
			return true
		} else {
			p = point
		}
	}
	return false
}

func part2(filename string) {
	grid, err := readGrid(filename)
	if err != nil {
		return
	}
	p, i, j, dir, found := getStarting(grid)
	if !found {
		return
	}
	inBounds := true

	visited := make([][]int, len(grid))
	for i := range visited {
		visited[i] = make([]int, len(grid[0]))
		for j := range visited[i] {
			visited[i][j] = -1
		}
	}

	procs := []DirPoint{}
	for inBounds {
		visited[p.Row][p.Col] = dir
		point := Point{Row: p.Row + i, Col: p.Col + j}
		inBounds = checkBounds(grid, point)
		if inBounds {
			procs = append(procs, DirPoint{point: point, dir: dir})
		}
		if inBounds && grid[point.Row][point.Col] == '#' {
			i, j, dir = turn90(dir)
		} else {
			p = point
		}
	}

	process(grid, procs)
}

type DirPoint struct {
	point Point
	dir   int
}

func worker(grid [][]rune, tasks <-chan DirPoint, results chan<- bool, wg *sync.WaitGroup) {
	for task := range tasks {
		result := checkCircular(grid, task.dir, task.point)
		results <- result
	}
	wg.Done()
}

func process(grid [][]rune, procs []DirPoint) {
	numWorkers := runtime.NumCPU()
	numTasks := len(procs)

	tasks := make(chan DirPoint, numTasks)
	results := make(chan bool, numTasks)

	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(grid, tasks, results, &wg)
	}

	for _, p := range procs {
		tasks <- p
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(results)
	}()

	count := 0
	for result := range results {
		if result {
			count++
		}
	}

	fmt.Printf("Solution: %d\n", count)
}

func main() {
	// example := part2("./input1.txt")
	// if example != 6 {
	// 	fmt.Println("assertion failed :(", example)
	// 	return
	// }
	// fmt.Println("assertion passed!!")
	part2("input2.txt")
}
