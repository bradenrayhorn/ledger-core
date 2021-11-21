package http

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

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

func (s *Server) UpdateMarketProvider(w http.ResponseWriter, req *http.Request) {

}
