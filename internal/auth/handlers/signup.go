package authHandlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"
	supa "github.com/nedpals/supabase-go"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	ctx := context.Background()
	supabaseClient := supabase.InitClient()
	_user, err := supabaseClient.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if err != nil {
		fmt.Println("[SignupHandler:signup]", err)
	}
	dbRepository := supabase.NewSupabaseRepository(supabaseClient)

	userService := user.NewUserService(dbRepository)
	_err := userService.CreateUser(ctx, user.User{
		ID: _user.ID,
	})

	if _err != nil {
		fmt.Println("[SignupHandler:createuser]", _err)
	}
}
