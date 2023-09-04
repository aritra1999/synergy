package processor

import "relational-db/parser"

type ResponseSuccess struct {
	Query       string   `json:"query"`
	QueryTokens []string `json:"queryTokens"`
	Response	string   `json:"response"`
}

type ResponseError struct {
	Query string `json:"query"`
	Error string `json:"error"`
}

func QueryProcessor(query string) Response {
	var response Response

	response.Query = query
	response.QueryTokens = parser.Tokenize(query)

	return response
}
