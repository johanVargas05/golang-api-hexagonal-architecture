package health_controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	controller := New()

	assert.NotNil(t, controller)
}

func TestExecute(t *testing.T) {
	controller := New()

	t.Run("Success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/health", http.NoBody)
		rec := httptest.NewRecorder()
		ctx := e.NewContext(req, rec)

		expectedResponse := struct {
			Code int                          `json:"code"`
			Data       map[string]any `json:"data"`
			Message string `json:"message"`
		}{}
		
		response := controller.Execute(ctx)
		err := json.Unmarshal([]byte(rec.Body.Bytes()), &expectedResponse)

		assert.NoError(t, err)
		assert.Nil(t, response)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 200, expectedResponse.Code)
		assert.Equal(t, "API is healthy", expectedResponse.Message)
		assert.Equal(t, "UP", expectedResponse.Data["status"])
		assert.NotNil(t, expectedResponse.Data["current_time"])
	})
}