package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jacobmiller22/aoc-2024/math"
)

const WS rune = 0x20
const HT rune = 0x23
const GUARD rune = 0x5E // Guard ^

const UP int = 0
const RIGHT int = 1
const DOWN int = 2
const LEFT int = 3

var errLoop = fmt.Errorf("infinite loop")

func main() {

	path := flag.String("input", "", "Path to input file")

	flag.Parse()

	f, err := os.Open(*path)

	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}

	p1, p2, err := solution(f)

	if err != nil {
		log.Fatalf("error executing solution: %v", err)
	}

	log.Printf("Part 1: %d :==: Part 2: %d\n", p1, p2)
}

type Marking struct {
	kind rune
	y    int
}

func scanMap(f io.Reader) (*board, *coordinate, error) {
	scr := bufio.NewScanner(f)

	b := newBoard()

	y := 0
	maxX := 0
	var guardCoord *coordinate
	for scr.Scan() {
		line := scr.Text()

		for x, c := range line {
			maxX = math.Max(maxX, x)
			if c == GUARD {
				guardCoord = &coordinate{x, y}
			}
			if c == HT {
				b.Mark(x, y)
			}
		}
		y++
	}

	if err := scr.Err(); err != nil {
		return nil, nil, err
	}

	if guardCoord == nil {
		return nil, nil, fmt.Errorf("can't find guard location")
	}

	b.SetWidth(maxX + 1)
	b.SetHeight(y)

	return b, guardCoord, nil
}

type coordinate struct {
	x int
	y int
}

type board struct {
	m map[coordinate]bool
	h int
	w int
}

func newBoard() *board {
	return &board{
		m: make(map[coordinate]bool),
	}
}

func (b *board) Clear() {
	for k := range b.m {
		b.m[k] = false
	}
}

func (b *board) Mark(x, y int) {
	coord := coordinate{x, y}
	b.m[coord] = true
}

func (b *board) Unmark(x, y int) {
	coord := coordinate{x, y}
	b.m[coord] = false
}

func (b *board) Has(x, y int) bool {
	coord := coordinate{x, y}
	v, _ := b.m[coord]
	return v
}

func (b *board) Height() int {
	return b.h
}

func (b *board) SetHeight(h int) {
	b.h = h
}

func (b *board) Width() int {
	return b.w
}
func (b *board) SetWidth(w int) {
	b.w = w
}

func printVisualization(ob, tb *board) {
	fmt.Println("\n---Board Viewer---")
	for col := 0; col < ob.Width(); col++ {
		line := ""
		for row := 0; row < ob.Height(); row++ {
			if ob.Has(row, col) {
				line += "#"
			} else if tb.Has(row, col) {
				line += "X"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}

}

func solution(f io.Reader) (int, int, error) {

	/// ob is obstacle board, tb is traveled board
	ob, gc, err := scanMap(f)

	if err != nil {
		return 0, 0, err
	}

	startX, startY := gc.x, gc.y
	tb := newBoard()
	tb.SetWidth(ob.Width())
	tb.SetHeight(ob.Height())

	sum1, sum2 := 0, 0

	history, err := simulate(ob, tb, *gc)
	if err != nil {
		log.Fatalf("error during first simulation: %v", err)
	}

	sum1 = len(*history)

	// For part 2 we will inspect the travel original travel history,
	// we will place a new obstruction at each newly discovered spot and see if this will give us an
	// infinite loop error

	for _, discovery := range *history {

		if discovery.x == startX && discovery.y == startY {
			continue
		}
		ob.Mark(discovery.x, discovery.y)
		tb := newBoard()

		_, err := simulate(ob, tb, *gc)
		if errors.Is(err, errLoop) {
			sum2++
		}
		ob.Unmark(discovery.x, discovery.y)

	}

	return sum1, sum2, err
}

func simulate(ob, tb *board, gc coordinate) (*[]coordinate, error) {

	// Guard starts UP
	gd := UP

	discoveryCount := 0
	history := make([]coordinate, 0, 0)
	spotsMoved := 0

	for {
		nx, ny := gc.x, gc.y
		if gd == DOWN {
			if gc.y >= ob.Height() {
				// on the edge of the map, we out
				break
			}
			// get next position
			ny++

		} else if gd == UP {
			// get next up
			if gc.y <= 0 {
				// on the edge of the map, we out
				break
			}
			ny--
		} else if gd == LEFT {
			if gc.x <= 0 {
				// on the edge of the map, we out
				break
			}
			nx--
		} else {
			if gc.x >= ob.Width() {
				// on the edge of the map, we out
				break
			}
			nx++
		}

		// check next pos for obstacles
		if ob.Has(nx, ny) {
			// turn to the right
			gd = (gd + 1) % 4
			if gd == UP {
				// Take a snapshot of our discovery count, if we reach
				// UP again with no new discoveries then we are in a loop
				if len(history) == discoveryCount {
					return nil, errLoop
				}
				discoveryCount = len(history)
			}
			continue
		}

		if !tb.Has(gc.x, gc.y) {
			history = append(history, coordinate{gc.x, gc.y})
		}
		tb.Mark(gc.x, gc.y)
		gc.x, gc.y = nx, ny
		spotsMoved++
	}
	return &history, nil
}
