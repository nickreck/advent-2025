package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err.Error())
	}
	var countP1 int
	var countP2 int
	for _, line := range strings.Split(string(file), ",") {
		bounds := strings.Split(line, "-")
		lowerBound, err := strconv.Atoi(bounds[0])
		if err != nil {
			log.Fatalf("failed to convert %s to a number: %s", bounds[0], err.Error())
		}
		upperBound, err := strconv.Atoi(bounds[1])
		if err != nil {
			log.Fatalf("failed to convert %s to a number: %s", bounds[1], err.Error())
		}
		for i := range upperBound - lowerBound + 1 {
			stringInt := strconv.Itoa(lowerBound + i)
			partOne := stringInt[:(len(stringInt) / 2)]
			partTwo := stringInt[(len(stringInt) / 2):]
			if partOne == partTwo {
				countP1 += lowerBound + i
			}
			if !isValidP2(stringInt) {
				countP2 += lowerBound + i
			}
		}
	}
	fmt.Println("Part 1:", countP1)
	fmt.Println("Part 2:", countP2)
}

func isValidP2(num string) bool {
	arr := []rune(num)
outer:
	for i := range len(arr) / 2 {
		if len(arr)%(i+1) != 0 {
			continue
		}
		chunks := slices.Collect(slices.Chunk(arr, i+1))
		for _, sli := range chunks[1:] {
			if !reflect.DeepEqual(chunks[0], sli) {
				continue outer
			}
		}
		return false
	}
	return true
}
