package util

import (
	"strconv"
	"strings"
)

func ParseInt(in string) int {
	parsed, err := strconv.Atoi(in)

	if err != nil {
		panic("couldn't parse number")
	}

	return parsed
}

func ParseIntList(list string, separator string) []int {
	items := strings.Split(list, separator)
	var ret []int

	for _, item := range items {
		ret = append(ret, ParseInt(item))
	}

	return ret
}
