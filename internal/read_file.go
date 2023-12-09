package internal

import (
	"bufio"
	"errors"
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

type Reader struct {
	file *os.File
}

func (r *Reader) New(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return errors.Join(errors.New("unable to read file"), err)
	}
	defer file.Close()

	r.file = file

	return nil
}

func (r *Reader) Lines(lines chan string) error {
	defer close(lines)

	if r.file == nil {
		return errors.New("didn't initialise file")
	}

	fileScanner := bufio.NewScanner(r.file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines <- fileScanner.Text()
	}

	return nil
}
