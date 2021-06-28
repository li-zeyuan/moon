package schema

type Schema struct {
	fieldMap  map[uintptr]Expression
	tableName Expression
}

func Alias() string {
	return ""
}
