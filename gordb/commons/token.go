package commons

import "strings"

func Tokenize(query string) []string {
	return strings.Split(query, " ")
}
