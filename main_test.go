package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

var router *httprouter.Router

func TestAuth(t *testing.T) {
	initDB()
	req, _ := http.NewRequest("GET", "/api/getuser/test1@gmail.com", nil)
	req.SetBasicAuth("abc", "123")

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	t.Log(responseRecorder.Body.String())

	res := response{}
	json.Unmarshal([]byte(responseRecorder.Body.String()), &res)
	//t.Log(res.Statuscode)
	assert.Equal(t, "200", res.Statuscode)

}

func TestAuth1(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/getuser/test1@gmail.com", nil)
	req.SetBasicAuth("abc", "12345")

	responseRecorder := httptest.NewRecorder()

	router.ServeHTTP(responseRecorder, req)

	//t.Log(responseRecorder.Body.String())
	//t.Log(responseRecorder.HeaderMap)

	assert.Equal(t, 401, responseRecorder.Code)
}

//TestMain sets up the testing environment
func TestMain(m *testing.M) {
	router = setupRoutes()
	code := m.Run()
	os.Exit(code)
}
