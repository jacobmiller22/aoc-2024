package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/jacobmiller22/aoc-2024/math"
)

func main() {
	path := flag.String("input", "", "Path to input file")

	flag.Parse()

	if *path == "" {
		log.Fatalf("Provided path must not be empty: '%v'\n", *path)
	}

	f, err := os.Open(*path)
	if err != nil {
		log.Fatalf("error opening input file: %v", err)
	}
	p1, p2, err := solution(f)

	if err != nil {
		log.Fatalf("error running solution: %v", err)
	}

	log.Printf("Results: p1=%d; p2=%d\n", p1, p2)
}

type Block struct {
	id    int
	used  int
	free  int
	moved bool
}

func NewBlock(id, used, free int, moved bool) *Block {
	return &Block{id, used, free, moved}
}

func part1(blocks []*Block) int {

	j := len(blocks) - 1 // the back pointer; index of of the last block
	for i := 0; i < len(blocks) && i < j; i++ {
		if blocks[i].free <= 0 {
			continue
		}

		transfer := math.Min(blocks[i].free, blocks[j].used)
		blocks = slices.Insert(blocks, i+1, NewBlock(blocks[j].id, transfer, blocks[i].free-transfer, false))
		j++
		blocks[i].free = 0
		blocks[j].used -= transfer
		if blocks[j].used == 0 {
			// this block has been copied completely, remove it
			blocks = slices.Delete(blocks, j, j+1)
			j--
		}

	}

	sum1 := 0
	k := 0
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < blocks[i].used; j++ {
			sum1 += k * blocks[i].id
			k += 1
		}
	}

	return sum1

}

func visualizeDiskMap(blocks []*Block) {
	var w strings.Builder
	for _, b := range blocks {
		for i := 0; i < b.used; i++ {
			w.WriteString(strconv.Itoa(b.id))
		}
		for i := 0; i < b.free; i++ {
			w.WriteString(".")
		}
	}
	w.WriteString("\n")
	fmt.Print(w.String())
}

func part2chatgpt(blocks []*Block) int {
	// Step 1: Sort blocks by ID in descending order (file movement order)
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].id > blocks[j].id
	})

	for i := 0; i < len(blocks); i++ {
		if blocks[i].moved || blocks[i].used == 0 {
			continue
		}

		// Attempt to move the file to the leftmost span of free space
		for j := 0; j < i; j++ {
			if blocks[j].free >= blocks[i].used {
				// Move the file to blocks[j]
				blocks[j].free -= blocks[i].used
				blocks = slices.Insert(blocks, j+1, NewBlock(blocks[i].id, blocks[i].used, 0, true))

				// Update the original block (blocks[i])
				blocks[i].free += blocks[i].used
				blocks[i].used = 0
				blocks[i].moved = true
				break
			}
		}
	}

	// Step 2: Calculate the checksum
	sum2 := 0
	k := 0
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < blocks[i].used; j++ {
			sum2 += k * blocks[i].id
			k++
		}
		k += blocks[i].free
	}

	return sum2
}

func part2(blocks []*Block) int {

	for i := len(blocks) - 1; i >= 0; i-- {
		if blocks[i].moved {
			continue
		}

		for j := 0; j < i; j++ {
			if blocks[j].free >= blocks[i].used {
				blocks = slices.Insert(blocks, j+1, NewBlock(blocks[i].id, blocks[i].used, blocks[j].free-blocks[i].used, true))
				blocks[i].free += blocks[i+1].used + blocks[i+1].free
				blocks = slices.Delete(blocks, i+1, i+2)
				blocks[j].free = 0
				i++
				break
			}
		}
	}

	sum2 := 0
	k := 0
	for i := 0; i < len(blocks); i++ {
		for j := 0; j < blocks[i].used; j++ {
			sum2 += k * blocks[i].id
			k += 1
		}
		k += blocks[i].free
	}

	return sum2
}

func solution(f io.Reader) (int, int, error) {

	rdr := bufio.NewReader(f)
	line, err := rdr.ReadString('\n')
	if err != nil {
		return 0, 0, fmt.Errorf("error reading line: %v", err)
	}

	line = strings.TrimSuffix(line, "\n")

	blocks1 := make([]*Block, 0, len(line)/2)
	blocks2 := make([]*Block, 0, len(line)/2)

	fmt.Printf("len(line): %d\n", len(line))

	k := 0
	for i := 0; i < len(line); i += 2 {
		used := int(line[i] - '0')
		free := 0
		if i+1 < len(line) {
			free = int(line[i+1] - '0')
		}
		k += used + free

		blocks1 = append(blocks1, NewBlock(i/2, used, free, false))
		blocks2 = append(blocks2, NewBlock(i/2, used, free, false))
	}
	fmt.Printf("expected k: %d\n", k)

	return part1(blocks1), part2(blocks2), nil
}
