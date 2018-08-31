package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
	store "github.com/tomjaroszewskiwork/go-user-messages/app/store"
)

var testRouter *mux.Router

func TestMain(m *testing.M) {
	testRouter = NewRouter()
	store.InitDB()
	addTestData()
	retCode := m.Run()
	store.CleanDB()
	os.Exit(retCode)
}

func addTestData() {
	testInsert := `INSERT INTO user_messages (user_id, message_id, message, generated_at) VALUES
	('tom.j', 100, 'test string', '2018-08-08 20:08:08-0:00'),
	('tom.j', 101, 'lool', '2018-08-08 20:08:09-0:00'),
	('tom.j', 102, 'longer message filled with words', '2018-10-09 20:08:09-0:00'),
 	('tom.j', 103, 'even more words', '2018-10-10 20:08:09-0:00'),
	('tom.j', 104, 'longgest message possible', '2018-10-11 20:08:09-0:00'),
 	('tom.j', 105, 'ldsfsdfsdf test test este stest', '2018-10-11 21:08:09-0:00'),
 	('fun.dude', 150, 'sad message :-(', '2018-10-09 20:08:09-0:00'),
 	('fun.dude', 151, 'WOW!', '2018-10-09 20:08:09-0:00'),
 	('fun.dude', 152, 'a a1100000000011a a', '2018-10-09 20:08:09-0:00'),
	('bob.dole', 200, 'secret words secret', '2018-12-08 20:08:09-0:00')`

	err := store.DB.Exec(testInsert).Error
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func TestGetMessage(t *testing.T) {
	codeTest(t, "GET", "/v1/users/tom.j1/messages/100", 404)
	codeTest(t, "GET", "/v1/users/tom.j1/messages/11000", 404)
	codeTest(t, "GET", "/v1/users/tom.j/messages/100", 200)
}

func codeTest(t *testing.T, method string, url string, expectedCode int) {
	req, _ := http.NewRequest(method, url, nil)
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, req)
	assert.Equal(t, expectedCode, response.Code)
}
