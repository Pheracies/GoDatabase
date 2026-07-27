package go_database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

//go:embed sql_functions/create_table.sql
var createTableSQL string

//go:embed sql_functions/set_data.sql
var setDataSQL string

//go:embed sql_functions/get_data.sql
var getDataSQL string

type PropertiesType struct {
}

var Properties PropertiesType = PropertiesType{} // Public properties
var databases map[string]*sql.DB = make(map[string]*sql.DB)

func RegisterDatabase(database string) error {
	if databases == nil {
		databases = make(map[string]*sql.DB)
	}

	dir := "../databases"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	path := dir + "/" + database + ".db"
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	_, err = db.Exec(createTableSQL)
	if err != nil {
		db.Close()
		return fmt.Errorf("error creating users table: %w", err)
	}

	databases[database] = db
	return nil
}

func SetData[valType any](database string, key string, value valType) error {
	db, exists := databases[database]
	if !exists {
		return fmt.Errorf("database '%s' is not registered", database)
	}

	_, err := db.Exec(setDataSQL, key, value)
	if err != nil {
		return fmt.Errorf("error executing set_data: %w", err)
	}

	return nil
}

func GetData(database string, key string) error {
	db, exists := databases[database]
	if !exists {
		return fmt.Errorf("database '%s' is not registered", database)
	}

	data, err2 := db.Query(getDataSQL, key)
	if err2 != nil {
		return fmt.Errorf("error executing get_data: %w", err2)
	}
	defer data.Close()

	for data.Next() {
		var key, value string
		if err := data.Scan(&key, &value); err != nil {
			return fmt.Errorf("error scanning data: %w", err)
		}
		fmt.Println("Key:", key, "Value:", value)
	}

	return nil
}
