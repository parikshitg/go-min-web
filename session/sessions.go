package session

import (
	"github.com/gorilla/sessions"
)

// initialize session key
var key = []byte("super-secret-key")
var Store = sessions.NewCookieStore(key)
