package main

import (
	"advent2022/util"
	"fmt"
	"strings"
)

var Rock int = 1
var Paper int = 2
var Scissor int = 3

type game struct {
	opponent int
	player   int
}

func parseLetter(letter string) int {
	return map[string]int{
		"A": Rock,
		"X": Rock,
		"B": Paper,
		"Y": Paper,
		"C": Scissor,
		"Z": Scissor,
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

	score := 0

	if g.opponent == g.player {
		score = 3
	} else if g.opponent == losesAgainst[g.player] {
		score = 6
	}

	return score + g.player
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
