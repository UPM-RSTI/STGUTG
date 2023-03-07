package http_wrapper_test

import (
	"free5gc/lib/http_wrapper"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewResponse(t *testing.T) {
	response := http_wrapper.NewResponse(http.StatusCreated, map[string][]string{
		"Location": {"https://www.nctu.edu.tw/"},
		"Refresh":  {"url=https://free5gc.org"},
	}, 1000)
	assert.Equal(t, "https://www.nctu.edu.tw/", response.Header.Get("Location"))
	assert.Equal(t, "url=https://free5gc.org", response.Header.Get("Refresh"))
	assert.Equal(t, 1000, response.Body)
}
