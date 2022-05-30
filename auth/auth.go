package auth

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

func LoggedInUser(s *sessions.CookieStore, r *http.Request) (uint, bool) {
	session, err := s.Get(r, "session")
	if err != nil {
		log.Fatalf("LoggedInUser: %v\n", err)
	}

	id, isset := session.Values["userID"]
	if !isset {
		return 0, false
	}

	userID, ok := id.(uint)
	if !ok {
		log.Fatalf("LoggedInUser: %v\n", err)
	}

	return userID, true
}
