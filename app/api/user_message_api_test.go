package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	store "github.com/tomjaroszewskiwork/go-user-messages/app/store"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

var testRouter *mux.Router

const testFixterDir = "../../test/data/"

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
	('tom.j', 100, 'test string', '2018-08-08T20:08:08'),
	('tom.j', 101, 'lool', '2018-08-08T20:08:09:00.000'),
	('tom.j', 102, 'longer message filled with words', '2018-10-09T20:08:09:00.000'),
 	('tom.j', 103, 'even more words', '2018-10-10T20:08:09:00.000'),
	('tom.j', 104, 'longgest message possible', '2018-10-11T20:08:09:00.000'),
 	('tom.j', 105, 'ldsfsdfsdf test test este stest', '2018-10-11T21:08:09:00.000'),
 	('fun.dude', 150, 'sad message :-(', '2018-10-09T20:08:09:00.000'),
 	('fun.dude', 151, 'WOW!', '2018-10-09T20:08:09:00.000'),
 	('fun.dude', 152, 'a a1100000000011a a', '2018-10-09T20:08:09:00.000'),
	('bob.dole', 200, 'secret words secret', '2018-12-08T20:08:09:00.000')`

	err := store.DB.Exec(testInsert).Error
	if err != nil {
		fmt.Println("Test data isnert failed " + err.Error())
		os.Exit(1)
	}
}

func TestGetMessage(t *testing.T) {
	codeTest(t, "GET", "/v1/users/fake.tom.j/messages/100", 404)
	codeTest(t, "GET", "/v1/users/tom.j1/messages/11000", 404)
	codeTest(t, "GET", "/v1/users/tom.j1/messages/-1", 404)
	codeTest(t, "GET", "/v1/users/tom.j1/messages/abc", 400)
	bodyResponseTest(t, "GET", "/v1/users/tom.j/messages/100", nil, 200, "message")
}

// Just makes sure that the return code matches
func codeTest(t *testing.T, method string, url string, expectedCode int) {
	responseTest(t, method, url, nil, expectedCode)
}

// Calls the API and compares the full body return
func bodyResponseTest(t *testing.T, method string, url string, body interface{}, exectedCode int, expectedResponseFile string) {
	response := responseTest(t, method, url, body, exectedCode)
	// Does a JSON diff on the body compared to the expected JSON file
	differ := diff.New()
	testFile := testFixterDir + expectedResponseFile + ".json"
	expectedString, err := ioutil.ReadFile(testFile)
	if err != nil {
		require.FailNow(t, testFile+" "+err.Error())
	}
	delta, err := differ.Compare(expectedString, []byte(response.Body.String()))
	if err != nil {
		require.FailNow(t, err.Error())
	}
	if delta.Modified() {
		formatter := formatter.NewDeltaFormatter()
		diffString, _ := formatter.Format(delta)
		require.FailNow(t, diffString)
	}
}

// Calls the API, tests for the code and returns the response
func responseTest(t *testing.T, method string, url string, body interface{}, expectedCode int) *httptest.ResponseRecorder {
	var bodyJSON []byte
	var err error
	if body != nil {
		bodyJSON, err = json.Marshal(body)
		require.NotNil(t, err)
	}
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bodyJSON))
	response := httptest.NewRecorder()
	testRouter.ServeHTTP(response, req)
	require.Equal(t, expectedCode, response.Code)
	return response
}
