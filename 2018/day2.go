package main

import (
	"strconv"
)

// Returns two values:
//   - number of letters that appear exactly twice in ‘s’
//   - number of letters that appear exactly thrice in ‘s’
func dupecounts(s string) (int, int) {
	two := 0
	three := 0
	counts := make(map[rune]int)

	for _, c := range s {
		counts[c] += 1
	}

	for _, v := range counts {
		if v == 2 {
			two++
		} else if v == 3 {
			three++
		}
	}

	return two, three
}

func day2_A(lines []string) int {
	mul1 := 0
	mul2 := 0
	for _, s := range lines {
		two, three := dupecounts(s)
		if two > 0 {
			mul1++
		}
		if three > 0 {
			mul2++
		}
	}
	return mul1 * mul2
}

func day2_B(lines []string) string {
	// Algorithm from https://www.reddit.com/r/adventofcode/comments/a2damm/_/eaxco3u.
	// Also interesting: https://www.reddit.com/r/adventofcode/comments/a2damm/_/eaynnaw.
	size := len(lines[0])

	for i := 0; i < size; i++ {
		seen := make(map[string]bool, len(lines))
		for _, s := range lines {
			ss := s[:i] + s[i+1:]
			if seen[ss] {
				return ss
			}
			seen[ss] = true
		}
	}
	return ""
}

func day2() (string, string) {
	lines := readLines("input/02")
	a := day2_A(lines)
	b := day2_B(lines)
	return strconv.Itoa(a), b
}
