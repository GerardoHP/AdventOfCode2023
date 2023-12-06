package main

import (
	"fmt"
	"github.com/GerardoHP/AdventOfCode2023/Common"
	"github.com/GerardoHP/AdventOfCode2023/Days"
)

func main() {
	err, strs := Common.ReadLines("inputs/input_day1_2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	r1 := Days.Day1(strs)
	fmt.Println(r1)
}
