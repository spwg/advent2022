package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type move struct {
	numCrates int
	from      int
	to        int
}

func parseCrates(s string) ([][]string, error) {
	lines := strings.Split(s, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("%q split by newlines has length 0", s)
	}
	n := ((len(lines[0]) + 1) / 4) - 1
	var stacks [][]string
	for i := 0; i <= n; i++ {
		stacks = append(stacks, []string{})
	}
	lines = lines[:len(lines)-1]
	for i := 0; i < len(lines); i++ {
		line := lines[len(lines)-1-i]
		p := 0
		for stack := 0; stack <= n; stack++ {
			q := p + 4
			if q >= len(line) {
				q = len(line)
			}
			container := line[p:q]
			if strings.TrimSpace(container) != "" {
				stacks[stack] = append(stacks[stack], container)
			}
			p += 4
		}
	}
	return stacks, nil
}

func parseMoves(s string) ([]move, error) {
	var moves []move
	for _, line := range strings.Split(s, "\n") {
		line = strings.TrimPrefix(line, "move ")
		line, n, err := trimNumber(line)
		if err != nil {
			return nil, fmt.Errorf("trim number: %v", err)
		}
		line = strings.TrimPrefix(line, " from ")
		line, from, err := trimNumber(line)
		if err != nil {
			return nil, fmt.Errorf("trim number: %v", err)
		}
		line = strings.TrimPrefix(line, " to ")
		_, to, err := trimNumber(line)
		if err != nil {
			return nil, fmt.Errorf("trim number: %v", err)
		}
		if from == to {
			return nil, fmt.Errorf(`"from" %v cannot equal "to" %v`, from, to)
		}
		moves = append(moves, move{n, from - 1, to - 1})
	}
	return moves, nil
}

func trimNumber(s string) (string, int, error) {
	i := strings.Index(s, " ")
	if i == -1 {
		i = len(s)
	}
	n, err := strconv.Atoi(s[:i])
	if err != nil {
		return "", 0, err
	}
	s = s[i:]
	return s, n, nil
}

func applyMoves(moves []move, crates [][]string) ([][]string, error) {
	for _, m := range moves {
		for i := 0; i < m.numCrates; i++ {
			j := len(crates[m.from]) - 1
			crates[m.to] = append(crates[m.to], crates[m.from][j])
			crates[m.from] = crates[m.from][:j]
		}
	}
	return crates, nil
}

func run(input string) error {
	before, after, found := strings.Cut(input, "\n\n")
	if !found {
		return fmt.Errorf("strings.Cut(%q, %q): not found", input, "\n\n")
	}
	moves, err := parseMoves(after)
	if err != nil {
		return err
	}
	crates, err := parseCrates(before)
	if err != nil {
		return fmt.Errorf("parse crates: %v", err)
	}
	crates, err = applyMoves(moves, crates)
	if err != nil {
		return err
	}
	for _, c := range crates {
		log.Printf("%#v", c)
	}
	return nil
}

func main() {
	err := run(`    [B]             [B] [S]        
    [M]             [P] [L] [B] [J]
    [D]     [R]     [V] [D] [Q] [D]
    [T] [R] [Z]     [H] [H] [G] [C]
    [P] [W] [J] [B] [J] [F] [J] [S]
[N] [S] [Z] [V] [M] [N] [Z] [F] [M]
[W] [Z] [H] [D] [H] [G] [Q] [S] [W]
[B] [L] [Q] [W] [S] [L] [J] [W] [Z]
 1   2   3   4   5   6   7   8   9 

move 3 from 5 to 2
move 5 from 3 to 1
move 4 from 4 to 9
move 6 from 1 to 4
move 6 from 8 to 7
move 5 from 2 to 7
move 1 from 5 to 4
move 11 from 9 to 7
move 1 from 1 to 9
move 6 from 4 to 6
move 12 from 6 to 7
move 1 from 9 to 2
move 2 from 4 to 6
move 1 from 8 to 9
move 1 from 9 to 4
move 1 from 6 to 1
move 2 from 7 to 5
move 2 from 6 to 7
move 2 from 1 to 6
move 2 from 4 to 7
move 1 from 5 to 4
move 1 from 5 to 6
move 1 from 6 to 1
move 1 from 1 to 3
move 1 from 4 to 1
move 1 from 1 to 4
move 1 from 4 to 5
move 1 from 3 to 9
move 1 from 5 to 1
move 4 from 2 to 1
move 20 from 7 to 8
move 24 from 7 to 3
move 3 from 6 to 4
move 1 from 1 to 9
move 1 from 9 to 3
move 2 from 1 to 2
move 2 from 4 to 1
move 2 from 2 to 1
move 14 from 3 to 6
move 6 from 1 to 6
move 10 from 3 to 2
move 1 from 2 to 3
move 6 from 6 to 5
move 2 from 3 to 4
move 13 from 8 to 4
move 1 from 9 to 7
move 1 from 6 to 3
move 10 from 4 to 2
move 1 from 3 to 6
move 2 from 8 to 7
move 1 from 7 to 2
move 11 from 6 to 8
move 2 from 6 to 1
move 2 from 1 to 3
move 1 from 8 to 6
move 1 from 3 to 9
move 3 from 8 to 2
move 1 from 3 to 6
move 2 from 6 to 4
move 1 from 6 to 5
move 11 from 2 to 9
move 2 from 4 to 6
move 1 from 6 to 1
move 1 from 1 to 5
move 11 from 2 to 7
move 12 from 7 to 5
move 1 from 6 to 2
move 10 from 8 to 7
move 6 from 5 to 3
move 4 from 5 to 4
move 11 from 9 to 7
move 7 from 4 to 9
move 4 from 9 to 6
move 12 from 7 to 3
move 1 from 8 to 9
move 1 from 5 to 1
move 1 from 1 to 2
move 1 from 6 to 9
move 3 from 4 to 1
move 1 from 9 to 7
move 8 from 7 to 2
move 3 from 6 to 1
move 8 from 2 to 3
move 1 from 7 to 4
move 2 from 7 to 2
move 1 from 5 to 2
move 8 from 5 to 1
move 3 from 9 to 6
move 1 from 6 to 2
move 1 from 4 to 5
move 1 from 5 to 4
move 2 from 9 to 3
move 1 from 8 to 6
move 1 from 4 to 5
move 1 from 5 to 1
move 1 from 6 to 8
move 1 from 8 to 1
move 7 from 1 to 5
move 11 from 3 to 7
move 1 from 1 to 9
move 4 from 2 to 1
move 5 from 1 to 3
move 1 from 5 to 9
move 1 from 6 to 3
move 6 from 2 to 1
move 5 from 7 to 3
move 1 from 6 to 8
move 1 from 8 to 4
move 6 from 7 to 9
move 4 from 9 to 8
move 2 from 8 to 9
move 2 from 5 to 8
move 13 from 3 to 7
move 1 from 3 to 8
move 2 from 1 to 9
move 3 from 1 to 5
move 1 from 4 to 1
move 6 from 5 to 9
move 8 from 9 to 8
move 2 from 7 to 3
move 1 from 9 to 7
move 1 from 5 to 2
move 5 from 9 to 8
move 1 from 8 to 7
move 1 from 2 to 9
move 7 from 1 to 2
move 4 from 7 to 5
move 6 from 2 to 3
move 1 from 2 to 1
move 10 from 8 to 9
move 3 from 8 to 9
move 4 from 5 to 1
move 2 from 8 to 6
move 9 from 9 to 8
move 1 from 9 to 6
move 8 from 8 to 4
move 12 from 3 to 5
move 1 from 4 to 2
move 3 from 8 to 1
move 3 from 9 to 7
move 1 from 3 to 2
move 1 from 6 to 9
move 8 from 3 to 8
move 6 from 4 to 5
move 1 from 7 to 6
move 1 from 8 to 1
move 6 from 8 to 7
move 1 from 3 to 6
move 7 from 1 to 5
move 1 from 4 to 9
move 4 from 6 to 5
move 13 from 7 to 5
move 1 from 8 to 2
move 2 from 9 to 3
move 4 from 7 to 2
move 1 from 3 to 8
move 1 from 3 to 4
move 4 from 1 to 2
move 1 from 5 to 7
move 23 from 5 to 6
move 1 from 8 to 6
move 1 from 9 to 4
move 5 from 2 to 6
move 1 from 4 to 9
move 1 from 9 to 3
move 1 from 7 to 8
move 1 from 4 to 3
move 1 from 3 to 7
move 1 from 7 to 5
move 1 from 8 to 7
move 12 from 6 to 1
move 1 from 2 to 5
move 1 from 3 to 1
move 20 from 5 to 2
move 14 from 2 to 4
move 11 from 2 to 6
move 1 from 7 to 8
move 13 from 1 to 8
move 9 from 8 to 4
move 3 from 8 to 6
move 10 from 6 to 8
move 6 from 6 to 4
move 4 from 8 to 5
move 26 from 4 to 2
move 2 from 5 to 2
move 5 from 8 to 1
move 1 from 8 to 3
move 2 from 1 to 3
move 2 from 3 to 7
move 27 from 2 to 7
move 2 from 8 to 1
move 1 from 3 to 7
move 6 from 6 to 2
move 4 from 6 to 1
move 4 from 6 to 4
move 2 from 5 to 4
move 4 from 2 to 1
move 3 from 1 to 8
move 1 from 2 to 8
move 8 from 4 to 3
move 1 from 2 to 8
move 5 from 8 to 6
move 1 from 4 to 2
move 1 from 2 to 1
move 6 from 3 to 1
move 13 from 7 to 1
move 1 from 2 to 8
move 1 from 8 to 2
move 1 from 6 to 2
move 1 from 2 to 8
move 1 from 8 to 2
move 14 from 7 to 1
move 5 from 6 to 3
move 2 from 3 to 1
move 3 from 3 to 2
move 3 from 7 to 4
move 1 from 4 to 9
move 1 from 9 to 7
move 2 from 3 to 6
move 5 from 2 to 7
move 1 from 7 to 6
move 5 from 7 to 6
move 2 from 6 to 7
move 1 from 6 to 8
move 1 from 4 to 7
move 4 from 6 to 9
move 35 from 1 to 8
move 3 from 7 to 2
move 1 from 2 to 5
move 24 from 8 to 3
move 1 from 5 to 8
move 13 from 3 to 6
move 2 from 2 to 6
move 6 from 6 to 4
move 11 from 1 to 6
move 12 from 6 to 1
move 1 from 8 to 1
move 2 from 1 to 3
move 5 from 4 to 1
move 1 from 6 to 4
move 1 from 8 to 3
move 13 from 3 to 9
move 3 from 8 to 2
move 3 from 2 to 7
move 1 from 3 to 6
move 3 from 7 to 8
move 14 from 1 to 3
move 1 from 1 to 9
move 6 from 3 to 8
move 17 from 8 to 6
move 1 from 3 to 7
move 1 from 7 to 8
move 26 from 6 to 7
move 1 from 1 to 9
move 3 from 4 to 1
move 2 from 3 to 8
move 1 from 8 to 4
move 14 from 9 to 7
move 12 from 7 to 3
move 2 from 1 to 4
move 2 from 7 to 8
move 2 from 8 to 3
move 4 from 9 to 8
move 1 from 4 to 7
move 1 from 1 to 3
move 2 from 4 to 2
move 24 from 7 to 6
move 1 from 8 to 1
move 1 from 7 to 2
move 1 from 7 to 9
move 3 from 2 to 9
move 1 from 1 to 6
move 5 from 8 to 2
move 5 from 3 to 4
move 1 from 2 to 5
move 3 from 9 to 8
move 2 from 4 to 9
move 16 from 6 to 3
move 14 from 3 to 8
move 1 from 7 to 9
move 8 from 6 to 9
move 4 from 8 to 5
move 8 from 8 to 3
move 1 from 5 to 8
move 1 from 2 to 4
move 4 from 8 to 7
move 1 from 5 to 6
move 12 from 9 to 5
move 15 from 5 to 8
move 1 from 6 to 1
move 2 from 2 to 6
move 3 from 4 to 2
move 4 from 2 to 7
move 8 from 7 to 3
move 1 from 1 to 4
move 3 from 6 to 9
move 16 from 8 to 3
move 3 from 9 to 4
move 1 from 8 to 9
move 2 from 9 to 4
move 24 from 3 to 8
move 19 from 8 to 7
move 2 from 8 to 7
move 7 from 4 to 5
move 13 from 7 to 5
move 4 from 7 to 8
move 7 from 8 to 1
move 3 from 5 to 3
move 3 from 7 to 2
move 1 from 1 to 4
move 1 from 7 to 2
move 3 from 2 to 4
move 8 from 3 to 1
move 11 from 1 to 3
move 12 from 3 to 4
move 1 from 2 to 5
move 18 from 3 to 8
move 3 from 1 to 9
move 1 from 3 to 5
move 15 from 5 to 4
move 4 from 5 to 1
move 23 from 4 to 6
move 3 from 1 to 6
move 13 from 8 to 3
move 25 from 6 to 2
move 1 from 9 to 5
move 5 from 3 to 8
move 17 from 2 to 8
move 4 from 4 to 1
move 1 from 9 to 7
move 5 from 2 to 6
move 2 from 2 to 4
move 1 from 9 to 4
move 6 from 3 to 9
move 16 from 8 to 3
move 2 from 1 to 8
move 1 from 7 to 4
move 5 from 4 to 7
move 1 from 5 to 3
move 2 from 7 to 1
move 9 from 8 to 4
move 3 from 7 to 2
move 2 from 8 to 3
move 10 from 4 to 1
move 1 from 2 to 3
move 5 from 3 to 7
move 2 from 8 to 9
move 2 from 9 to 8
move 1 from 2 to 1
move 3 from 9 to 6
move 2 from 2 to 8
move 4 from 7 to 3
move 4 from 8 to 6
move 1 from 7 to 1
move 1 from 4 to 8
move 4 from 3 to 4
move 4 from 4 to 2
move 6 from 1 to 2
move 1 from 4 to 3
move 5 from 3 to 8
move 6 from 3 to 8
move 2 from 2 to 8
move 3 from 2 to 9
move 8 from 1 to 6
move 3 from 2 to 7
move 2 from 7 to 2
move 13 from 6 to 5
move 7 from 5 to 9
move 3 from 2 to 7
move 1 from 2 to 9
move 2 from 5 to 2
move 3 from 8 to 5
move 5 from 3 to 4
move 2 from 2 to 1
move 9 from 8 to 7
move 1 from 1 to 8
move 6 from 5 to 2
move 4 from 2 to 8
move 4 from 7 to 1
move 1 from 2 to 6
move 5 from 1 to 6
move 1 from 8 to 2
move 1 from 2 to 9
move 13 from 6 to 5
move 2 from 7 to 2
move 1 from 8 to 7
move 4 from 4 to 7
move 1 from 4 to 1
move 4 from 8 to 4
move 6 from 5 to 9
move 2 from 1 to 4
move 1 from 8 to 6
move 11 from 9 to 5
move 1 from 7 to 8
move 1 from 8 to 1
move 1 from 1 to 3
move 6 from 4 to 8
move 1 from 8 to 4
move 1 from 1 to 6
move 6 from 9 to 7
move 1 from 4 to 5
move 3 from 2 to 1
move 1 from 8 to 2
move 1 from 3 to 2
move 20 from 5 to 6
move 3 from 1 to 6
move 2 from 2 to 9
move 3 from 8 to 3
move 5 from 3 to 8
move 1 from 1 to 6
move 2 from 8 to 9
move 7 from 9 to 5
move 3 from 5 to 4
move 3 from 8 to 3
move 9 from 7 to 9
move 1 from 8 to 5
move 7 from 7 to 9
move 2 from 5 to 2
move 9 from 9 to 2
move 1 from 7 to 3
move 2 from 9 to 1
move 2 from 5 to 9
move 2 from 1 to 4
move 2 from 3 to 7
move 18 from 6 to 7
move 7 from 9 to 1
move 7 from 6 to 8
move 4 from 4 to 9
move 4 from 8 to 3
move 2 from 8 to 2
move 1 from 8 to 5
move 1 from 4 to 7
move 1 from 5 to 1
move 2 from 9 to 3
move 12 from 2 to 5
move 6 from 5 to 6
move 5 from 7 to 2
move 3 from 6 to 4
move 1 from 4 to 7
move 1 from 4 to 1
move 2 from 5 to 8
move 1 from 8 to 2
move 2 from 9 to 7
move 8 from 1 to 8
move 11 from 7 to 1
move 5 from 8 to 2
move 7 from 7 to 5
move 1 from 9 to 4
move 1 from 7 to 5
move 7 from 5 to 7
move 2 from 6 to 1
move 1 from 8 to 2
move 12 from 1 to 7
move 2 from 1 to 2
move 3 from 8 to 5
move 3 from 5 to 2
move 8 from 7 to 3
move 1 from 3 to 1
move 3 from 6 to 4
move 4 from 5 to 6
move 14 from 2 to 9
move 3 from 6 to 9
move 3 from 4 to 2
move 1 from 1 to 7
move 1 from 7 to 1
move 3 from 3 to 5
move 8 from 7 to 4
move 1 from 5 to 9
move 3 from 2 to 4
move 1 from 3 to 4
move 4 from 2 to 6
move 2 from 6 to 7
move 3 from 5 to 4
move 16 from 4 to 1
move 7 from 9 to 8
move 1 from 5 to 1
move 3 from 7 to 9
move 3 from 9 to 4
move 7 from 1 to 7
move 6 from 7 to 1
move 5 from 3 to 1
move 11 from 9 to 2
move 3 from 4 to 6
move 9 from 2 to 8
move 6 from 3 to 5
move 2 from 8 to 6
move 5 from 5 to 3
move 2 from 7 to 1
move 3 from 3 to 9
move 1 from 2 to 4
move 1 from 5 to 1
move 13 from 1 to 2
move 5 from 8 to 6
move 2 from 3 to 9
move 2 from 4 to 7
move 5 from 6 to 9
move 7 from 9 to 1
move 3 from 7 to 2
move 6 from 8 to 6
move 5 from 6 to 2
move 2 from 8 to 3
move 2 from 9 to 4
move 6 from 2 to 5
move 1 from 3 to 7`)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}
