package day9

import (
	"aoc2023/internal"
	"errors"
	"fmt"
	"strconv"
)

func differences(numbers []int) []int {
	result := make([]int, len(numbers)-1)

	for i, end := 0, len(numbers)-1; i < end; i++ {
		result[i] = numbers[i+1] - numbers[i]
	}

	return result
}

func allZeros(numbers []int) bool {
	for _, number := range numbers {
		if number != 0 {
			return false
		}
	}

	return true
}

func predict(numbers []int) int {
	result := 0

	stack := make([][]int, 0)
	stack = append(stack, numbers)
	for !allZeros(numbers) {
		numbers = differences(numbers)
		stack = append(stack, numbers)
	}

	stackLastIndex := len(stack) - 1
	stack[stackLastIndex] = append(stack[stackLastIndex], 0)
	for i := len(stack) - 2; i >= 0; i-- {
		layer := stack[i]
		prevLayer := stack[i+1]

		layer = append(layer, prevLayer[len(prevLayer)-1]+layer[len(layer)-1])

		stack[i] = layer
	}

	result = stack[0][len(stack[0])-1]

	return result
}

func predictBackwards(numbers []int) int {
	result := 0

	stack := make([][]int, 0)
	stack = append(stack, numbers)
	for !allZeros(numbers) {
		numbers = differences(numbers)
		stack = append(stack, numbers)
	}

	stackLastIndex := len(stack) - 1
	stack[stackLastIndex] = append(stack[stackLastIndex], 0)
	for i := len(stack) - 2; i >= 0; i-- {
		layer := stack[i]
		prevLayer := stack[i+1]

		layer = append([]int{layer[0] - prevLayer[0]}, layer...)

		stack[i] = layer
	}

	result = stack[0][0]

	return result
}

func Exec() error {
	lines, err := internal.Read("days/day9/input.txt")
	if err != nil {
		return errors.Join(errors.New("bad read file"), err)
	}

	regex, err := internal.NumberRegex()
	if err != nil {
		return errors.Join(errors.New("cannot get regex"), err)
	}

	sum := 0
	backwardSum := 0
	for _, line := range lines {
		foundNumbers := regex.FindAllString(line, -1)

		numbers := make([]int, len(foundNumbers))
		for i, foundNumber := range foundNumbers {
			num, err := strconv.Atoi(foundNumber)
			if err != nil {
				return errors.Join(errors.New("cannot convert number"), err)
			}

			numbers[i] = num
		}

		sum += predict(numbers)
		backwardSum += predictBackwards(numbers)
	}

	fmt.Println(sum)
	fmt.Println(backwardSum)

	return nil
}
