package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(dat), "\n")
	re := *regexp.MustCompile(`^(\d+)-(\d+) (.): (.*?)$`)
	cnt := 0

	for _, v := range lines {
		res := re.FindAllStringSubmatch(v, -1)
		if res != nil {
			min, _ := strconv.Atoi(res[0][1])
			max, _ := strconv.Atoi(res[0][2])
			chr := res[0][3]
			str := res[0][4]

			if (string(str[min-1]) == chr) != (string(str[max-1]) == chr) {
				cnt++
			}

		}
	}

	fmt.Println(cnt)
}
