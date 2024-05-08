package mutation

import (
	"errors"
	"strings"
)

type Response struct {
	Response    string    `json:"response"`
	Query       string    `json:"query"`
	Statement   Statement `json:"statement"`
	QueryTokens []string  `json:"queryTokens"`
}

type Statement struct {
	Type string
}

var POSSIBLE_COMMANDS = []string{"create", "insert", "select", "update", "delete", "alter", "drop", "truncate", "describe", "show"}

func PrepareStatement(query string) (Statement, error) {
	for _, command := range POSSIBLE_COMMANDS {
		if strings.Contains(strings.ToLower(query), command) {
			return Statement{Type: command}, nil
		}
	}

	return Statement{}, errors.New("invalid commanding")
}

func MutationProcessor(table string) (Response, error) {
	var response Response
	query := "SELECT * FROM"
	statement, err := PrepareStatement(query)

	if err != nil {
		return response, err
	}

	response.Query = query
	response.Statement = statement

	return response, nil
}
