package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	dat, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	ap := make([]map[string]string, 0)
	passes := strings.Split(string(dat), "\n\n")
	reg := *regexp.MustCompile("((.*?):(.*?))\\s")

	for _, v := range passes {
		pass := strings.ReplaceAll(v, "\n", " ")
		matches := reg.FindAllStringSubmatch(fmt.Sprintf("%s ", pass), -1)
		m := make(map[string]string)
		for _, k := range matches {
			m[k[2]] = k[3]
		}
		ap = append(ap, m)
	}

	req_fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} // missing cid is fine
	cvalid := 0
	pp.Println(ap)
	for _, v := range ap { // each map
		valid := true
		for _, r := range req_fields {
			if _, ok := v[r]; ok {
				valid = valid && ok
			} else {
				valid = false
			}
		}
		if valid && isvalid(v) {
			cvalid++
		}
	}
	fmt.Println(cvalid)
}

func isvalid(input map[string]string) bool {
	val := true
	for k, v := range input {
		switch k {
		case "byr":
			i, _ := strconv.Atoi(v)
			val = val && bt(1920, 2002, i)
			break

		case "iyr":
			i, _ := strconv.Atoi(v)
			val = val && bt(2010, 2020, i)
			break

		case "eyr":
			i, _ := strconv.Atoi(v)
			val = val && bt(2020, 2030, i)
			break

		case "hgt":
			if strings.Contains(v, "cm") {
				i, _ := strconv.Atoi(v[:len(v)-2])
				val = val && bt(150, 193, i)
			} else if strings.Contains(v, "in") {
				i, _ := strconv.Atoi(v[:len(v)-2])
				val = val && bt(59, 76, i)
			} else {
				val = false
			}
			break

		case "hcl":
			reg := *regexp.MustCompile("#[0-9a-fA-F]{6}$")
			val = val && reg.MatchString(v)
			break

		case "ecl":
			reg := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
			val = val && reg.Match([]byte(v))
			break

		case "pid":
			reg := *regexp.MustCompile("^[0-9]{9}$")
			val = val && reg.MatchString(v)
			break
		}
		fmt.Println(k, val)
	}

	return val
}

func bt(min, max, num int) bool {
	return num >= min && num <= max
}
