package config

var (
	DbTypeMap = map[string]string{
		"tinyint":    "int8",
		"smallint":   "int16",
		"mediumint":  "int32",
		"int":        "int32",
		"integer":    "int32",
		"bigint":     "int64",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
		"date":       "string",
		"datetime":   "time.Time",
		"timestamp":  "time.Time",
		"time":       "time.Time",
		"year":       "uint8",
		"varbinary":  "string",
		"enum":       "string",
		"set":        "string",
		"char":       "string",
		"varchar":    "string",
		"binary":     "string",
		"tinytext":   "string",
		"text":       "string",
		"mediumtext": "string",
		"longtext":   "string",
		"tinyblob":   "[]byte",
		"mediumblob": "[]byte",
		"blob":       "[]byte",
		"longblob":   "[]byte",
	}
	DbUnsignedTypeMap = map[string]string{
		"tinyint":   "uint8",
		"smallint":  "uint16",
		"mediumint": "uint32",
		"int":       "uint32",
		"integer":   "uint32",
		"bigint":    "uint64",
	}
	DbType    = "mysql"
	FileModel = "0777"
)
