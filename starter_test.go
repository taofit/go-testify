package starter_test

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	starter "github.com/williaminfante/go_test_starter"
)

func TestSayHello(t *testing.T) {
	greeting := starter.SayHello("Steven")
	assert.Equal(t, "Hello Steven. Welcome!", greeting)

	greeting = starter.SayHello("Anastasia")
	assert.Equal(t, "Hello Anastasia. Welcome!", greeting)
}

func TestOddOrEven(t *testing.T) {
	t.Run("check non negative number", func(t *testing.T) {
		assert.Equal(t, "45 is an odd number", starter.OddOrEven(45))
		assert.Equal(t, "0 is an even number", starter.OddOrEven(0))
	})
	t.Run("check 2", func(t *testing.T) {
		assert.Equal(t, "-34 is an even number", starter.OddOrEven(-34))
		assert.Equal(t, "-3421 is an odd number", starter.OddOrEven(-3421))
	})
}

func TestCheckHealth(t *testing.T) {
	t.Run("Check health status", func(t *testing.T) {
		writer := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "http://someurl", nil)
		starter.CheckHealth(writer, req)
		response := writer.Result()
		body, err := io.ReadAll(response.Body)
		assert.Equal(t, "http check passed", string(body))
		assert.Equal(t, 200, response.StatusCode)
		assert.Equal(t, "text/plain; charset=utf-8", response.Header.Get("Content-type"))
		assert.Equal(t, nil, err)
	})
}
