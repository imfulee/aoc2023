package day8

import (
	"aoc2023/internal"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type paths struct {
	left  string
	right string
}

const start = "AAA"
const end = "ZZZ"

const left = "L"
const right = "R"

func NormalWalk(nodes map[string]paths, directions []string, currentNodeKey string) (int, error) {
	steps := 0
	directionLen := len(directions)
	currentDirectionIndex := 0

	for currentNodeKey != end {
		currentNode, hasNode := nodes[currentNodeKey]
		if !hasNode {
			return -1, fmt.Errorf("node %s not found", currentNodeKey)
		}

		currentDirection := directions[currentDirectionIndex]
		switch currentDirection {
		case left:
			currentNodeKey = currentNode.left
		case right:
			currentNodeKey = currentNode.right
		default:
			return -1, errors.New("direction not found")
		}

		currentDirectionIndex = (currentDirectionIndex + 1) % directionLen
		steps++
	}

	return steps, nil
}

func gcd(a, b int) int {
	if a > b {
		a, b = b, a
	}

	if a == 0 {
		return b
	}

	return gcd(a, b%a)
}

func lcm(numbers []int) (int, error) {
	numberLen := len(numbers)

	if numberLen < 2 {
		return -1, errors.New("not enough amounts of numbers")
	}

	lcm := 1
	for i := 0; i < numberLen; i++ {
		b := numbers[i]
		lcm = (lcm * b) / gcd(lcm, b)
	}

	return lcm, nil
}

func ghostWalker(nodes map[string]paths, directions []string, currentNodeKey string) (int, error) {
	steps := 0
	directionLen := len(directions)
	currentDirectionIndex := 0

	for string(currentNodeKey[2]) != "Z" {
		currentNode, hasNode := nodes[currentNodeKey]
		if !hasNode {
			return -1, fmt.Errorf("node %s not found", currentNodeKey)
		}

		currentDirection := directions[currentDirectionIndex]
		switch currentDirection {
		case left:
			currentNodeKey = currentNode.left
		case right:
			currentNodeKey = currentNode.right
		default:
			return -1, errors.New("direction not found")
		}

		currentDirectionIndex = (currentDirectionIndex + 1) % directionLen
		steps++
	}

	return steps, nil
}

func GhostWalk(nodes map[string]paths, directions []string) (int, error) {
	currentNodeKeys := make([]string, 0)
	for nodeKey := range nodes {
		if string(nodeKey[2]) == "A" {
			currentNodeKeys = append(currentNodeKeys, nodeKey)
		}
	}

	steps := make([]int, len(currentNodeKeys))

	for i := range currentNodeKeys {

		step, err := ghostWalker(nodes, directions, currentNodeKeys[i])
		if err != nil {
			return -1, errors.Join(errors.New("normal walk"), err)
		}

		steps[i] = step
	}

	lcmSteps, err := lcm(steps)
	if err != nil {
		return -1, errors.Join(errors.New("lcm error"), err)
	}

	return lcmSteps, nil
}

func Exec() error {
	lines, err := internal.Read("days/day8/input.txt")
	if err != nil {
		return errors.Join(errors.New("bad read file"), err)
	}

	if len(lines) < 2 {
		return errors.New("too few lines")
	}

	nodeRegex, err := regexp.Compile("[0-9A-Z]+")
	if err != nil {
		return errors.Join(errors.New("cannot compile node regex"), err)
	}

	directions := strings.Split(strings.Trim(lines[0], " "), "")
	if len(directions) < 1 {
		return errors.New("bad amount of directions")
	}

	nodes := make(map[string]paths)
	for _, line := range lines[2:] {
		if line == "" {
			continue
		}

		data := nodeRegex.FindAllString(line, -1)
		if len(data) != 3 {
			return errors.New("bad input, should have three nodes")
		}

		nodeKey := data[0]
		nodes[nodeKey] = paths{left: data[1], right: data[2]}
	}

	if steps, err := NormalWalk(nodes, directions, start); err != nil {
		return err
	} else {
		fmt.Println(steps)
	}

	if steps, err := GhostWalk(nodes, directions); err != nil {
		return err
	} else {
		fmt.Println(steps)
	}

	return nil
}
