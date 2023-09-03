package processor

import "postgresql/parser"

type Response struct {
	Query       string   `json:"query"`
	QueryTokens []string `json:"queryTokens"`
}

func QueryProcessor(query string) Response {
	var response Response

	response.Query = query
	response.QueryTokens = parser.Tokenize(query)

	return response
}
