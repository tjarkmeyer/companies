package httpadapter

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjarkmeyer/companies/companies/internal/v1/repositories"
)

var (
	expected = map[error]int{
		repositories.ErrNotFound:     http.StatusNotFound,
		repositories.ErrAlreadyExist: http.StatusConflict,
	}
)

func Test_Adapter(t *testing.T) {
	defaultCode := http.StatusInternalServerError
	adapter := New(defaultCode, AdaptBadRequestError)

	assert.Equal(t, expected[repositories.ErrNotFound], adapter.AdaptToHttpCode(repositories.ErrNotFound))
	assert.Equal(t, expected[repositories.ErrAlreadyExist], adapter.AdaptToHttpCode(repositories.ErrAlreadyExist))
	assert.Equal(t, defaultCode, adapter.AdaptToHttpCode(errors.New("")))
}
