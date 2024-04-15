package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestApiTimeoff(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "api/timeoff", nil)

	r := gin.Default()
	r.ServeHTTP(rec, req)
	res := rec.Result()

	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
	t.Log(rec.Body.String())
}
