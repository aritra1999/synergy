package parser

import "strings"

func Tokenize(query string) []string {
	return strings.Split(query, " ")
}
