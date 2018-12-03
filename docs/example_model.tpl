package {package}

type {struct_name} struct {
	{struct_field} `json:"{column_name}" db:"{column_name}`
}

func (p {struct_name}) TableName() string {
	return "{table_name}"
}