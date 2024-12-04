package main

import (
	"bufio"

	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	path := flag.String("input", "", "Path to input data")
	part := flag.Int("part", 1, "Problem part to run")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var result int64
	switch *part {
	case -1:
		{
			line := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
			result, _ = sumValid(&line, true)
		}
	case 1:
		result, err = part1(f)
	case 2:
		result, err = part2(f)
	}
	if err != nil {
		log.Fatalf("Error while running part %d", *part)
	}
	log.Printf("Result for part %d: %d", *part, result)
}

func isByteAsciiDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func sumValid(p *string, default_ bool) (int64, bool) {
	line := *p
	var sum int64 = 0
	include := default_
	for i := 0; i < len(line)-7; i++ {
		j := i
		if line[j] == 'd' && line[j+1] == 'o' && line[j+2] == '(' && line[j+3] == ')' {
			include = true
			i += 2
			continue
		}
		if line[j] == 'd' && line[j+1] == 'o' && line[j+2] == 'n' && line[j+3] == '\'' && line[j+4] == 't' && line[j+5] == '(' && line[j+6] == ')' {
			include = false
			i += 5
			continue
		}

		if line[j] != 'm' || line[j+1] != 'u' || line[j+2] != 'l' || line[j+3] != '(' {
			continue
		}
		j += 4

		var left int64
		if isByteAsciiDigit(line[j]) && isByteAsciiDigit(line[j+1]) && isByteAsciiDigit(line[j+2]) {
			left = (int64(line[j])-48)*100 + (int64(line[j+1])-48)*10 + (int64(line[j+2]) - 48)
			j += 3
		} else if isByteAsciiDigit(line[j]) && isByteAsciiDigit(line[j+1]) {
			left = (int64(line[j])-48)*10 + (int64(line[j+1]) - 48)
			j += 2
		} else if isByteAsciiDigit(line[j]) {
			left = int64(line[j]) - 48
			j += 1
		} else {
			continue
		}

		if line[j] != ',' {
			continue
		}
		j += 1

		var right int64
		if isByteAsciiDigit(line[j]) && isByteAsciiDigit(line[j+1]) && isByteAsciiDigit(line[j+2]) {
			right = (int64(line[j])-48)*100 + (int64(line[j+1])-48)*10 + (int64(line[j+2]) - 48)
			j += 3
		} else if isByteAsciiDigit(line[j]) && isByteAsciiDigit(line[j+1]) {
			right = (int64(line[j])-48)*10 + (int64(line[j+1]) - 48)
			j += 2
		} else if isByteAsciiDigit(line[j]) {
			right = int64(line[j]) - 48
			j += 1
		} else {
			continue
		}

		if line[j] != ')' {
			continue
		}

		i = j // j is at the last index, the for loop will also increment by one
		if include {
			sum += left * right
		}
	}
	return sum, include
}

func part1(f io.Reader) (int64, error) {

	scr := bufio.NewScanner(f)

	var valid int64 = 0

	include := true
	for scr.Scan() {

		line := scr.Text()
		n, newInclude := sumValid(&line, include)
		include = newInclude
		valid += n
	}

	if err := scr.Err(); err != nil {
		return 0, fmt.Errorf("error scanning file: %v", err)
	}

	return valid, nil
}

func part2(f io.Reader) (int64, error) {
	return 0, nil
}
