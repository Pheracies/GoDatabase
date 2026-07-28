package go_database

import (
	"database/sql"
	_ "embed"
	"fmt"
	"os"

	"github.com/Heaplyn/GoDatabase/src/json_conversion"
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

type DatabaseOptions[T any] struct {
	DefaultValue T
}

type Database[T any] struct {
	Options *DatabaseOptions[T]
	db      *sql.DB
}

var Properties PropertiesType = PropertiesType{} // Public properties
var databases map[string]*Database[any] = make(map[string]*Database[any])

func setupDefaultOptions[T any](options *DatabaseOptions[T]) *DatabaseOptions[T] {
	if options == nil {
		options = &DatabaseOptions[T]{}
	}
	return options
}

func RegisterDatabase[T any](database string, options *DatabaseOptions[T]) error {
	options = setupDefaultOptions(options)

	dir := "./src/databases"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	if existing, ok := databases[database]; ok && existing != nil && existing.db != nil {
		existing.db.Close()
	}

	path := dir + "/" + database + ".db"

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	_, _ = db.Exec("PRAGMA journal_mode=WAL;")
	_, _ = db.Exec("PRAGMA busy_timeout=5000;")

	currentDatabase := &Database[any]{
		Options: &DatabaseOptions[any]{
			DefaultValue: options.DefaultValue,
		},
		db: db,
	}

	_, err = currentDatabase.db.Exec(createTableSQL)
	if err != nil {
		db.Close()
		return fmt.Errorf("error creating users table: %w", err)
	}

	databases[database] = currentDatabase
	return nil
}

func SetData[valType any](database string, key string, value valType) error {
	db, exists := databases[database]
	if !exists {
		return fmt.Errorf("database '%s' is not registered", database)
	}
	options := db.Options
	if any(value) == nil && options != nil {
		if def, ok := options.DefaultValue.(valType); ok {
			value = def
		}
	}
	fmt.Println("Key: ", key)
	fmt.Println("Name: ", value)
	_, err := db.db.Exec(setDataSQL, key, value)
	if err != nil {
		return fmt.Errorf("error executing set_data: %w", err)
	}

	return nil
}

func GetData[T any](database string, key string) (T, error) {
	fmt.Println("Getting data for ", key, " in ", database)
	var zero T
	db, exists := databases[database]
	if !exists {
		return zero, fmt.Errorf("database '%s' is not registered", database)
	}
	if db.Options != nil {
		if def, ok := db.Options.DefaultValue.(T); ok {
			zero = def
		}
	}

	data := db.db.QueryRow(getDataSQL, key)

	var key2 string
	var value2 T
	if err := data.Scan(&key2, &value2); err != nil {
		return zero, fmt.Errorf("error scanning data: %w", err)
	}

	return value2, nil
}

func Encrypt(data any) (string, error) {
	return json_conversion.Encrypt(data)
}

func Decrypt[T any](data string) (T, error) {
	return json_conversion.Decrypt[T](data)
}
