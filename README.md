# GoDB

A lightweight, file-driven Key-Value database abstraction built on top of SQLite in Go.

## Overview

GoDB provides a simple wrapper for managing SQLite databases as key-value stores. It isolates SQL logic into external `.sql` query files rather than hardcoding raw query strings inside Go code. Built with `modernc.org/sqlite`, it requires no CGO tooling and runs cross-platform out of the box.

## Key Features

- 🚀 **Zero-CGO Dependency**: Uses a pure Go SQLite driver (`modernc.org/sqlite`), making cross-compilation fast and effortless.
- 📁 **File-Driven SQL Logic**: SQL queries (`create_table.sql`, `set_data.sql`, `get_data.sql`) are separated into clean `.sql` files.
- 🔑 **Generic Key-Value API**: Built with Go generics (`[valType any]`) to support writing arbitrary value types into SQLite storage.
- 🔄 **Automatic Upserts**: Implements atomic insert/update semantics (`INSERT OR REPLACE`) backed by unique key indexing.
- 📂 **Multi-Database Management**: Registers and isolates separate named database files inside `./src/databases/`.

## Architecture & Directory Structure

```text
├── README.md
├── go.mod
├── go.sum
└── src/
    ├── main.go                  # Example application entrypoint
    ├── databases/               # Generated SQLite database files (.db)
    └── go_database/
        ├── framework.go         # Core database manager & connection pool
        └── sql_functions/
            ├── create_table.sql # Table initialization & index cleanup
            ├── set_data.sql     # Upsert query (INSERT OR REPLACE)
            └── get_data.sql     # Key retrieval query
```

## Quick Start Example

```go
package main

import (
	"fmt"
	"godb/src/go_database"
)

func main() {
	// 1. Register or create a SQLite database file ("app.db")
	err := go_database.RegisterDatabase("app")
	if err != nil {
		fmt.Println("Error registering database:", err)
		return
	}

	// 2. Set or update a key-value pair
	err = go_database.SetData("app", "Alice", 5590)
	if err != nil {
		fmt.Println("Error setting data:", err)
		return
	}

	// 3. Query data by key
	val, err := go_database.GetData[int]("app", "Alice")
	if err != nil {
		fmt.Println("Error getting data:", err)
		return
	}
	fmt.Println("Value:", val)
}
```

## Current Limitations & Potential Improvements

- **Runtime Disk Reads**: SQL files are currently read from disk at runtime using `os.ReadFile` on every query call. This can be optimized using Go's `embed.FS` or in-memory string caching.
- **Single-Table Schema**: Currently structured for key-value pairs stored in a `users` table.
