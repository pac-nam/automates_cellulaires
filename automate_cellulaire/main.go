package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", os.Args[0], "<rule>")
		return
	}
	rule, err := strconv.Atoi(os.Args[1])
	if err != nil || rule > 255 || rule < 0 {
		fmt.Println("rule have to be between 0 and 255")
		return
	}
	var tab [8]bool
	cpt := 7
	for mask := 1; mask < 256; mask*=2 {
		if rule & mask > 0 {
			tab[cpt] = true
		}
		cpt--
	}
	PrintRules(tab)
	automate(tab)
}