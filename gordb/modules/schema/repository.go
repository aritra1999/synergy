package schema

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
