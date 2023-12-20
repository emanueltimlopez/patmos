package authHandlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	supa "github.com/nedpals/supabase-go"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	ctx := context.Background()
	supabase := supabase.InitClient()
	_user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		fmt.Println("[LoginHandler:signin]", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: _user.AccessToken,
	})

	w.Header().Set("HX-Redirect", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte{})
}
