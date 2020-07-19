package router

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var router *mux.Router
var req *http.Request
var err error
var respRec *httptest.ResponseRecorder

func init() {
	router = mux.NewRouter()
	APIRouter := router.PathPrefix("/api/v1").Subrouter()

	createReaderRouter(APIRouter)
	createWriterRouter(APIRouter)
}

func createTestRequest(t *testing.T, method, url string, payload []byte) *http.Request {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		t.Errorf("Creating '%s %s' request failed!", method, url)
	}

	return req
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	return recorder
}

func checkResponseCode(t *testing.T, expected int, actual int) {
	assert.Equal(
		t,
		expected,
		actual,
		fmt.Sprintf("Expected response code %d. Got %d", expected, actual),
	)
}
