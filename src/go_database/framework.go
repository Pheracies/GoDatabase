package go_database

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

type PropertiesType struct {
}

var Properties PropertiesType = PropertiesType{} // Public properties
var databases map[string]*sql.DB = make(map[string]*sql.DB)

func RegisterDatabase(database string) error {
	if databases == nil {
		databases = make(map[string]*sql.DB)
	}

	dir := "./src/databases"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	path := dir + "/" + database + ".db"
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	createTableSQL, err := os.ReadFile("./src/go_database/sql_functions/create_table.sql")
	if err != nil {
		return fmt.Errorf("error reading create_table.sql: %w", err)
	}

	_, err = db.Exec(string(createTableSQL))
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

	setDataSQL, err := os.ReadFile("./src/go_database/sql_functions/set_data.sql")
	if err != nil {
		return fmt.Errorf("error reading set_data.sql: %w", err)
	}

	_, err = db.Exec(string(setDataSQL), key, value)
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

	setDataSQL, err := os.ReadFile("./src/go_database/sql_functions/get_data.sql")
	if err != nil {
		return fmt.Errorf("error reading get_data.sql: %w", err)
	}

	data, err2 := db.Query(string(setDataSQL), key)
	if err2 != nil {
		return fmt.Errorf("error executing get_data: %w", err2)
	}

	for data.Next() {
		var key, value string
		if err := data.Scan(&key, &value); err != nil {
			return fmt.Errorf("error scanning data: %w", err)
		}
		fmt.Println("Key:", key, "Value:", value)
	}

	return nil
}
