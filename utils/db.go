package utils

import (
	"database/sql"
	"fmt"

	"github.com/wnote/db2struct/config"
)

type field struct {
	Column  string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open(config.CustomCfg["db_type"], config.CustomCfg["db_dn"])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func GetAllTables() []string {
	rows, err := db.Query("show tables;")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	var allTables []string
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		allTables = append(allTables, tableName)
	}
	return allTables
}

func getAllFields(tableName string) ([]field, error) {
	rows, err := db.Query("show fields from " + tableName + ";")
	if err != nil {
		return nil, err
	}
	var fieldsMap []field
	for rows.Next() {
		var column, fType, fNull, fKey, fDefault, fExtra string
		rows.Scan(&column, &fType, &fNull, &fKey, &fDefault, &fExtra)
		fieldsMap = append(fieldsMap, field{
			Type:    fType,
			Null:    fNull,
			Key:     fKey,
			Default: fDefault,
			Extra:   fExtra,
			Column:  column,
		})
	}
	return fieldsMap, nil
}
