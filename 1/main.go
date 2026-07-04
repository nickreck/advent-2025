package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("error opening input text file: %s", err.Error())
	}

	position := 50

	var landedOnZero int
	var passedZero int
	for _, line := range strings.Split(string(file), "\n") {
		var direction int
		switch line[0] {
		case 'L':
			direction = -1
		case 'R':
			direction = 1
		}

		clicks, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("failed to convert %s to integer: %s", line[1:], err.Error())
		}

		for range clicks {
			position += direction
			switch position {
			case -1:
				position = 99
			case 100:
				position = 0
			}
			if position == 0 {
				passedZero++
			}
		}
		if position == 0 {
			landedOnZero++
		}
	}
	fmt.Println("Part 1:", landedOnZero)
	fmt.Println("Part 2:", passedZero)
}
