package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func main() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err.Error())
	}
	grid := strings.Split(string(fileBytes), "\n")
	count1 := traverseGrid(slices.Clone(grid))
	var count2 int
	var currentIteration = 1
	for currentIteration != 0 {
		cache = [][]int{}
		currentIteration = traverseGrid(grid)
		count2 += currentIteration
		removePaper(grid)
	}
	fmt.Println("Part 1:", count1)
	fmt.Println("Part 2:", count2)
}

var cache [][]int

func traverseGrid(grid []string) int {
	var count int
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != '@' {
				continue
			}
			if isValid(grid, r, c) {
				count++
				cache = append(cache, []int{r, c})
			}
		}
	}
	return count
}
func isValid(grid []string, r, c int) bool {
	var countRolls int8
	if !isOutOfBounds(grid, r-1, c) && grid[r-1][c] == '@' {
		countRolls++
	}
	if !isOutOfBounds(grid, r-1, c-1) && grid[r-1][c-1] == '@' {
		countRolls++
	}
	if !isOutOfBounds(grid, r-1, c+1) && grid[r-1][c+1] == '@' {
		countRolls++
	}
	if !isOutOfBounds(grid, r+1, c) && grid[r+1][c] == '@' {
		countRolls++
	}
	if !isOutOfBounds(grid, r+1, c-1) && grid[r+1][c-1] == '@' {
		countRolls++
	}
	if !isOutOfBounds(grid, r+1, c+1) && grid[r+1][c+1] == '@' {
		countRolls++
	}
	if !isOutOfBounds(grid, r, c-1) && grid[r][c-1] == '@' {
		countRolls++
	}
	if !isOutOfBounds(grid, r, c+1) && grid[r][c+1] == '@' {
		countRolls++
	}
	return countRolls < 4
}

func isOutOfBounds(grid []string, r, c int) bool {
	return r < 0 || r == len(grid) || c < 0 || c == len(grid[r])
}

func removePaper(grid []string) {
	for _, point := range cache {
		grid[point[0]] = string(slices.Replace([]rune(grid[point[0]]), point[1], point[1]+1, 'X'))
	}
}
