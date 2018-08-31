package api

import (
	"fmt"
	"os"
	"testing"

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
	wd, _ := os.Getwd()
	fmt.Println(wd)
	test.CodeTest(t, "GET", "/v1/users/fake.tom.j/messages/100", 404)
	test.CodeTest(t, "GET", "/v1/users/tom.j1/messages/11000", 404)
	test.CodeTest(t, "GET", "/v1/users/tom.j1/messages/-1", 404)
	test.CodeTest(t, "GET", "/v1/users/tom.j1/messages/abc", 400)
	test.BodyResponseTest(t, "GET", "/v1/users/tom.j/messages/100", nil, 200, "message")
}
