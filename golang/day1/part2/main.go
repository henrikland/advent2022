package main

import (
	"advent2022/util"
	"fmt"
	"sort"
	"strings"
)

func elfCalories(elves []string) []int {
	var ret []int

	for _, elf := range elves {
		calories := util.ParseIntList(elf, "\n")
		ret = append(ret, util.ListSum(calories))
	}

	return ret
}

func main() {
	input := util.ReadInput()
	calorieTotals := elfCalories(strings.Split(input, "\n\n"))
	sort.Ints(calorieTotals)

	fmt.Println(calorieTotals[len(calorieTotals)-1] + calorieTotals[len(calorieTotals)-2] + calorieTotals[len(calorieTotals)-3])
}
