// https://stepik.org/edit-lesson/771778/step/10

// Linter log:
/*
task10\task10.go:35:2: variable name 'db' is too short for the scope of its usage (varnamelen)
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

// Create the 'Track' model below:
type Track struct {
	gorm.Model
	Name     string
	Duration int
	AlbumID  uint
}

// Create the 'Album' model below:
type Album struct {
	gorm.Model
	Title    string
	Tracks   []Track
	ArtistID uint
}

// Create the 'Artist' model below:
type Artist struct {
	gorm.Model
	Name   string
	Albums []Album
}

func main() {
	// DO NOT modify the gorm.Open code block, it connects to an in-memory database to create the 'songs' schema:
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Write the code to migrate the schema conformed by 'Track', 'Album' and 'Artist' models:
	err = db.AutoMigrate(&Track{}, &Album{}, &Artist{})
	if err != nil {
		log.Fatal(err)
	}

	// DO NOT delete the code line below, it checks if the tables exist after the migration:
	checkTablesExists(db)
}

//::footer
func checkTablesExists(db *gorm.DB) {
	if !db.Migrator().HasTable("tracks") {
		log.Fatal("Table 'tracks' not found")
	}

	if !db.Migrator().HasTable("albums") {
		log.Fatal("Table 'albums' not found")
	}

	if !db.Migrator().HasTable("artists") {
		log.Fatal("Table 'artists' not found")
	}

	fmt.Println("All models successfully migrated!")
}
