package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to read input file: %s", err.Error())
	}
	var partOneJoltage int
	var partTwoJoltage int
	start := time.Now()
	for line := range strings.SplitSeq(string(file), "\n") {
		for i := 9; i > 0; i-- {
			joltage = ""
			jolt, err := calculateJoltage(line, 2, i)
			if err == nil {
				partOneJoltage += jolt
				break
			}
		}
	}
	fmt.Println("Part 1:", partOneJoltage, "execution time:", time.Since(start))
	start = time.Now()
	for line := range strings.SplitSeq(string(file), "\n") {
		for i := 9; i > 0; i-- {
			joltage = ""
			jolt, err := calculateJoltage(line, 12, i)
			if err == nil {
				partTwoJoltage += jolt
				break
			}
		}
	}
	fmt.Println("Part 2:", partTwoJoltage, "execution time:", time.Since(start))
}

var joltage string

func calculateJoltage(line string, batteryCount, findJoltage int) (int, error) {
	if len(joltage) == batteryCount {
		return strconv.Atoi(joltage)
	}
	if findJoltage == 0 || len(line) < batteryCount-len(joltage) {
		return 0, errors.New("not found")
	}
	index := slices.Index([]rune(line), rune(strconv.Itoa(findJoltage)[0]))
	if index == -1 {
		return 0, errors.New("not found")
	}
	joltage += strconv.Itoa(findJoltage)
	for i := 9; i > 0; i-- {
		jolt, err := calculateJoltage(line[index+1:], batteryCount, i)
		if err == nil {
			return jolt, nil
		}
	}
	joltage = joltage[:len(joltage)-1]
	return 0, errors.New("not found")
}
