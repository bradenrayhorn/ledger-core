package http

import (
	"context"
	"net/http"

	core "github.com/bradenrayhorn/ledger-core"
	"github.com/google/uuid"
)

func (s *Server) authentication() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session_id")
			if err != nil {
				s.HandleError(w, r, err)
				return
			}

			userID, err := s.SessionService.GetUserFromSession(r.Context(), cookie.Value, r.RemoteAddr, r.UserAgent())
			if err != nil {
				s.HandleError(w, r, core.WrapError(err, core.ErrorAuth))
				return
			}

			uuid, err := uuid.Parse(userID)
			if err != nil {
				s.HandleError(w, r, err)
				return
			}

			ctx := context.WithValue(r.Context(), "userID", uuid)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func isAuthenticated(r *http.Request) bool {
	userID := r.Context().Value("userID")
	return userID != nil && len(userID.(string)) > 0
}

func getUser(r *http.Request) uuid.UUID {
	return r.Context().Value("userID").(uuid.UUID)
}
