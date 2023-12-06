package day4

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"

	sets "github.com/deckarep/golang-set/v2"
)

type Card struct {
	CardNumber   int
	Copies       int
	WinningNums  sets.Set[int]
	SelectedNums sets.Set[int]
}

func Exec() error {
	file, err := os.Open("days/day4/input.txt")
	if err != nil {
		return errors.Join(errors.New("bad file read"), err)
	}
	defer file.Close()

	numberRegex, err := regexp.Compile("[0-9]+")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	cards := make([]Card, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}

		splitedStr := strings.Split(line, ":")
		if len(splitedStr) != 2 {
			return errors.New("cannot split string by ':'")
		}

		cardNumber, err := strconv.Atoi(numberRegex.FindString(splitedStr[0]))
		if err != nil {
			return errors.New("cannot convert card number")
		}

		splitedStr = strings.Split(splitedStr[1], "|")
		if len(splitedStr) != 2 {
			return errors.New("cannot split string by '|'")
		}

		winningNums := sets.NewSet[int]()
		for _, winningNumStr := range numberRegex.FindAllString(splitedStr[0], -1) {
			winningNum, err := strconv.Atoi(winningNumStr)
			if err != nil {
				return errors.New("cannot convert winning number")
			}

			winningNums.Add(winningNum)
		}

		selectedNums := sets.NewSet[int]()
		for _, selectedNumStr := range numberRegex.FindAllString(splitedStr[1], -1) {
			selectedNum, err := strconv.Atoi(selectedNumStr)
			if err != nil {
				return errors.New("cannot convert winning number")
			}

			selectedNums.Add(selectedNum)
		}

		cards = append(cards, Card{CardNumber: cardNumber, Copies: 1, WinningNums: winningNums, SelectedNums: selectedNums})
	}

	pointSum := 0
	copiesSum := 0
	for cardIndex, card := range cards {
		matches := card.SelectedNums.Intersect(card.WinningNums).Cardinality()

		copiesSum += card.Copies

		if matches == 0 {
			continue
		}

		pointSum += int(math.Pow(2, float64(matches-1)))

		start := cardIndex + 1
		end := cardIndex + 1 + matches
		for wonCardIndex := start; wonCardIndex < end; wonCardIndex++ {
			cards[wonCardIndex].Copies += card.Copies
		}
	}

	fmt.Println(pointSum)
	fmt.Println(copiesSum)

	return nil
}
