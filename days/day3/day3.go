package day3

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Number struct {
	num    int
	points []Point
}

type Point struct {
	x int
	y int
}

func (p Point) isAdjacent(adjP Point) bool {
	deltaX := math.Abs(float64(p.x - adjP.x))
	deltaY := math.Abs(float64(p.y - adjP.y))

	if deltaX > 1 || deltaY > 1 {
		return false
	}

	return true
}

func Exec() error {
	file, err := os.Open("days/day3/input.txt")
	if err != nil {
		return errors.Join(errors.New("bad file read"), err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	numberRegex, err := regexp.Compile("[0-9]+")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}

	spCharRegex, err := regexp.Compile("[^.0-9\n]")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}

	numbers := make([]Number, 0)
	spChars := make([]Point, 0)

	lineNumber := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}

		numberIndexes := numberRegex.FindAllStringIndex(line, -1)
		for _, numberIndex := range numberIndexes {
			points := make([]Point, 0)
			for col := numberIndex[0]; col < numberIndex[1]; col++ {
				points = append(points, Point{x: col, y: lineNumber})
			}

			num, err := strconv.Atoi(line[numberIndex[0]:numberIndex[1]])
			if err != nil {
				return errors.Join(errors.New("cannot convert number"), err)
			}

			numbers = append(numbers, Number{num: num, points: points})
		}

		spCharIndexes := spCharRegex.FindAllStringIndex(line, -1)
		for _, spCharIndex := range spCharIndexes {
			x := spCharIndex[0]
			y := lineNumber

			spChars = append(spChars, Point{x: x, y: y})
		}

		lineNumber++
	}

	esSum := 0
	for _, number := range numbers {
		hasNumber := false
		for _, point := range number.points {
			if hasNumber {
				break
			}

			for _, spChar := range spChars {
				if !point.isAdjacent(spChar) {
					continue
				}

				esSum += number.num
				hasNumber = true
				break
			}
		}
	}

	fmt.Println(esSum)

	grSum := 0
	for _, spChar := range spChars {
		matches := make([]Number, 0)

		for _, number := range numbers {
			for _, point := range number.points {
				if !point.isAdjacent(spChar) {
					continue
				}

				matches = append(matches, number)
				break
			}
		}

		if len(matches) == 2 {
			grSum += matches[0].num * matches[1].num
		}
	}

	fmt.Println(grSum)

	return nil
}
