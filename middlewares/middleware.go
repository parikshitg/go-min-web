package middlewares

import (
	"net/http"

	s "github.com/parikshitg/go-min-web/session"
)

// Checks if user is authenticated (middleware)
func AuthenticatedUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := s.Store.Get(r, "cookie-name")

		cookie, ok := session.Values["authenticated"]

		if cookie == true && ok {
			f(w, r)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	}
}

// checks unauthenticated user (middleware)
func UnauthenticatedUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := s.Store.Get(r, "cookie-name")

		cookie, ok := session.Values["authenticated"]

		if cookie == true && ok {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}

		f(w, r)

	}
}
