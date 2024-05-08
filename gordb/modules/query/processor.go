package query

import "time"

func QueryProcessor(query Query) (Response, error) {
	var response Response
	startProcessing := time.Now()

	data := map[string]interface{}{
		"table": query.Table,
	}

	response.Data = []map[string]interface{}{data, data}

	response.Performance = float64(time.Since(startProcessing).Nanoseconds())
	return response, nil
}
