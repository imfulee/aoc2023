package main

import (
	"aoc2023/days/day1"
	"fmt"
)

func main() {
	if err := day1.ExecA(); err != nil {
		fmt.Println(err)
	}

	if err := day1.ExecB(); err != nil {
		fmt.Println(err)
	}
}
