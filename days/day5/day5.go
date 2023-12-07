package day5

import (
	"aoc2023/internal"
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"sync"
)

type Range struct {
	start int
	end   int
}

type RangePair struct {
	source      Range
	destination Range
}

func translate(x int, rps *[]RangePair) (translated int) {
	isTranslated := false

	for _, rp := range *rps {
		if rp.source.start <= x && x <= rp.source.end {
			diff := x - rp.source.start
			translated = rp.destination.start + diff

			isTranslated = true
			break
		}
	}

	if !isTranslated {
		return x
	}

	return
}

var (
	seedSoilMap            = make([]RangePair, 0)
	soilFertilzerMap       = make([]RangePair, 0)
	fertilizerWaterMap     = make([]RangePair, 0)
	waterLightMap          = make([]RangePair, 0)
	lightTemperatureMap    = make([]RangePair, 0)
	temperatureHumidityMap = make([]RangePair, 0)
	humidityLocationMap    = make([]RangePair, 0)
)

func getMappings(lines []string, m *[]RangePair, regex *regexp.Regexp) error {
	found := false
	for _, line := range lines {
		if !found {
			if regex.FindString(line) != "" {
				found = true
				continue
			}

			continue
		}

		if found && line == "" {
			break
		}

		nRegex, err := internal.NumberRegex()
		if err != nil {
			return errors.Join(errors.New("unable to init regex"), err)
		}

		numbers := nRegex.FindAllString(line, -1)
		if len(numbers) != 3 {
			return errors.New("incorrect number of numbers in line")
		}

		srcStart, err := strconv.Atoi(numbers[1])
		if err != nil {
			return errors.Join(errors.New("cannot convert source start number"), err)
		}

		desStart, err := strconv.Atoi(numbers[0])
		if err != nil {
			return errors.Join(errors.New("cannot convert destination start number"), err)
		}

		rpRange, err := strconv.Atoi(numbers[2])
		if err != nil {
			return errors.Join(errors.New("cannot convert range number"), err)
		}

		*m = append(*m, RangePair{source: Range{start: srcStart, end: srcStart + rpRange - 1}, destination: Range{start: desStart, end: desStart + rpRange - 1}})
	}

	return nil
}

func transform(seed int) (location int) {
	soil := translate(seed, &seedSoilMap)
	fertilizer := translate(soil, &soilFertilzerMap)
	water := translate(fertilizer, &fertilizerWaterMap)
	light := translate(water, &waterLightMap)
	temperature := translate(light, &lightTemperatureMap)
	humidity := translate(temperature, &temperatureHumidityMap)
	location = translate(humidity, &humidityLocationMap)

	return

}

func Exec() error {
	lines, err := internal.Read("days/day5/input.txt")
	if err != nil {
		return errors.Join(errors.New("read file error"), err)
	}

	numberRegex, err := regexp.Compile("[0-9]+")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}

	seedStrs := numberRegex.FindAllString(lines[0], -1)
	seeds := make([]int, 0)
	for _, seedStr := range seedStrs {
		seed, err := strconv.Atoi(seedStr)
		if err != nil {
			return errors.Join(errors.New("cannot convert seed string"), err)
		}

		seeds = append(seeds, seed)
	}

	regex, err := regexp.Compile("seed-to-soil map:")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}
	if err = getMappings(lines, &seedSoilMap, regex); err != nil {
		return errors.Join(errors.New("cannot create seedSoilMap"), err)
	}

	regex, err = regexp.Compile("soil-to-fertilizer map:")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}
	if err = getMappings(lines, &soilFertilzerMap, regex); err != nil {
		return errors.Join(errors.New("cannot create soilFertilzerMap"), err)
	}

	regex, err = regexp.Compile("fertilizer-to-water map:")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}
	if err = getMappings(lines, &fertilizerWaterMap, regex); err != nil {
		return errors.Join(errors.New("cannot create fertilizerWaterMap"), err)
	}

	regex, err = regexp.Compile("water-to-light map:")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}
	if err = getMappings(lines, &waterLightMap, regex); err != nil {
		return errors.Join(errors.New("cannot create waterLightMap"), err)
	}

	regex, err = regexp.Compile("water-to-light map:")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}
	if err = getMappings(lines, &waterLightMap, regex); err != nil {
		return errors.Join(errors.New("cannot create waterLightMap"), err)
	}

	regex, err = regexp.Compile("light-to-temperature map:")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}
	if err = getMappings(lines, &lightTemperatureMap, regex); err != nil {
		return errors.Join(errors.New("cannot create lightTemperatureMap"), err)
	}

	regex, err = regexp.Compile("temperature-to-humidity map:")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}
	if err = getMappings(lines, &temperatureHumidityMap, regex); err != nil {
		return errors.Join(errors.New("cannot create temperatureHumidityMap"), err)
	}

	regex, err = regexp.Compile("humidity-to-location map:")
	if err != nil {
		return errors.Join(errors.New("cannot create regex"), err)
	}
	if err = getMappings(lines, &humidityLocationMap, regex); err != nil {
		return errors.Join(errors.New("cannot create humidityLocationMap"), err)
	}

	lowestLocNum := math.MaxInt64
	for _, seed := range seeds {
		location := transform(seed)

		if location < lowestLocNum {
			lowestLocNum = location
		}
	}

	fmt.Println(lowestLocNum)

	if len(seeds)%2 != 0 {
		return errors.New("bad amount of seeds")
	}

	var wg sync.WaitGroup
	lowChan := make(chan int, len(seeds)/2)

	for sp := 0; sp < len(seeds); sp += 2 {
		start := seeds[sp]
		length := seeds[sp+1]

		wg.Add(1)

		go func(c chan int, s int, l int) {
			lowest := math.MaxInt64
			for seed, end := s, s+l; seed < end; seed++ {
				location := transform(seed)

				if location < lowest {
					lowest = location
				}
			}

			c <- lowest
			wg.Done()
		}(lowChan, start, length)
	}

	go func() {
		wg.Wait()
		close(lowChan)
	}()

	lowestLocNum = math.MaxInt64
	for lowest := range lowChan {
		if lowest < lowestLocNum {
			lowestLocNum = lowest
		}
	}

	fmt.Println(lowestLocNum)

	return nil
}
