package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	part1("input.txt")
	part2("input.txt")
}

func part1(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		length := len(input)
		part1 := input[:length/2]
		part2 := input[length/2 : length]
		res += eval(part1, part2)
	}
	fmt.Println("Part 1: ", res)
}

func part2(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	res2 := 0

	input, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")

	for i := 0; i < len(lines); i = i + 3 {
		res2 += eval2(lines[i], lines[i+1], lines[i+2])
	}
	fmt.Println("Part 2: ", res2)
}

func eval(part1 string, part2 string) int {
	for i := 0; i < len(part1); i++ {
		for j := 0; j < len(part2); j++ {
			if string(part1[i]) == string(part2[j]) {
				if 65 <= int(part1[i]) && int(part1[i]) <= 90 {
					return int(part1[i] - 38)
				} else {
					return int(part1[i] - 96)
				}
			}
		}
	}
	return -1
}

func eval2(s1 string, s2 string, s3 string) int {
	cs := ""
	if len(s1) >= len(s2) && len(s1) >= len(s3) {
		cs = s1
	} else if len(s2) >= len(s1) && len(s2) >= len(s3) {
		cs = s2
	} else {
		cs = s3
	}

	for i := 0; i < len(cs); i++ {
		if strings.Contains(s1, string(cs[i])) && strings.Contains(s2, string(cs[i])) && strings.Contains(s3, string(cs[i])) {
			if 65 <= int(cs[i]) && int(cs[i]) <= 90 {
				return int(cs[i] - 38)
			} else {
				return int(cs[i] - 96)
			}
		}
	}
	return 0
}
