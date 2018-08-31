package store

import (
	"github.com/jinzhu/gorm"
	// Imports sql lite so the db starts
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB Global database pointer
var DB *gorm.DB

// InitDB connects to the database and initializes the store
func InitDB() {
	var err error
	DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	CleanDB()
	// Migrate the schema
	DB.AutoMigrate(&UserMessage{})
}

// CleanDB resets the table
func CleanDB() {
	DB.DropTableIfExists(&UserMessage{})
}
