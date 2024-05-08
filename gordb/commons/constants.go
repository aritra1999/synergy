package commons

// Can be changed to a Set later for performance enhancements
var ValidWhereOperators = []string{"=", "!=", ">", "<", ">=", "<=", "like", "in", "between", "is", "is not"}

var PossibleCommands = []string{"create", "insert", "select", "update", "delete", "alter", "drop", "truncate", "describe", "show"}
