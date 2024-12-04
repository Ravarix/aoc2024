package day4

import (
	"bufio"
	"log"
	"os"
)

var directions = [][]int{{0, 1}, {1, 0}, {1, 1}, {0, -1}, {-1, 0}, {-1, -1}, {1, -1}, {-1, 1}}
var word = []rune("XMAS")

type matrix [][]rune

func Run() {
	m := matrix(parseFile())
	hits := 0
	for i, row := range m {
		for j, char := range row {
			if char == 'X' {
				hits += m.testLocation(i, j)
			}
		}
	}
	log.Println("crossword hits: ", hits)

	//Part 2
	hits = 0
	for i, row := range m {
		for j, char := range row {
			if char == 'A' && m.testStep2(i, j) {
				hits++
			}
		}
	}
	log.Println("X-MAS hits: ", hits)
}

func (m matrix) testStep2(x, y int) bool {
	if !m.inBounds(x+1, y+1) || !m.inBounds(x-1, y-1) {
		return false
	}
	// Add opposite corners to sets
	line1 := map[rune]struct{}{
		m[x+1][y+1]: {},
		m[x-1][y-1]: {},
	}
	line2 := map[rune]struct{}{
		m[x+1][y-1]: {},
		m[x-1][y+1]: {},
	}
	// Check each set contains 'M' & 'S'
	if _, ok := line1['M']; !ok {
		return false
	}
	if _, ok := line1['S']; !ok {
		return false
	}
	if _, ok := line2['M']; !ok {
		return false
	}
	if _, ok := line2['S']; !ok {
		return false
	}
	return true
}

func (m matrix) inBounds(x, y int) bool {
	if x < 0 || y < 0 || x >= len(m) || y >= len(m[x]) {
		return false // out of bounds
	}
	return true
}

func (m matrix) testLocation(x, y int) int {
	hits := 0
	for _, dir := range directions {
		if m.testDirection(x, y, dir) {
			hits++
		}
	}
	return hits
}

func (m matrix) testDirection(x, y int, dir []int) bool {
	for _, char := range word {
		if !m.inBounds(x, y) {
			return false
		}
		if !(m[x][y] == char) {
			return false
		}
		x += dir[0]
		y += dir[1]
	}
	return true
}

func parseFile() [][]rune {
	file, err := os.Open("day4/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var out [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		out = append(out, line)
	}
	return out
}
