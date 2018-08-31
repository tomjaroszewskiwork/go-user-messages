package api

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	store "github.com/tomjaroszewskiwork/go-user-messages/app/store"
	test "github.com/tomjaroszewskiwork/go-user-messages/test"
)

func TestMain(m *testing.M) {
	test.StartTestDB()
	test.SetTestRouter(NewRouter())
	retCode := m.Run()
	store.CleanDB()
	os.Exit(retCode)
}

func TestGetMessage(t *testing.T) {
	test.CodeTest(t, "GET", "/v1/users/fake.tom.j/messages/100", nil, 404)
	test.CodeTest(t, "GET", "/v1/users/tom.j1/messages/11000", nil, 404)
	test.CodeTest(t, "GET", "/v1/users/tom.j1/messages/-1", nil, 404)
	test.CodeTest(t, "GET", "/v1/users/tom.j1/messages/abc", nil, 400)
	test.BodyResponseTest(t, "GET", "/v1/users/tom.j/messages/100", nil, 200, "message")
}

func TestAddMessage(t *testing.T) {
	test.CodeTest(t, "POST", "/v1/users/tom.j/messages", nil, 400)
	messageBody := MessageBody{Message: ""}
	test.CodeTest(t, "POST", "/v1/users/tom.j/messages", messageBody, 400)
	messageBody.Message = "values"
	response := test.CodeTest(t, "POST", "/v1/users/tom.j/messages", messageBody, 201)
	newLocation := response.HeaderMap.Get("Location")
	response = test.CodeTest(t, "GET", newLocation, nil, 200)
	// Can't do auto JSON compare as time will differ, freezing time is too complicated
	var newMessage store.UserMessage
	err := json.NewDecoder(response.Body).Decode(&newMessage)
	if err != nil {
		require.FailNow(t, err.Error())
	}
	require.Equal(t, "tom.j", newMessage.UserID)
	require.Equal(t, messageBody.Message, newMessage.Message)
}
