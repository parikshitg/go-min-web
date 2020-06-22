package handlers

import (
	"net/http"

	s "github.com/parikshitg/go-min-web/session"
)

// Logout function
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := s.Store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
