package day10

import (
	"aoc2023/internal"
	"errors"
	"fmt"
	"log"
	"slices"
	"strings"
)

var startingCharacter = "S"

type Position struct {
	x int
	y int
}

func samePosition(p1 Position, p2 Position) bool {
	return p1.x == p2.x && p1.y == p2.y
}

const (
	DIRECTION_UP    = "up"
	DIRECTION_DOWN  = "down"
	DIRECTION_LEFT  = "left"
	DIRECTION_RIGHT = "right"
)

func Exec() error {
	lines, err := internal.Read("days/day10/input.txt")
	if err != nil {
		return errors.Join(errors.New("bad read file"), err)
	}

	grid := make([][]string, 0)
	for _, line := range lines {
		characters := strings.Split(line, "")
		horizontal := make([]string, 0)
		horizontal = append(horizontal, characters...)
		grid = append(grid, horizontal)
	}

	startingPosition := Position{}
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == startingCharacter {
				startingPosition.x = x
				startingPosition.y = y
			}
		}
	}

	direction := ""
	moves := 0
	currentPosition := startingPosition

	if currentPosition.x < len(grid[0])-1 && slices.Contains([]string{"-", "J", "7"}, grid[currentPosition.y][currentPosition.x+1]) {
		direction = DIRECTION_RIGHT
	} else if currentPosition.y < len(grid)-1 && slices.Contains([]string{"|", "J", "L"}, grid[currentPosition.y+1][currentPosition.x]) {
		direction = DIRECTION_DOWN
	} else if currentPosition.x > 0 && slices.Contains([]string{"-", "F", "L"}, grid[currentPosition.y][currentPosition.x-1]) {
		direction = DIRECTION_LEFT
	} else if currentPosition.y > 0 && slices.Contains([]string{"|", "F", "7"}, grid[currentPosition.y-1][currentPosition.x]) {
		direction = DIRECTION_UP
	} else {
		log.Fatalln("Cannot find way out from starting position")
	}

	for {
		switch direction {
		case DIRECTION_DOWN:
			currentPosition.y++
		case DIRECTION_UP:
			currentPosition.y--
		case DIRECTION_LEFT:
			currentPosition.x--
		case DIRECTION_RIGHT:
			currentPosition.x++
		}

		switch grid[currentPosition.y][currentPosition.x] {
		case "L":
			if direction == DIRECTION_DOWN {
				direction = DIRECTION_RIGHT
			} else {
				direction = DIRECTION_UP
			}
		case "J":
			if direction == DIRECTION_DOWN {
				direction = DIRECTION_LEFT
			} else {
				direction = DIRECTION_UP
			}
		case "7":
			if direction == DIRECTION_UP {
				direction = DIRECTION_LEFT
			} else {
				direction = DIRECTION_DOWN
			}
		case "F":
			if direction == DIRECTION_UP {
				direction = DIRECTION_RIGHT
			} else {
				direction = DIRECTION_DOWN
			}
		}

		moves++

		if samePosition(currentPosition, startingPosition) {
			break
		}
	}

	fmt.Println(moves / 2)

	return nil
}
