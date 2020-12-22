package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	lines, _ := ioutil.ReadFile("input")
	ls := strings.Split(string(lines), "\n")
	ll := make([]string, len(ls))
	lll := make([]string, len(ls))
	for i, l := range ls { // iterate through all entries and 'store' index
		_ = copy(ll, ls)
		for j, k := range append(ll[:i], ll[i+1:]...) { // take all other entries but the stored one, for reasons
			_ = copy(lll, ls)
			for _, v := range append(lll[:j], lll[j+1:]...) { // take all other entries but the stored one, for reasons
				// fmt.Println(l, k, v)
				if barrtoi(l)+barrtoi(k)+barrtoi(v) == 2020 {
					fmt.Println(barrtoi(l) * barrtoi(k) * barrtoi(v))
				}
			}
		}
	}
}

func barrtoi(in string) int {
	i, err := strconv.Atoi(in)
	if err != nil {
		return -1
	}
	return i
}
