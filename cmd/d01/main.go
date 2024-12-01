package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/jacobmiller22/aoc-2024/math"
)

func main() {

	inputPath := flag.String("input", "", "Path to input")
	part := flag.Int("part", 1, "The question part to run")

	flag.Parse()

	if *inputPath == "" {
		log.Panic("Input path cannot be empty")
	}

	f, err := os.Open(*inputPath)

	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}

	scanner := bufio.NewScanner(f)

	var result int
	switch *part {
	case 1:
		result, err = part1(scanner)
	case 2:
		result, err = part2(scanner)

	}

	if err != nil {
		log.Fatalf("Error occurred while running part %d: %v", part, err)
	}

	log.Printf("Part %d Answer is: %d", part, result)
}

func part1(scanner *bufio.Scanner) (int, error) {
	var leftNums sort.IntSlice = make([]int, 0, 0)
	var rightNums sort.IntSlice = make([]int, 0, 0)

	for scanner.Scan() {
		line := scanner.Text()
		left, err := strconv.Atoi(line[0:5])
		if err != nil {
			return 0, fmt.Errorf("Left side is not a number: %v", err)
		}
		leftNums = append(leftNums, left)

		right, err := strconv.Atoi(line[8:13])
		if err != nil {
			return 0, fmt.Errorf("Right side is not a number: %v", err)
		}
		rightNums = append(rightNums, right)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error while scanning input: %v", err)
	}

	// sort
	leftNums.Sort()
	rightNums.Sort()

	if leftNums.Len() != rightNums.Len() {
		return 0, fmt.Errorf("The length of left nums does not equal the length of right nums!")
	}

	sum := 0
	for i, left := range leftNums {
		sum += math.Abs(left - rightNums[i])
	}

	return sum, nil
}

func part2(scanner *bufio.Scanner) (int, error) {
	leftNums := make([]int, 0, 0)
	rightNums := make([]int, 0, 0)

	for scanner.Scan() {
		line := scanner.Text()
		left, err := strconv.Atoi(line[0:5])
		if err != nil {
			return 0, fmt.Errorf("Left side is not a number: %v", err)
		}
		leftNums = append(leftNums, left)

		right, err := strconv.Atoi(line[8:13])
		if err != nil {
			return 0, fmt.Errorf("Right side is not a number: %v", err)
		}
		rightNums = append(rightNums, right)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error while scanning input: %v", err)
	}

	if len(leftNums) != len(rightNums) {
		return 0, fmt.Errorf("The length of left nums does not equal the length of right nums!")
	}

	rightNumCounts := make(map[int]int, len(leftNums))

	for _, right := range rightNums {
		rightNumCounts[right] += 1 // default value for int is 0, so we don't need to worry about the missing key case
	}

	sum := 0
	for _, left := range leftNums {
		freq, _ := rightNumCounts[left]
		sum += left * freq
	}
	return sum, nil
}
