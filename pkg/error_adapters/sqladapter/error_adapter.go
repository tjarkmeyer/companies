package sqladapter

import (
	"errors"

	"github.com/jackc/pgconn"
	"github.com/lib/pq"
)

type IErrorAdapter interface {
	AdaptSqlErr(err error) (adapted error)
}

type ErrorAdapter struct {
	errorsMap map[string]error
}

func New(errorsMap map[string]error) *ErrorAdapter {
	return &ErrorAdapter{
		errorsMap: errorsMap,
	}
}

func (a *ErrorAdapter) AdaptSqlErr(err error) (adapted error) {
	if err == nil {
		return nil
	}

	var (
		code      string
		pqErr     *pq.Error
		pqConnErr *pgconn.PgError
	)

	if ok := errors.As(err, &pqErr); ok {
		code = string(pqErr.Code)
	}
	if ok := errors.As(err, &pqConnErr); ok {
		code = pqConnErr.Code
	}

	foundErr, found := a.errorsMap[code]
	if !found {
		return err
	}

	return foundErr
}
