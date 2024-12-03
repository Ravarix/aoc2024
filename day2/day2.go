package day2

import (
	"bufio"
	"fmt"
	"log"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	reports := parseFile()
	safeCount := 0
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		}
	}
	fmt.Println("safe reports: ", safeCount)

	safeCount = 0
	// Part 2
	for _, report := range reports {
		if isSafe(report) {
			safeCount++
		} else {
			slog.Debug(fmt.Sprintf("attempting error correction"))
			for i := 0; i < len(report); i++ {
				corrected := append(slices.Clone(report[:i]), report[i+1:]...)
				if isSafe(corrected) {
					safeCount++
					break
				}
			}
		}
	}

	fmt.Println("safe reports after correction: ", safeCount)

}

func isSafe(report []int) bool {
	slog.Debug(fmt.Sprintf("evaluating %v", report))
	safe := true
	previous := report[0]
	previousDelta := 0
	for _, val := range report[1:] {
		delta := val - previous
		if delta == 0 || delta > 3 || delta < -3 {
			safe = false
			slog.Debug(fmt.Sprintf("\n	invalid delta [%v, %v]\n", val, previous))
			break
		}
		if delta > 0 && previousDelta < 0 || delta < 0 && previousDelta > 0 {
			safe = false
			slog.Debug(fmt.Sprintf("\n	flipped delta [%v, %v]\n", delta, previousDelta))
			break
		}
		previous = val
		previousDelta = delta
	}
	if safe {
		slog.Debug(fmt.Sprintf(" safe"))
	}
	return safe
}

func parseFile() [][]int {
	file, err := os.Open("day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var out [][]int
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())
		var ints []int
		for _, s := range tokens {
			val, _ := strconv.Atoi(s)
			ints = append(ints, val)
		}
		out = append(out, ints)
	}

	return out
}
