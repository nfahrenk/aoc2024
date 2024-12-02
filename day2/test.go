package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func helper(splits []string, ndx int) (int, error) {
	if ndx < len(splits) {
		return processLine(splits, ndx+1)
	} else {
		return 0, nil
	}
}

func processLine(splits []string, ndx int) (int, error) {
	prev := -1
	var asc bool

	num1, err1 := strconv.Atoi(splits[1])
	num2, err2 := strconv.Atoi(splits[2])
	num0, err3 := strconv.Atoi(splits[0])
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("coult not parse int", err1, err2, err3)
	}
	if ndx == 0 {
		asc = num1 < num2
	} else if ndx == 1 {
		asc = num0 < num2
	} else {
		asc = num0 < num1
	}

	for n, elem := range splits {
		if n == ndx {
			continue
		}
		i, err := strconv.Atoi(elem)
		if err != nil {
			return 0, err
		}
		if prev >= 0 {
			diff := i - prev
			if !asc {
				diff = -diff
			}
			if diff <= 0 || diff > 3 {
				return helper(splits, ndx)
			}
		}
		prev = i
	}
	return 1, nil
}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, " ")
		chunk, err := processLine(splits, -1)
		if err != nil {
			fmt.Println("Error parsing int:", err)
			return
		}
		counter += chunk
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Output", strconv.Itoa(counter))
}
