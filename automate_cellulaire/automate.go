package main

import (
	"fmt"
	"strings"
)

const (
	SLEEP = 1
	CYCLE = 60
	HALFSIZE = 75
	SIZE = (HALFSIZE*2) + 1
	ALIVES = "█"
	ALIVEB = '█'
	DEADS = "."
	DEADB = '.'
)

func PrintRules(rule [8]bool) {
	var char, str string
	for _, r := range rule {
		if r {
			char = ALIVES
		} else {
			char = DEADS
		}
		str += "   " + char
	}
	fmt.Println("███ ██. █.█ █.. .██ .█. ..█ ...")
	fmt.Println(str[2:]+"\n")
}

func automate(rule [8]bool) {
	var print, old []rune
	start := strings.Repeat(DEADS, HALFSIZE) + ALIVES + strings.Repeat(DEADS, HALFSIZE)
	print = make([]rune, SIZE)
	old = make([]rune, SIZE)
	print[0], print[SIZE - 1] = DEADB, DEADB
	old = []rune(start)
	for i := 0; i < CYCLE; i++ {
		fmt.Println(string(old))
		for index := 1; index < SIZE - 1; index++ {
			if old[index - 1] == ALIVEB && old[index] == ALIVEB && old[index + 1] == ALIVEB {
				if rule[0] {
					print[index] = ALIVEB
				} else {
					print[index] = DEADB
				}
			} else if old[index - 1] == ALIVEB && old[index] == ALIVEB && old[index + 1] == DEADB {
				if rule[1] {
					print[index] = ALIVEB
				} else {
					print[index] = DEADB
				}
			} else if old[index - 1] == ALIVEB && old[index] == DEADB && old[index + 1] == ALIVEB {
				if rule[2] {
					print[index] = ALIVEB
				} else {
					print[index] = DEADB
				}
			} else if old[index - 1] == ALIVEB && old[index] == DEADB && old[index + 1] == DEADB {
				if rule[3] {
					print[index] = ALIVEB
				} else {
					print[index] = DEADB
				}
			} else if old[index - 1] == DEADB && old[index] == ALIVEB && old[index + 1] == ALIVEB {
				if rule[4] {
					print[index] = ALIVEB
				} else {
					print[index] = DEADB
				}
			} else if old[index - 1] == DEADB && old[index] == ALIVEB && old[index + 1] == DEADB {
				if rule[5] {
					print[index] = ALIVEB
				} else {
					print[index] = DEADB
				}
			} else if old[index - 1] == DEADB && old[index] == DEADB && old[index + 1] == ALIVEB {
				if rule[6] {
					print[index] = ALIVEB
				} else {
					print[index] = DEADB
				}
			} else if old[index - 1] == DEADB && old[index] == DEADB && old[index + 1] == DEADB {
				if rule[7] {
					print[index] = ALIVEB
				} else {
					print[index] = DEADB
				}
			}
				}
		for tmp := 0; tmp < SIZE; tmp++ {
			old[tmp] = print[tmp]
		}
	}
}