package merchantservices

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PasinduYeshan/go-sample-oauth/internal/common/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllServices(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := GetAllServices(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var got response.APIResponse
	if assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &got)) {
		assert.Equal(t, "Ads retrieved successfully.", got.Message)
		assert.Equal(t, "success", got.Status)
	}
}
