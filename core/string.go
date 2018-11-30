package core

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/wnote/db2struct/config"
)

func ReplaceStruct(structStr string, tablePrefix string, tableName string) (string, string) {
	parseTableName := tableName
	if tablePrefix != "" {
		pos := strings.Index(tableName, tablePrefix)
		if pos == 0 {
			parseTableName = tableName[len(tablePrefix):]
		}
	}
	structName := parseName(parseTableName)
	structStr = strings.Replace(structStr, "{struct_name}", structName, -1)
	structStr = strings.Replace(structStr, "{table_name}", tableName, -1)
	return structStr, parseTableName
}

func ReplaceFields(db *sql.DB, structStr string, packageName string, tableName string) string {
	fieldList, err := getAllFields(db, tableName)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var timeTimeExist bool
	structStr = regexp.MustCompile("\\s*\\{struct_field\\}.*").ReplaceAllStringFunc(structStr, func(s string) string {
		var replaceStructList [][]string
		var maxStructFieldLen, maxTypeLen int
		for _, field := range fieldList {

			typeName := getGoTypeFromColumnType(field.Type)

			structField := parseName(field.Column)
			if fLen := len(structField); fLen > maxStructFieldLen {
				maxStructFieldLen = fLen
			}
			if fLen := len(typeName); fLen > maxTypeLen {
				maxTypeLen = fLen
			}
			if typeName == "time.Time" {
				timeTimeExist = true
			}
			replaceStructList = append(replaceStructList, []string{structField, typeName, field.Column})
		}

		var rtnStr string
		for _, replaceStruct := range replaceStructList {
			curColumn := s
			curColumn = strings.Replace(curColumn, "{struct_field}", replaceStruct[0]+strings.Repeat(" ", maxStructFieldLen-len(replaceStruct[0])+1)+replaceStruct[1]+strings.Repeat(" ", maxTypeLen-len(replaceStruct[1])), -1)
			curColumn = strings.Replace(curColumn, "{column_name}", replaceStruct[2], -1)
			rtnStr += curColumn
		}
		return rtnStr
	})
	structStr = strings.Replace(structStr, "{package}", packageName, -1)
	if timeTimeExist {
		structStr = regexp.MustCompile("\\s*package\\s.*").ReplaceAllStringFunc(structStr, func(s string) string {
			s = s + "\n\nimport (\n	\"time\"\n)"
			return s
		})
	}
	return structStr
}

func getGoTypeFromColumnType(columnTypeName string) string {
	fieldType := regexp.MustCompile("[^a-zA-Z].*").ReplaceAllString(columnTypeName, "")
	var structTypeName string
	if regexp.MustCompile("\\sunsigned").MatchString(columnTypeName) {
		structTypeName, _ = config.DbUnsignedTypeMap[fieldType]
	}
	if structTypeName == "" {
		var exist bool
		structTypeName, exist = config.DbTypeMap[fieldType]
		if !exist {
			structTypeName = "string"
			fmt.Println(columnTypeName, " not exist in mysql type map, we will change it to string")
		}
	}

	return structTypeName
}

func parseName(name string) string {
	f := func(s string) string {
		b := []byte(s)
		lastLen := len(b) - 1
		if b[lastLen] >= 97 {
			b[lastLen] -= 32
		}
		return string(b[lastLen:])
	}
	name = regexp.MustCompile("^[a-zA-Z]").ReplaceAllStringFunc(name, f)
	name = regexp.MustCompile("_[a-zA-Z]").ReplaceAllStringFunc(name, f)
	return name
}
