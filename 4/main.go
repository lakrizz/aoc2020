package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/k0kubun/pp"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	regex := *regexp.MustCompile(`^(?:.+[\n\r])+`)
	fmt.Println(string(dat))
	passports := regex.FindAllStringSubmatch(string(dat), -1)
	pp.Println(passports)

	for _, v := range passports {
		pp.Println(v)
	}
}
