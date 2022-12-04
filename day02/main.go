package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	RockScore    = 1
	PaperScore   = 2
	ScissorScore = 3
	WinScore     = 6
	LoseScore    = 0
	DrawScore    = 3
)

func main() {
	f, err := os.Open("day02/input.txt")
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	total_score := 0
	for fileScanner.Scan() {
		round := strings.Split(fileScanner.Text(), " ")
		opponent := round[0]
		me := round[1]
		total_score += determineScore(me, opponent)
		//fmt.Println(round)
	}
	err = f.Close()
	if err != nil {
		return
	}
	fmt.Println(total_score)
}

func determineScore(me string, opponent string) int {
	/*
		A = Rock
		B = Paper
		C = Scissor

		X = Rock
		Y = Paper
		Z = Scissor
	*/
	score := 0
	switch opponent {
	case "A":
		if me == "X" {
			score += DrawScore
		} else if me == "Y" {
			score += WinScore
		} else if me == "Z" {
			score += LoseScore
		}
	case "B":
		if me == "X" {
			score += LoseScore
		} else if me == "Y" {
			score += DrawScore
		} else if me == "Z" {
			score += WinScore
		}
	case "C":
		if me == "X" {
			score += WinScore
		} else if me == "Y" {
			score += LoseScore
		} else if me == "Z" {
			score += DrawScore
		}
	}

	switch me {
	case "X":
		score += RockScore
	case "Y":
		score += PaperScore
	case "Z":
		score += ScissorScore
	}
	return score
}
