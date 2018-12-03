package config

// Config the model template for go file.
// {package},{struct_name},{struct_field} is required for template string
const MODEL_STRUCT_TEMPLATE = `package {package}

type {struct_name} struct {
	{struct_field} ` + "`" + `json:"{column_name}" db:"{column_name}"` + "`" + `
}

func (p {struct_name}) TableName() string {
	return "{table_name}"
}
`
