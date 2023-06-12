package http_adapter

type IErrorAdapter interface {
	AdaptToHttpCode(err error) (adapted int)
}

type ErrorAdapter struct {
	defaultErrorCode int
	adapters         []ErrorToHttpCodeAdapter
}

func New(defaultErrorCode int, adapters ...ErrorToHttpCodeAdapter) *ErrorAdapter {
	return &ErrorAdapter{
		defaultErrorCode: defaultErrorCode,
		adapters:         adapters,
	}
}

func (a *ErrorAdapter) AdaptToHttpCode(err error) (adapted int) {
	for i := range a.adapters {
		if code := a.adapters[i](err); code > 0 {
			return code
		}
	}

	return a.defaultErrorCode
}
