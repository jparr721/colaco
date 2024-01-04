package db

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// MockColacoDB implements the ColacoDBInterface for testing purposes.
type MockColacoDB struct {
	Sodas  map[string]any
	Promos map[string]any
}

// NewMockColacoDB creates a new instance of MockColacoDB.
func NewMockColacoDB() *MockColacoDB {
	return &MockColacoDB{
		Sodas:  make(map[string]any),
		Promos: make(map[string]any),
	}
}

func (mdb *MockColacoDB) Init() error {
	// Initialization logic for the mock database (if needed).
	return nil
}

func (mdb *MockColacoDB) Get(query string, dest any, args ...interface{}) error {
	destVal := reflect.ValueOf(dest).Elem()
	if destVal.Kind() != reflect.Slice {
		return errors.New("destination is not a slice")
	}

	// Determine which table we are querying
	if strings.Contains(query, "FROM promos") {
		for _, record := range mdb.Promos {
			recordVal := reflect.ValueOf(record)
			if destVal.Type().Elem() != recordVal.Type() {
				continue // Skip if the types do not match
			}
			destVal.Set(reflect.Append(destVal, recordVal))
		}
	} else if strings.Contains(query, "FROM sodas") {
		for _, record := range mdb.Sodas {
			recordVal := reflect.ValueOf(record)
			if destVal.Type().Elem() != recordVal.Type() {
				continue // Skip if the types do not match
			}
			destVal.Set(reflect.Append(destVal, recordVal))
		}
	} else {
		return errors.New("unknown table in query")
	}

	return nil
}

func (mdb *MockColacoDB) GetOne(query string, dest any, args ...interface{}) error {
	// Simulate fetching a record by ID.
	// In a real scenario, you would parse the SQL query.
	// For simplicity, assume we're always querying by ID.

	// Check if the query is for promos or sodas
	var record any
	var found bool
	if strings.Contains(query, "FROM promos") {
		// Fetch a promo
		id, _ := args[0].(string)
		record, found = mdb.Promos[id]
	} else if strings.Contains(query, "FROM sodas") {
		// Fetch a soda
		id, _ := args[0].(string)
		record, found = mdb.Sodas[id]
	} else {
		return errors.New("unknown table in query")
	}

	if !found {
		return sql.ErrNoRows // Mimic the behavior of the sql package when no rows are found
	}

	// Assuming dest is a pointer to a struct, and record is also a struct
	destVal := reflect.ValueOf(dest).Elem()
	recordVal := reflect.ValueOf(record)

	// Copy record to dest
	if destVal.Type() != recordVal.Type() {
		return errors.New("type mismatch between dest and record")
	}
	destVal.Set(recordVal)

	return nil
}

func (mdb *MockColacoDB) Create(query string, args ...interface{}) (string, error) {
	if strings.Contains(query, "INTO promos") {
		// Assume args is in the same order as the struct fields and just parse directly into
		// the db
		id := "10"
		fmt.Println(args)
		mdb.Promos[id] = args[0]
		return id, nil
	} else if strings.Contains(query, "INTO sodas") {
		id := "10"
		mdb.Sodas[id] = args[0]
		return id, nil
	} else {
		return "", errors.New("unknown table in query")
	}
}

func (mdb *MockColacoDB) Update(query string, args ...interface{}) error {
	if strings.Contains(query, "INTO promos") {
		// Assume args is in the same order as the struct fields and just parse directly into
		// the db
		id := args[0].(string)
		mdb.Promos[id] = args[1]
		return nil
	} else if strings.Contains(query, "INTO sodas") {
		id := args[0].(string)
		mdb.Sodas[id] = args[1]
		return nil
	} else {
		return errors.New("unknown table in query")
	}
}
