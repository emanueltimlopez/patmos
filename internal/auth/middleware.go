package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	supa "github.com/nedpals/supabase-go"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, *supa.User)

type Auth struct {
	handler AuthenticatedHandler
}

func GetAuthenticatedUser(r *http.Request) (*supa.User, error) {
	c, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ctx := context.Background()
	supabase := supabase.InitClient()
	u, err := supabase.Auth.User(ctx, c.Value)
	if err != nil {
		fmt.Println(err)
	}
	return u, nil
}

func (ea *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	ea.handler(w, r, user)
}

func NewAuth(handler AuthenticatedHandler) *Auth {
	return &Auth{handler}
}
