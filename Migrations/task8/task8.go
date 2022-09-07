// https://stepik.org/edit-lesson/771778/step/8

// Linter log:
/*
task8.go:48:2: variable name 'db' is too short for the scope of its usage (varnamelen)
        db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
        ^
task8.go:64:24: parameter name 'db' is too short for the scope of its usage (varnamelen)
func checkTablesExists(db *gorm.DB) {
*/

package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

// DO NOT modify the 'tournament' models below:
type Team struct { // Creates the "teams" table
	gorm.Model
	Name string
}

type Player struct { // Creates the "players" table
	gorm.Model
	FirstName string
	LastName  string
	Age       int
	TeamID    uint
}

type Game struct { // Creates the "games" table
	gorm.Model
	TeamID        uint
	HomeTeam      Team
	HomeTeamID    uint
	HomeTeamScore uint
	AwayTeam      Team
	AwayTeamID    uint
	AwayTeamScore uint
	Date          time.Time
}

func main() {
	// DO NOT modify the gorm.Open code block, it connects
	// to an in-memory database to create the 'tournament' schema:
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Write the code to run an Auto Migration
	// for the 'Team', 'Player' and 'Game' models below:
	err = db.AutoMigrate(&Team{}, &Player{}, &Game{})
	if err != nil {
		log.Fatal(err)
	}

	// DO NOT delete the code line below, it checks if the tables exist after the migration:
	checkTablesExists(db)
}

func checkTablesExists(db *gorm.DB) {
	if !db.Migrator().HasTable("teams") {
		log.Fatal("Table 'teams' not found")
	}
	fmt.Println("Table 'teams' successfully migrated!")

	if !db.Migrator().HasTable("players") {
		log.Fatal("Table 'players' not found")
	}
	fmt.Println("Table 'players' successfully migrated!")

	if !db.Migrator().HasTable("games") {
		log.Fatal("Table 'games' not found")
	}
	fmt.Println("Table 'games' successfully migrated!")
}
