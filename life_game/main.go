package main

import (
	"math/rand"
	"time"
)

func InitGrid() [HEIGHT][WEIGHT]int8 {
	var grid[HEIGHT][WEIGHT]int8
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range grid {
		for j := range grid[i] {
			if rand.Int() % 2 == 1 {
				grid[i][j] = 1
			}
		}
	}
	return grid
}

func main() {
	grid := InitGrid()
	play(grid)
}