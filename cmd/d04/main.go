package main

import (
	"bytes"
	"flag"
	"log"
	"os"
)

func main() {

	path := flag.String("input", "", "Path to input")
	part := flag.Int("part", 1, "The problem part to run")

	flag.Parse()

	data, err := os.ReadFile(*path)

	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	var result int
	switch *part {
	case 1:
		result, err = part1(&data)
	case 2:
		result, err = part2(&data)
	}

	if err != nil {
		log.Fatalf("error from part func: %v", err)
	}

	log.Printf("Result from part %d: %d\n", *part, result)
}

// Find all occurences of xmas from the given pointer in the provided matrix
// Assumes that the given coordiantes are already on an `X`
func findXmas(p *[][]byte, r, c int) int {
	mat := *p
	n := 0
	numRows := len(mat) - 1
	numCols := len(mat[0])

	// Inline Horizontal
	if c < numCols-3 && mat[r][c+1] == 'M' && mat[r][c+2] == 'A' && mat[r][c+3] == 'S' {
		n += 1
	}
	if c >= 3 && mat[r][c-1] == 'M' && mat[r][c-2] == 'A' && mat[r][c-3] == 'S' {
		n += 1
	}

	// Inline Vertical
	if r < numRows-3 && mat[r+1][c] == 'M' && mat[r+2][c] == 'A' && mat[r+3][c] == 'S' {
		n += 1
	}

	if r >= 3 && mat[r-1][c] == 'M' && mat[r-2][c] == 'A' && mat[r-3][c] == 'S' {
		n += 1
	}

	// Diagonal top left
	if r >= 3 && c >= 3 && mat[r-1][c-1] == 'M' && mat[r-2][c-2] == 'A' && mat[r-3][c-3] == 'S' {
		n += 1
	}

	// Diagonal top right
	if r >= 3 && c < numCols-3 && mat[r-1][c+1] == 'M' && mat[r-2][c+2] == 'A' && mat[r-3][c+3] == 'S' {
		n += 1
	}

	// Diagonal bottom left
	if r < numRows-3 && c >= 3 && mat[r+1][c-1] == 'M' && mat[r+2][c-2] == 'A' && mat[r+3][c-3] == 'S' {
		n += 1
	}

	// Diagonal bottom right
	if r < numRows-3 && c < numCols-3 && mat[r+1][c+1] == 'M' && mat[r+2][c+2] == 'A' && mat[r+3][c+3] == 'S' {
		n += 1
	}
	return n

}

func findMas(p *[][]byte, r, c int) int {
	mat := *p
	numRows := len(mat) - 1
	numCols := len(mat[0])
	if r <= 0 || r >= numRows-1 || c <= 0 || c >= numCols-1 {
		return 0
	}

	n := 0
	// diagonal down right MAS
	// M M
	//  A
	// S S
	if mat[r-1][c-1] == 'M' && mat[r-1][c+1] == 'M' && mat[r+1][c-1] == 'S' && mat[r+1][c+1] == 'S' {
		n += 1
	}

	// diagonal down right SAM
	// M S
	//  A
	// M S
	if mat[r-1][c-1] == 'M' && mat[r-1][c+1] == 'S' && mat[r+1][c-1] == 'M' && mat[r+1][c+1] == 'S' {
		n += 1
	}

	// diagonal down left MAS
	// S S
	//  A
	// M M
	if mat[r-1][c-1] == 'S' && mat[r-1][c+1] == 'S' && mat[r+1][c-1] == 'M' && mat[r+1][c+1] == 'M' {
		n += 1
	}

	// diagonal down left SAM
	// S M
	//  A
	// S M
	if mat[r-1][c-1] == 'S' && mat[r-1][c+1] == 'M' && mat[r+1][c-1] == 'S' && mat[r+1][c+1] == 'M' {
		n += 1
	}

	return n

}

func part1(data *[]byte) (int, error) {
	matrix := bytes.Split(*data, []byte{'\n'})
	sum := 0
	for r := 0; r < len(matrix)-1; r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == 'X' {
				sum += findXmas(&matrix, r, c)
			}
		}
	}

	return sum, nil
}

func part2(data *[]byte) (int, error) {
	matrix := bytes.Split(*data, []byte{'\n'})
	sum := 0
	for r := 0; r < len(matrix)-1; r++ {
		for c := 0; c < len(matrix[0]); c++ {
			if matrix[r][c] == 'A' {
				sum += findMas(&matrix, r, c)
			}
		}
	}

	return sum, nil
}
