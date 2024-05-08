package query

type Where = [][]string
type Order = [][]string

type Query struct {
	Table string `json:"table" binding:"required"`
	Where `json:"where"`
	Order `json:"order"`
}

type Response struct {
	Performance float64                  `json:"performance"`
	Data        []map[string]interface{} `json:"data"`
}
