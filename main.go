package main

import (
	"fmt"
	"github.com/GerardoHP/AdventOfCode2023/Common"
	"github.com/GerardoHP/AdventOfCode2023/Days"
)

func main() {
	err, strs := Common.ReadLines("inputs/day2_part2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	r1 := Days.Day2(strs)
	fmt.Println(r1)
}
