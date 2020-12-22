package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	hits := 1
	mr := []int{1, 3, 5, 7, 1}
	md := []int{1, 1, 1, 1, 2}

	for j := 0; j < len(mr); j++ {
		lhits := 0
		for i := 0; i < len(lines)-1; i += md[j] {
			if string(lines[i][i*mr[j]%len(lines[i])]) == "#" {
				lhits++
			}
		}
		hits = hits * lhits
	}

	fmt.Println(hits)
}
