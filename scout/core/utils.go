package core

import "strings"

func CreateSlug(name string) string {
	return strings.ToLower(strings.ReplaceAll(name, " ", "-"))
}