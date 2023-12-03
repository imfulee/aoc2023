package day1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var numberStrings = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func setNumbers(firstNum, lastNum, num int) (int, int) {
	if firstNum == -1 {
		return num, num
	}

	return firstNum, num
}

func getCalibrationValue(line string) (int, error) {
	lineLength := len(line)
	firstNum := -1
	lastNum := -1

	for indexStart := range line {
		number, err := strconv.Atoi(string(line[indexStart]))
		if err == nil {
			firstNum, lastNum = setNumbers(firstNum, lastNum, number)
			continue
		}

		for numberString := range numberStrings {
			indexEnd := indexStart + len(numberString)
			if indexEnd > lineLength {
				continue
			}

			subString := line[indexStart:indexEnd]
			if subString == numberString {
				number = numberStrings[subString]
				firstNum, lastNum = setNumbers(firstNum, lastNum, number)
				break
			}
		}
	}

	if firstNum == -1 || lastNum == -1 {
		return -1, errors.New("cannot find calibratiuon value")
	}

	return firstNum*10 + lastNum, nil
}

func ExecB() error {
	file, err := os.Open("data/day1/puzzleB.txt")
	if err != nil {
		return errors.Join(errors.New("bad file read"), err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	total := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}

		if cValue, err := getCalibrationValue(line); err == nil {
			total += cValue
		} else {
			return err
		}
	}

	fmt.Println(total)

	return nil
}
