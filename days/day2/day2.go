package day2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
)

type Game struct {
	red   int
	green int
	blue  int
}

func intepretLine(line string) (gameID int, games []Game, err error) {
	games = make([]Game, 0)

	idAndGame := strings.Split(line, ":")
	if len(idAndGame) != 2 {
		err = errors.New("cannot split line into game id and games")
		return
	}

	wordAndId := strings.Split(idAndGame[0], " ")
	if len(idAndGame) != 2 {
		err = errors.New("cannot split line into game id and games")
		return
	}

	gameID, err = strconv.Atoi(wordAndId[1])
	if err != nil {
		err = errors.New("cannot convert game ID")
		return
	}

	turns := strings.Split(idAndGame[1], ";")
	for _, turn := range turns {
		if turn == "" {
			continue
		}

		game := Game{}

		rgbSplits := strings.Split(turn, ",")
		for _, rgbSplit := range rgbSplits {
			countColourTuple := strings.Split(strings.Trim(rgbSplit, " "), " ")
			if len(countColourTuple) != 2 {
				err = errors.New("cannot split count and colour")
				return
			}

			if count, cErr := strconv.Atoi(strings.Trim(countColourTuple[0], " ")); cErr == nil {
				colour := strings.Trim(countColourTuple[1], " ")
				switch colour {
				case "red":
					game.red = count
				case "blue":
					game.blue = count
				case "green":
					game.green = count
				default:
					err = errors.New("invalid string for colour")
				}

				if err != nil {
					return
				}
			} else {
				err = errors.Join(errors.New("cannot convert count"), cErr)
				return
			}
		}

		games = append(games, game)
	}

	return
}

func ExecA() error {
	file, err := os.Open("days/day2/day2A.txt")
	if err != nil {
		return errors.Join(errors.New("bad file read"), err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	gameIDSum := 0
	powerSum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}

		if gameID, games, err := intepretLine(line); err == nil && gameID != -1 {
			goodGame := true
			least := map[string]int{
				"red":   0,
				"green": 0,
				"blue":  0,
			}

			for _, game := range games {
				if game.red > redLimit || game.blue > blueLimit || game.green > greenLimit {
					goodGame = false
				}

				if game.red > least["red"] {
					least["red"] = game.red
				}

				if game.green > least["green"] {
					least["green"] = game.green
				}

				if game.blue > least["blue"] {
					least["blue"] = game.blue
				}
			}

			fmt.Println(least)

			if goodGame {
				gameIDSum += gameID
			}

			powerSum += least["red"] * least["blue"] * least["green"]
		} else {
			return errors.Join(errors.New("cannot intepret line"), err)
		}
	}

	fmt.Println(gameIDSum)
	fmt.Println(powerSum)

	return nil
}
