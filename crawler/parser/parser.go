package parser

import "regexp"

func ParseList(contents []byte, regex string) [][]byte {
	regexp := regexp.MustCompile(regex)
	submatch := regexp.FindAllSubmatch(contents, -1)
	var result [][]byte
	for _, m := range submatch {
		result = append(result, m[1:]...)
	}
	return result
}
