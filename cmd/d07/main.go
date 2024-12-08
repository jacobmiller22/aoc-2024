package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := flag.String("input", "", "path to input file")

	flag.Parse()

	if *path == "" {
		log.Fatalf("provided an empty input: '%s'", *path)
	}

	f, err := os.Open(*path)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	r1, r2, err := solution(f)
	if err != nil {
		log.Fatalf("error executing solution: %v", err)
	}

	log.Printf("Results: Part 1=%d; Part 2=%d\n", r1, r2)
}

type equation struct {
	ans      int64
	operands []int64
}

const SUM_OP = 1
const MULT_OP = 2
const CONCAT_OP = 3

type operator struct {
	kind int
}

func (o *operator) Exec(a, b int64) (int64, error) {
	switch o.kind {
	case SUM_OP:
		return a + b, nil
	case MULT_OP:
		return a * b, nil
	case CONCAT_OP:
		l := strconv.FormatInt(a, 10)
		r := strconv.FormatInt(b, 10)
		c := l + r
		res, err := strconv.ParseInt(c, 10, 64)
		return res, err
	}
	return 0, fmt.Errorf("invalid operator.. shouldn't happen")
}

func (o *operator) String() string {
	switch o.kind {
	case SUM_OP:
		return "+"
	case MULT_OP:
		return "*"
	case CONCAT_OP:
		return "||"
	}
	panic("shouldn't happen")
}

func GeneratePermutations(operators []operator, n int) [][]operator {
	var results [][]operator
	var helper func([]operator)
	helper = func(current []operator) {
		if len(current) == n {
			// Make a copy of current to avoid issues with slice reuse
			temp := make([]operator, len(current))
			copy(temp, current)
			results = append(results, temp)
			return
		}
		for _, op := range operators {
			helper(append(current, op))
		}
	}
	helper([]operator{})
	return results
}

func (e *equation) ValidOperations(ops []operator) ([][]operator, error) {
	valid := make([][]operator, 0, 0)
	perms := GeneratePermutations(ops, len(e.operands)-1)
	for _, perm := range perms {
		var ans int64 = e.operands[0]
		for i, op := range perm {
			res, err := op.Exec(ans, e.operands[i+1])
			if err != nil {
				return nil, fmt.Errorf("error executing operation: %v", err)
			}
			ans = res
		}
		if ans == e.ans {
			valid = append(valid, perm)
		}
	}
	return valid, nil
}

func NewEquationFromString(e string) (*equation, error) {
	split1 := strings.Split(e, ":")

	if len(split1) != 2 {
		return nil, fmt.Errorf("invalid equation string")
	}

	ans, err := strconv.ParseInt(split1[0], 10, 64)

	if err != nil {
		return nil, fmt.Errorf("error parsing answer to int64: %v", err)
	}

	operands := make([]int64, 0, 0)

	split2 := strings.Split(split1[1], " ")

	for _, opStr := range split2 {
		if opStr == "" {
			continue
		}
		op, err := strconv.ParseInt(opStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing operand to int64: %v", err)
		}

		operands = append(operands, op)
	}

	return &equation{ans, operands}, nil
}

func solution(f io.Reader) (int64, int64, error) {
	scr := bufio.NewScanner(f)

	var sum1 int64 = 0
	var sum2 int64 = 0

	var ops1 = []operator{{kind: SUM_OP}, {kind: MULT_OP}}
	var ops2 = []operator{{kind: SUM_OP}, {kind: MULT_OP}, {kind: CONCAT_OP}}

	for scr.Scan() {
		line := scr.Text()

		// ans :=
		eq, err := NewEquationFromString(line)
		if err != nil {
			return 0, 0, fmt.Errorf("error parsing equation: %v", err)
		}
		validOps, err := eq.ValidOperations(ops1)
		if err != nil {
			return 0, 0, err
		}
		if len(validOps) > 0 {
			sum1 += eq.ans
		}

		validOps, err = eq.ValidOperations(ops2)
		if err != nil {
			return 0, 0, err
		}
		if len(validOps) > 0 {
			sum2 += eq.ans
		}
	}

	if err := scr.Err(); err != nil {
		return 0, 0, fmt.Errorf("error scanning file: %v", err)
	}

	return sum1, sum2, nil
}
