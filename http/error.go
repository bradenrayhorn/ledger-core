package http

import (
	"errors"
	"net/http"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/go-chi/render"
)

type errorResponse struct {
	Error      string `json:"error"`
	StatusCode int    `json:"-"`
}

func (r *errorResponse) Render(w http.ResponseWriter, req *http.Request) error {
	render.Status(req, r.StatusCode)
	return nil
}

func (s *Server) HandleError(w http.ResponseWriter, req *http.Request, err error) {
	if err == nil {
		return
	}

	var apiError core.APIError
	var statusCode = 500
	var msg = "internal error"
	if errors.As(err, &apiError) {
		statusCode, msg = apiError.APIError()
	}

	if statusCode == 500 {
		s.Logger.Error(err.Error())
	}

	response := &errorResponse{Error: msg, StatusCode: statusCode}
	renderErr := render.Render(w, req, response)
	if renderErr != nil {
		w.WriteHeader(500)
		s.Logger.Error(renderErr.Error())
	}
}
