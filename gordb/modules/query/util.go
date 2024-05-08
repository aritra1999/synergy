package query

import (
	"gordb/commons"
	"slices"
)

func ValidateQuery(query Query) bool {
	if !ValidateTable(query.Table) || !ValidateWhere(query.Where) || !ValidateOrder(query.Order) {
		return false
	}

	return true
}

func ValidateTable(table string) bool {
	return true
}

func ValidateWhere(where Where) bool {
	for _, conditions := range where {
		if len(conditions) != 3 {
			return false
		}

		if !slices.Contains(commons.ValidWhereOperators, conditions[1]) {
			return false
		}
	}

	return true
}

func ValidateOrder(order Order) bool {
	return true
}
