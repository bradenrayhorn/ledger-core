package http

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) authentication() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			if err != nil {
				failedAuthentication(w, err)
				return
			}

			userID, err := s.SessionService.GetUserFromSession(r.Context(), cookie.Value, r.RemoteAddr, r.UserAgent())
			if err != nil {
				failedAuthentication(w, err)
				return
			}

			uuid, err := uuid.Parse(userID)
			if err != nil {
				failedAuthentication(w, err)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", uuid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getUser(r *http.Request) uuid.UUID {
	return r.Context().Value("userID").(uuid.UUID)
}

func failedAuthentication(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	w.WriteHeader(http.StatusUnauthorized)
}
