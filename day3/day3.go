package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Run() {
	regex := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	data := parseFile()
	matches := regex.FindAllStringSubmatch(data, -1)
	sum := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		sum += a * b
	}
	fmt.Println("sum: ", sum)

	// Part 2
	toggleRegex := regexp.MustCompile("(mul\\((\\d{1,3}),(\\d{1,3})\\)|don't\\(\\)|do\\(\\))")
	matches = toggleRegex.FindAllStringSubmatch(data, -1)
	sum = 0
	active := true
	for _, match := range matches {
		if match[0] == "do()" {
			active = true
		} else if match[0] == "don't()" {
			active = false
		} else {
			if active {
				a, _ := strconv.Atoi(match[2])
				b, _ := strconv.Atoi(match[3])
				sum += a * b
			}
		}
	}
	fmt.Println("sum: ", sum)
}

func parseFile() string {
	file, err := os.Open("day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
