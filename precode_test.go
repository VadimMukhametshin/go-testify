package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func NewRec(req *http.Request, handler http.HandlerFunc) *httptest.ResponseRecorder {
	responseRecorder := httptest.NewRecorder()
	handler = http.HandlerFunc(handler)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder
}

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=2&city=moscow", nil)

	responseRecorder := NewRec(req, mainHandle)

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=wocsom", nil)

	responseRecorder := NewRec(req, mainHandle)

	expected := `wrong city value`

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, expected, responseRecorder.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=10&city=moscow", nil)

	responseRecorder := NewRec(req, mainHandle)

	resBody := strings.Split(responseRecorder.Body.String(), ",")

	assert.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Len(t, resBody, totalCount)
}
