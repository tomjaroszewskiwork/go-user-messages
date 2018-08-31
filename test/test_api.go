package test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

var testRouter *mux.Router

// TOOD path is harded, will fail outside of app/api folder
const testFixterDir = "../../test/testdata/"

// SetTestRouter sets the test router
func SetTestRouter(newRouter *mux.Router) {
	testRouter = newRouter
}

// CodeTest makes sure that the return code matches
func CodeTest(t *testing.T, method string, url string, expectedCode int) {
	responseTest(t, method, url, nil, expectedCode)
}

// BodyResponseTest calls the API and compares the full body return
func BodyResponseTest(t *testing.T, method string, url string, body interface{}, exectedCode int, expectedResponseFile string) {
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
