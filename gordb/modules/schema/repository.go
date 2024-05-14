package schema

func GetMeta() []Meta {
	content, err := commons.ReadFile("meta.json")
	if err != nil {
		return []Meta{}
	}

	var meta []Meta

	if err := json.Unmarshal(content, &meta); err != nil {
		return []Meta{}
	}

	return meta
}

func UpdateMeta(updatedMeta []Meta) error {
	return commons.WriteFile("meta.json", updatedMeta)
}

func MetaContains(meta []Meta, tableName string) bool {
	for _, table := range meta {
		if table.TableName == tableName {
			return true
		}
	}

	return false
}


func GetTables() []Table {
	return []Table{
		{
			Name: "table1",
			Columns: []Column{
				{
					Name: "column1",
					Type: "string",
				},
				{
					Name: "column2",
					Type: "int",
				},
			},
		},
		{
			Name: "table2",
			Columns: []Column{
				{
					Name: "column1",
					Type: "string",
				},
				{
					Name: "column2",
					Type: "int",
				},
			},
		},
	}
}
