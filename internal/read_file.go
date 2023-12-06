package internal

import (
	"bufio"
	"os"
)

func Read(path string) ([]string, error) {
	result := make([]string, 0)

	file, err := os.Open(path)
	if err != nil {
		return result, err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}

	return result, nil
}
