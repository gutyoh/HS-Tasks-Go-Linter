// https://stepik.org/edit-lesson/771778/step/10

// Linter log:
/*
task11\task11.go:41:1: Function name: main, Cyclomatic Complexity: 5, Halstead Volume: 541.78, Maintainability Index: 45 (maintidx)
func main() {
^
task11\task11.go:43:2: variable name 'db' is too short for the scope of its usage (varnamelen)
        db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
        ^
*/

package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

// DO NOT modify the model declarations below!
type Continent struct {
	gorm.Model
	Name      string
	Countries []Country
}

type Country struct {
	gorm.Model
	Name        string
	Capital     string
	ContinentID uint
}

type CountryStats struct {
	gorm.Model
	CountryID  uint
	Country    Country
	Population int
}

func main() {
	// DO NOT modify the gorm.Open code block, it connects to an in-memory database 'world':
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// DO NOT delete this code block! It creates the in-memory 'world' DB schema:
	err = db.AutoMigrate(&Continent{}, &Country{}, &CountryStats{})
	if err != nil {
		log.Fatal(err)
	}

	// 1. Rename the 'country_stats' table to 'country_statistics':
	err = db.Migrator().RenameTable("country_stats", "country_statistics")
	if err != nil {
		log.Fatal(err)
	}
	checkTablesExists(db) // DO NOT delete this line, it checks if the table was renamed!

	// 2. Adding a 'gdp' column to the 'country_statistics' table:
	/* When adding a column to a renamed table, we need to create a new Model
		That has the updated name of the table, and also add a field to the new Model
	    with the name of the new column we want to add: */

	// 2.1 Add the 'GDP' int type field to the new CountryStatistics Model below:
	type CountryStatistics struct {
		gorm.Model
		CountryID  uint
		Country    Country
		Population int
		GDP        int // Add the 'GDP' int type field here
	}

	// 2.2 Use AddColumn() to add the 'GDP' field to the CountryStatistics Model:
	err = db.Migrator().AddColumn(&CountryStatistics{}, "GDP")
	if err != nil {
		log.Fatal(err)
	}
	checkColumnsExists(db) // DO NOT delete this line, it checks if the column was added!
}

//::footer
const TableName = "country_statistics"

func checkTablesExists(db *gorm.DB) {
	if !db.Migrator().HasTable(TableName) {
		log.Fatal("'country_statistics' table does not exist!")
	}
	fmt.Println("'country_statistics' table successfully renamed!")
}

func checkColumnsExists(db *gorm.DB) {
	if !db.Migrator().HasColumn(TableName, "GDP") {
		log.Fatal("'gdp' column does not exist!")
	}
	fmt.Println("'gdp' column successfully added!")
}
