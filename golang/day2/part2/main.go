package main

import (
	"advent2022/util"
	"fmt"
	"strings"
)

const (
	Rock = iota + 1
	Paper
	Scissor
	Win
	Lose
	Draw
)

type game struct {
	opponent int
	outcome  int
}

func parseLetter(letter string) int {
	return map[string]int{
		"A": Rock,
		"X": Lose,
		"B": Paper,
		"Y": Draw,
		"C": Scissor,
		"Z": Win,
	}[letter]
}

func parseGames(games []string) []game {
	var ret []game
	for _, g := range games {
		letters := strings.Split(g, " ")
		o := parseLetter(letters[0])
		p := parseLetter(letters[1])
		ret = append(ret, game{o, p})
	}
	return ret
}

func scoreGame(g game) int {
	losesAgainst := map[int]int{
		Rock:    Scissor,
		Paper:   Rock,
		Scissor: Paper,
	}
	winsAgainst := map[int]int{
		Rock:    Paper,
		Paper:   Scissor,
		Scissor: Rock,
	}

	var score int
	var player int

	switch g.outcome {
	case Draw:
		player = g.opponent
		score = 3
	case Win:
		player = winsAgainst[g.opponent]
		score = 6
	case Lose:
		player = losesAgainst[g.opponent]
		score = 0
	}

	return score + player
}

func main() {
	input := util.ReadInput()
	games := parseGames(strings.Split(input, "\n"))

	totalScore := 0

	for _, g := range games {
		totalScore += scoreGame(g)
	}

	fmt.Println(totalScore)
}
