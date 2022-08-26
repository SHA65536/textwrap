package textwrap

import "strings"

const defaultDelim string = " "
const defaultEnd string = "\n"

func WrapCustom(input string, size int, delim, end string) ([]string, error) {
	var results = []string{}
	var strs = strings.Split(input, end)
	for _, str := range strs {
		if len(str) <= size {
			results = append(results, str)
			continue
		}
		var curStr string
		words := strings.Split(str, delim)
		for _, word := range words {
			if len(curStr)+len(delim)+len(word) > size {
				if curStr != "" {
					results = append(results, curStr)
				}
				for len(word) > size {
					results = append(results, word[:size])
					word = word[size:]
				}
				curStr = word
			} else {
				if curStr == "" {
					curStr = word
				} else {
					curStr += delim + word
				}
			}
		}
		if curStr != "" {
			results = append(results, curStr)
		}
	}
	return results, nil
}

func Wrap(input string, size int) (string, error) {
	res, err := WrapCustom(input, size, defaultDelim, defaultEnd)
	if err != nil {
		return "", err
	}
	return strings.Join(res, "\n"), nil
}
