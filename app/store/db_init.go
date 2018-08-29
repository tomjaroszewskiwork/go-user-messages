package store

import (
	"github.com/jinzhu/gorm"
	// Imports sql lite so the db starts
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// InitDB connects to the database and initializes the store
func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&UserMessage{})
}
