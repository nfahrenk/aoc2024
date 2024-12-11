package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	Row int
	Col int
}

func readInput(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	return line, nil
}

func buildLine(filename string) ([]int, []Block, []Block, error) {
	line, err := readInput(filename)
	if err != nil {
		fmt.Println("could not read input", err)
		return nil, nil, nil, err
	}

	free := []Block{}
	values := []Block{}

	id := 0
	updated := []int{}
	for ndx, char := range line {
		add := -1
		if ndx%2 == 0 {
			add = id
			id++
		}
		maxLen, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Println("could not parse int", err)
			return nil, nil, nil, err
		}

		startNdx := len(updated)
		b := Block{startNdx: startNdx, id: add, size: maxLen}
		if add == -1 {
			free = append(free, b)
		} else {
			values = append(values, b)
		}
		for i := 0; i < maxLen; i++ {
			updated = append(updated, add)
		}
	}
	return updated, values, free, nil

}

func countLine(updated []int) int {
	counter := 0
	for i, val := range updated {
		if val > 0 {
			counter += i * val
		}
	}
	return counter
}

func part1(filename string) int {
	updated, _, _, err := buildLine(filename)
	if err != nil {
		return -1
	}
	freePtr := 0
	valPtr := len(updated) - 1
	for freePtr < valPtr {
		if updated[freePtr] >= 0 {
			freePtr++
		} else if updated[valPtr] < 0 {
			valPtr--
		} else if updated[valPtr] >= 0 && updated[freePtr] < 0 {
			updated[freePtr] = updated[valPtr]
			updated[valPtr] = -1
			valPtr--
			freePtr++
		} else {
			fmt.Println("no operation happened")
		}
	}
	fmt.Println("updated", updated)
	return countLine(updated)
}

type Block struct {
	startNdx int
	size     int
	id       int
}

func part2(filename string) int {
	updated, values, free, err := buildLine(filename)
	if err != nil {
		return -1
	}
	fmt.Println(updated)
	fmt.Println()
	for i := len(values) - 1; i >= 0; i-- {
		valSize := values[i].size
		for ndx, space := range free {
			if space.startNdx > values[i].startNdx {
				break
			}
			if valSize <= space.size {
				for j := 0; j < valSize; j++ {
					updated[space.startNdx+j] = values[i].id
					updated[values[i].startNdx+j] = -1
				}
				free[ndx].size -= valSize
				free[ndx].startNdx += valSize
				break
			}
		}
	}
	fmt.Println()
	fmt.Println(updated)
	return countLine(updated)
}

func main() {
	example := part2("./input1.txt")
	if example != 2858 {
		fmt.Println("assertion failed :(", example)
		return
	}
	fmt.Println("assertion passed!!")
	output := part2("./input2.txt")
	fmt.Println("output", output)
}
