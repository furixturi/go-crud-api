// main_test.go
package main

import (
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	const (
		APP_DB_USERNAME = "postgres"
		APP_DB_PASSWORD = "password"
		APP_DB_NAME     = "testdb"
	)
	a.Initialize(APP_DB_USERNAME, APP_DB_PASSWORD, APP_DB_NAME)

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(	
	id		SERIAL,
	name	TEXT NOT NULL,
	price	NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_key PRIMARY KEY(id)
)`