package util

import "strconv"

func ParseInt(in string) int {
	parsed, err := strconv.Atoi(in)

	if err != nil {
		panic("couldn't parse number")
	}

	return parsed
}
