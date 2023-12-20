package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	supa "github.com/nedpals/supabase-go"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	supabase := supabase.InitClient()
	_user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    "timbislopez@gmail.com",
		Password: "meneame",
	})

	if err != nil {
		fmt.Println(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: _user.AccessToken,
	})
}
