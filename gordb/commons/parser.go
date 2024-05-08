package commons

import (
	"errors"
	"strings"
)

func Parse(query string) (Statement, error) {
	for _, command := range PossibleCommands {
		if strings.Contains(strings.ToLower(query), command) {
			return Statement{Type: command}, nil
		}
	}

	return Statement{}, errors.New("invalid commanding")
}
