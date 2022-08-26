// TextWrap - A package made to simplify formatting text in a size limited space.

// Package textwrap does the main part of formatting the string
package textwrap

import (
	"fmt"
	"strings"
)

const (
	// defaultDelim is the default seperator between words
	defaultDelim string = " "
	// defaultEnd is the default end of line
	defaultEnd string = "\n"
)

// WrapCustom is the main function of the repo
// Given an input, and desired size
func WrapCustom(input string, size int, delim, end string) ([]string, error) {
	if size < 1 {
		return nil, fmt.Errorf("size should be larget than 0")
	}
	if len(delim) >= size {
		return nil, fmt.Errorf("delimeter length bigger than size")
	}
	var results = []string{}
	// Splitting according to line ending
	var strs = strings.Split(input, end)
	for _, str := range strs {
		// If a line fits inside the size, add it whole
		if len(str) <= size {
			results = append(results, str)
			continue
		}
		var curStr string
		words := strings.Split(str, delim)
		// Working on the words, word by word
		for _, word := range words {
			if len(curStr)+len(delim)+len(word) > size {
				// If the next word doesn't fit
				if curStr != "" {
					// If the current string isn't empty, add it to results
					results = append(results, curStr)
				}
				// Split up the next word until it fits
				for len(word) > size {
					results = append(results, word[:size])
					word = word[size:]
				}
				curStr = word
			} else {
				// If the next word does fit, add it to the current string
				if curStr == "" {
					curStr = word
				} else {
					curStr += delim + word
				}
			}
		}
		// Add the rest
		if curStr != "" {
			results = append(results, curStr)
		}
	}
	return results, nil
}

// Wrap calls WrapCustom with default paramets
// And returns the output joined by the line ending.
func Wrap(input string, size int) (string, error) {
	res, err := WrapCustom(input, size, defaultDelim, defaultEnd)
	if err != nil {
		return "", err
	}
	return strings.Join(res, "\n"), nil
}
