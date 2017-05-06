package vars

import (
	"regexp"
	"strings"
)

// Parse parses a string to look for vars
func Parse(str string) (map[string]string, error) {
	data := make(map[string]string)
	regex, err := regexp.Compile(`\s?[a-zA-Z]+\s?=\s?\w+`)

	if err != nil {
		return data, err
	}

	matches := regex.FindAllString(str, -1)
	for _, match := range matches {
		s := strings.Split(match, "=")
		key, value := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
		data[key] = value
	}
	return data, nil
}
