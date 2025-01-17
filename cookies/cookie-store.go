package cookies

import (
	"encoding/gob"
	"goth/types"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

type CookieStore struct {
	Store *sessions.CookieStore
}

const (
	UserKey         = "user"
	UserEmail       = "user_email"
	UserAccessToken = "user_access_token"
)

func NewCookieStore() CookieStore {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	gob.Register(types.AuthenticatedUser{})
	return CookieStore{
		Store: store,
	}
}

func (cs CookieStore) GetUserSession(r *http.Request) *sessions.Session {
	sess, _ := cs.Store.Get(r, UserKey)

	return sess
}
