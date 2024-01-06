package authHandlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emanueltimlopez/patmos/internal/platform/supabase"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		fmt.Println("[LogoutHandler:cookie]", err)
	}

	ctx := context.Background()
	supabase := supabase.InitClient()
	_err := supabase.Auth.SignOut(ctx, c.Value)

	if _err != nil {
		fmt.Println("[LogoutHandler:signout]", _err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: "",
	})

	w.Header().Set("HX-Redirect", "/login")
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte{})
}
