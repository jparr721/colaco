package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type ColacoDBInterface interface {
	Get(query string, dest any, args ...interface{}) error
	GetOne(query string, dest any, args ...interface{}) error
	Create(query string, args ...interface{}) error
	Update(query string, args ...interface{}) error
	Delete(query string, args ...interface{}) error
}

// ColacoDB is a wrapper around sql.DB that provides additional methods and also
// makes mocking in the testing functions easier.
type ColacoDB struct {
	*sql.DB
}

// Init initializes a database connection in a way that allows us to mock for unit tests.
func (db *ColacoDB) Init() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Build connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	zap.L().Info("Successfully connected to database!")
	db.DB = conn
}

// Get is a wrapper around sql.DB.Query that scans the results into a slice of structs.
func (db *ColacoDB) Get(query string, dest any, args ...interface{}) error {
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}

	defer rows.Close()

	// Get the slice's type and create a new slice to store results
	sliceVal := reflect.ValueOf(dest).Elem()
	structType := sliceVal.Type().Elem()

	for rows.Next() {
		// Create a new struct for each sub value
		structVal := reflect.New(structType).Elem()

		// Get the struct's fields and scan the row's values into them
		valCount := structVal.NumField()
		vals := make([]interface{}, valCount)
		for i := 0; i < valCount; i++ {
			vals[i] = structVal.Field(i).Addr().Interface()
		}

		// Now scan the entries back into the type
		if err := rows.Scan(vals...); err != nil {
			return err
		}

		// Now append
		sliceVal.Set(reflect.Append(sliceVal, structVal))
	}

	return rows.Err()
}

func (db *ColacoDB) GetOne(query string, dest any, args ...interface{}) error {
	row := db.QueryRow(query, args...)

	// Get the struct's fields and scan the row's values into them
	valCount := reflect.ValueOf(dest).Elem().NumField()
	vals := make([]interface{}, valCount)
	for i := 0; i < valCount; i++ {
		vals[i] = reflect.ValueOf(dest).Elem().Field(i).Addr().Interface()
	}

	// Now scan the row into the type
	if err := row.Scan(vals...); err != nil {
		return err
	}

	return nil
}

// Create is a wrapper around sql.DB.Exec that creates a new row in the database and returns the ID of the new record.
func (db *ColacoDB) Create(query string, args ...interface{}) (string, error) {
	// Add the RETURNING clause to the query
	query = fmt.Sprintf("%s RETURNING id", query)

	var id uuid.UUID
	err := db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

// Update is a wrapper around sql.DB.Exec that updates a row in the database.
func (db *ColacoDB) Update(query string, args ...interface{}) error {
	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

// Delete is a wrapper around sql.DB.Exec that deletes a row in the database.
func (db *ColacoDB) Delete(query string, dest any, args ...interface{}) error {
	return nil
}
