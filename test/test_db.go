package test

import (
	"fmt"
	"os"

	"github.com/tomjaroszewskiwork/go-user-messages/app/store"
)

// StartTestDB starts the DB for testing and inserts test data
func StartTestDB() {
	store.InitDB()

	testInsert := `INSERT INTO user_messages (user_id, message_id, message, generated_at) VALUES
	('tom.j', 100, 'test string', '2018-08-08T20:08:08'),
	('tom.j', 101, 'lool', '2018-08-08T20:08:09'),
	('tom.j', 102, 'longer message filled with words', '2018-10-09T20:08:09'),
 	('tom.j', 103, 'even more words', '2018-10-10T20:08:09'),
	('tom.j', 104, 'longgest message possible', '2018-10-11T20:08:09'),
 	('tom.j', 105, 'ldsfsdfsdf test test este stest', '2018-10-11T21:08:09'),
 	('fun.dude', 150, 'sad message :-(', '2018-10-09T20:08:09'),
 	('fun.dude', 151, 'WOW!', '2018-10-09T20:08:09'),
 	('fun.dude', 152, 'a a1100000000011a a', '2018-10-09T20:08:09'),
	('bob.dole', 200, 'secret words secret', '2018-12-08T20:08:09')`

	err := store.DB.Exec(testInsert).Error
	if err != nil {
		fmt.Println("Test data isnert failed " + err.Error())
		os.Exit(1)
	}
}
