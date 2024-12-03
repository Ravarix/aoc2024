package day1

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	left, right := parseFile()
	slices.Sort(left)
	slices.Sort(right)
	dist := 0
	for i := range len(left) {
		dist += int(math.Abs(float64(left[i] - right[i]))) // bruh why doesn't go have math.Abs for ints
	}
	println("distance: ", dist)

	// STEP 2
	counts := make(map[int]int)
	for _, val := range right {
		counts[val] += 1
	}
	similarity := 0
	for _, val := range left {
		if count, ok := counts[val]; ok {
			similarity += val * count
		}
	}
	println("similarity: ", similarity)
}

func parseFile() ([]int, []int) {
	file, err := os.Open("day1/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var out1 []int
	var out2 []int

	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		left, _ := strconv.Atoi(tokens[0])
		right, _ := strconv.Atoi(tokens[1])
		out1 = append(out1, left)
		out2 = append(out2, right)
	}

	return out1, out2
}
