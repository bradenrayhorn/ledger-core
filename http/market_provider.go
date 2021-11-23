package http

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

// get market provider

type getMarketProviderResponse struct {
	Data *string `json:"data"`
}

func (g *getMarketProviderResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *Server) GetMarketProvider(w http.ResponseWriter, req *http.Request) {
	userID := getUser(req)

	userMarketProvider, err := s.UserMarketProviderRepo.GetUserMarketProvider(req.Context(), userID)
	if err != nil {
		log.Println(err)
		return
	}

	response := &getMarketProviderResponse{Data: nil}
	if userMarketProvider != nil {
		response.Data = &userMarketProvider.Provider
	}

	render.Render(w, req, response)
}

// update market provider

type updateMarketProviderRequest struct {
	Provider string `json:"provider"`
}

func (r *updateMarketProviderRequest) Bind(req *http.Request) error {
	r.Provider = strings.TrimSpace(r.Provider)
	return nil
}

func (s *Server) UpdateMarketProvider(w http.ResponseWriter, req *http.Request) {
	data := &updateMarketProviderRequest{}
	if err := render.Bind(req, data); err != nil {
		log.Println(err)
		return
	}

	userID := getUser(req)

	var err error
	if len(data.Provider) > 0 {
		err = s.UserMarketProviderRepo.SetUserMarketProvider(req.Context(), userID, data.Provider)
	} else {
		err = s.UserMarketProviderRepo.DeleteUserMarketProvider(req.Context(), userID)
	}

	if err != nil {
		log.Println(err)
	}
}
