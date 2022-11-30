package util

import "os"

func ReadInput() string {
	args := os.Args[1:]

	if len(args) != 1 {
		panic("Input arg missing")
	}

	result, err := os.ReadFile(args[0])

	if err != nil {
		panic(err)
	}

	return string(result)
}
