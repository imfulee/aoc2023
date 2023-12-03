package day1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getCalibrationValueFromNumbers(line string) (int, error) {
	firstNum := -1
	secondNum := -1

	for _, char := range line {
		number, err := strconv.Atoi(string(char))
		if err != nil {
			continue
		}

		if firstNum == -1 {
			firstNum = number
			secondNum = number
			continue
		}

		secondNum = number
	}

	if firstNum == -1 || secondNum == -1 {
		return -1, errors.New("did not get number")
	}

	return firstNum*10 + secondNum, nil
}

func ExecA() error {
	file, err := os.Open("data/day1/puzzleA.txt")
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
			break
		}

		if cValue, err := getCalibrationValueFromNumbers(line); err == nil {
			total += cValue
		} else {
			return err
		}
	}

	fmt.Println(total)

	return nil
}
