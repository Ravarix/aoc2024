package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() {
	rules, lines := parseFile()

	goodLines, badLines := make(map[int]struct{}), make(map[int]struct{})
	for i, line := range lines {
		goodLine := true
		for j := range line.arr {
			if !checkRules(line, j, rules, false) {
				goodLine = false
				break
			}
		}
		if goodLine {
			goodLines[i] = struct{}{}
		} else {
			badLines[i] = struct{}{}
		}
	}
	fmt.Printf("The good lines are: %v\n", goodLines)
	sum := 0
	for lineIdx := range goodLines {
		middleIdx := (len(lines[lineIdx].arr) - 1) / 2
		sum += lines[lineIdx].arr[middleIdx]
	}
	fmt.Printf("The sum of all good lines is %v\n", sum)

	// Part 2
	sum = 0
	for lineIdx := range badLines {
		for j := range lines[lineIdx].arr {
			for !checkRules(lines[lineIdx], j, rules, true) { // keep using rule replace until success
			}
		}
		middleIdx := (len(lines[lineIdx].arr) - 1) / 2
		sum += lines[lineIdx].arr[middleIdx]
	}
	fmt.Printf("The sum of all fixed lines is %v\n", sum)
}

func checkRules(line lineData, idx int, rules map[int][]int, fix bool) bool {
	num := line.arr[idx]
	if befores, exists := rules[num]; exists { //do i have a rule with this number as an 'after'?
		for _, before := range befores {
			// does this 'before' exist in the line & have a greater loc than our current idx
			if beforeLoc, ok := line.locs[before]; ok && beforeLoc > idx {
				if fix {
					line.arr[idx] = before // swappy
					line.arr[beforeLoc] = num
					line.locs[before] = idx
					line.locs[num] = beforeLoc
				}
				return false // still return false, can't guarantee only 1 swap needed
			}
		}
	}
	return true
}

type lineData struct {
	arr  []int
	locs map[int]int
}

func parseFile() (rules map[int][]int, lines []lineData) {
	file, err := os.Open("day5/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rules = make(map[int][]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if len(str) == 0 {
			break
		}
		parts := strings.Split(str, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules[after] = append(rules[after], before)
	}
	lines = make([]lineData, 0)
	for scanner.Scan() {
		str := scanner.Text()
		var numLine []int
		locs := make(map[int]int)
		for i, val := range strings.Split(str, ",") {
			num, _ := strconv.Atoi(val)
			numLine = append(numLine, num)
			locs[num] = i
		}
		lines = append(lines, lineData{
			arr:  numLine,
			locs: locs,
		})
	}
	return rules, lines
}
