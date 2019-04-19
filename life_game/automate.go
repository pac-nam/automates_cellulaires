package main

import (
	"time"
)
import "fmt"


func PrintGrid(grid [HEIGHT][WEIGHT]int8) {
	toPrint := make([]rune, (HEIGHT*WEIGHT) + WEIGHT)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				toPrint[(i * (WEIGHT + 1)) + j] = ALIVEB
			} else {
				toPrint[(i * (WEIGHT + 1)) + j] = DEADB
			}
		}
		toPrint[i * (WEIGHT + 1)] = '\n'
	}
	fmt.Println(string(toPrint))
}

func play(old [HEIGHT][WEIGHT]int8) {
	for {
		PrintGrid(old)
		time.Sleep(SLEEP)
		var toPrint [HEIGHT][WEIGHT]int8
		for y := 1; y < HEIGHT - 1; y++ {
			for x := 1; x < WEIGHT - 1; x++ {
				cpt := old[y-1][x-1] +
					old[y-1][x] +
					old[y-1][x+1] +
					old[y][x-1] +
					old[y][x+1] +
					old[y+1][x-1] +
					old[y+1][x] +
					old[y+1][x+1]
				if cpt == 3 {
					toPrint[y][x] = 1
				} else if old[y][x] == 1 && cpt == 2 {
					toPrint[y][x] = 1
				} else {
					toPrint[y][x] = 0
				}
			}
		}
		old = toPrint
	}
}