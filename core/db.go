package core

import (
	"database/sql"
)

type field struct {
	Column  string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}

func GetAllTables(db *sql.DB) []string {
	rows, err := db.Query("show tables;")
	if err != nil {
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

func getAllFields(db *sql.DB, tableName string) ([]field, error) {
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

func GetDb(dbType string, dbDn string) *sql.DB {
	db, err := sql.Open(dbType, dbDn)
	if err != nil {
		panic(err)
	}
	return db
}
