package http_adapter

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjarkmeyer/companies/companies/internal/v1/repositories"
)

var (
	expectedMap = map[error]int{
		repositories.ErrNotFound:     http.StatusNotFound,
		repositories.ErrAlreadyExist: http.StatusBadRequest,
	}
)

func Test_Adapter(t *testing.T) {
	defaultCode := http.StatusInternalServerError
	adapter := New(defaultCode, AdaptNotFoundError, AdaptBadRequestError)

	assert.Equal(t, expectedMap[repositories.ErrNotFound], adapter.AdaptToHttpCode(repositories.ErrNotFound))
	assert.Equal(t, expectedMap[repositories.ErrAlreadyExist], adapter.AdaptToHttpCode(repositories.ErrAlreadyExist))
	assert.Equal(t, defaultCode, adapter.AdaptToHttpCode(errors.New("")))
}
