package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/сafe?count=10&city=moscow", nil) // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code %d, but got %d", http.StatusOK, responseRecorder.Code)
	expectedBody := "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент"
	assert.Equal(t, expectedBody, responseRecorder.Body.String(), "Expected body %d, but got %d",
		expectedBody, responseRecorder.Body.String())
}

func TestMainHandlerValidRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/cafe?count=3&city=moscow", nil)
	require.NoError(t, err, "The request is not formed correctly")

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code, "Expected status code %d, but got %d", http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body.String(), "The response body is empty")
}
