package internal

import "regexp"

var numberRegexInst *regexp.Regexp

func NumberRegex() (*regexp.Regexp, error) {
	if numberRegexInst == nil {
		numberRegex, err := regexp.Compile("[0-9]+")
		if err != nil {
			return nil, err
		}

		numberRegexInst = numberRegex
	}

	return numberRegexInst, nil
}
