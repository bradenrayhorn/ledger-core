package core

var (
	ErrorAuth = &apiError{status: 401, msg: "invalid authorization"}
)

type APIError interface {
	APIError() (int, string)
}

type apiError struct {
	status int
	msg    string
}

func (e apiError) Error() string {
	return e.msg
}

func (e apiError) APIError() (int, string) {
	return e.status, e.msg
}

type wrappedAPIError struct {
	error
	apiError *apiError
}

func (e wrappedAPIError) Is(err error) bool {
	return e.apiError == err
}

func (e wrappedAPIError) Unwrap() error {
	return e.error
}

func (e wrappedAPIError) APIError() (int, string) {
	return e.apiError.APIError()
}

func WrapError(err error, apiError *apiError) error {
	return wrappedAPIError{error: err, apiError: apiError}
}
