package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"goth/types"
	"net/http"
	"time"

	"github.com/a-h/templ"
)

func RenderComponent(w http.ResponseWriter, r *http.Request, component templ.Component) {
	component.Render(r.Context(), w)
}

func Redirect(w http.ResponseWriter, r *http.Request, to string) {
	if len(r.Header.Get("Hx-Request")) > 0 {
		w.Header().Set("Hx-Redirect", to)
		w.WriteHeader(http.StatusSeeOther)
	}

	http.Redirect(w, r, to, http.StatusSeeOther)

}

func AddNotification(w http.ResponseWriter, t types.NotificationTypeKey, notification string) {
	newCookie := types.Notification{
		Value: notification,
		Type:  fmt.Sprint(t),
	}
	c, err := json.Marshal(newCookie)
	if err != nil {
		return
	}
	val := make([]byte, base64.StdEncoding.EncodedLen(len(c)))
	base64.StdEncoding.Encode(val, c)

	http.SetCookie(w, &http.Cookie{
		Name:    "notification",
		Value:   string(val),
		Path:    "/",
		Expires: time.Now().Add(2 * time.Second),
	})
}
