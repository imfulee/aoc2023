package day6

import (
	"aoc2023/internal"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Exec() error {
	lines, err := internal.Read("days/day6/input.txt")
	if err != nil {
		return errors.Join(errors.New("bad read file"), err)
	}

	numRegex, err := internal.NumberRegex()
	if err != nil {
		return errors.Join(errors.New("unable to init regex"), err)
	}

	timeStrs := numRegex.FindAllString(lines[0], -1)
	distStrs := numRegex.FindAllString(lines[1], -1)

	if len(timeStrs) != len(distStrs) {
		return errors.New("bad input")
	}

	times := make([]float64, len(timeStrs))
	for i := range timeStrs {
		time, err := strconv.Atoi(timeStrs[i])
		if err != nil {
			return errors.Join(errors.New("cannot convert time"), err)
		}

		times[i] = float64(time)
	}

	/**
	 * t => total time
	 * d => distance needed to beat
	 * x => time hold on button
	 * (t - x) * x > d
	 * which is a quadratic and all numbers within the two roots would be the answer
	 */
	dists := make([]float64, len(distStrs))
	for i := range distStrs {
		dist, err := strconv.Atoi(distStrs[i])
		if err != nil {
			return errors.Join(errors.New("cannot convert distance"), err)
		}

		dists[i] = float64(dist)
	}

	multiple := 1
	for i := 0; i < len(times); i++ {
		r1 := (times[i] + math.Sqrt((times[i]*times[i])-(4*dists[i]))) / 2
		r2 := (times[i] - math.Sqrt((times[i]*times[i])-(4*dists[i]))) / 2

		// reverting the ceiling and floor because for whole integers like r1=20, r2=10, would not beat the distance
		ways := int(math.Ceil(r1)) - int(math.Floor(r2)) - 1
		multiple *= ways
	}

	fmt.Println(multiple)

	timeStr := strings.Join(timeStrs, "")
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		return errors.Join(errors.New("cannot convert time"), err)
	}

	distStr := strings.Join(distStrs, "")
	dist, err := strconv.Atoi(distStr)
	if err != nil {
		return errors.Join(errors.New("cannot convert distance"), err)
	}

	r1 := (float64(time) + math.Sqrt((float64(time)*float64(time))-(4*float64(dist)))) / 2
	r2 := (float64(time) - math.Sqrt((float64(time)*float64(time))-(4*float64(dist)))) / 2

	ways := int(math.Ceil(r1)) - int(math.Floor(r2)) - 1

	fmt.Println(ways)

	return nil
}
