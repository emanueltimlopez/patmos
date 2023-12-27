package authHandlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/emanueltimlopez/books-motivation/internal/plan"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"
	supa "github.com/nedpals/supabase-go"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	name := r.Form.Get("name")
	year, err := strconv.Atoi(r.Form.Get("year"))
	if err != nil {
		fmt.Println("[SignupHandler:year]", err)
	}

	ctx := context.Background()
	supabaseClient := supabase.InitClient()
	_, _err := supabaseClient.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if _err != nil {
		fmt.Println("[SignupHandler:signup]", _err)
	}
	__user, __err := supabaseClient.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    email,
		Password: password,
	})

	if __err != nil {
		fmt.Println("[SignupHandler:signin]", __err)
	}

	dbRepository := supabase.NewSupabaseRepository(supabaseClient)

	userService := user.NewUserService(dbRepository)
	___err := userService.CreateUser(ctx, user.User{
		ID:    __user.User.ID,
		Birth: year,
		Name:  name,
		Plan: plan.Plan{
			Minutes:  15,
			Sessions: 1,
			Words:    250,
		},
	})

	if ___err != nil {
		fmt.Println("[SignupHandler:createuser]", ___err)
	}

	tmpl := template.Must(template.ParseFiles("./web/templates/components/signup-ok.html"))
	tmpl.Execute(w, nil)
}
