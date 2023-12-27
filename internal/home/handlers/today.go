package home

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/emanueltimlopez/books-motivation/internal/user"

	supa "github.com/nedpals/supabase-go"
)

func daysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func TodayHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	tmpl := template.Must(template.ParseFiles("./web/templates/today.html"))
	ctx := context.Background()
	dbClient := supabase.InitClient()

	dbRepository := supabase.NewSupabaseRepository(dbClient)

	userService := user.NewUserService(dbRepository)

	_user, err := userService.GetUser(ctx, userSupa.ID)
	if err != nil {
		fmt.Println("[TodayHandler:getuser]", err)
	}

	if _user == nil {
		tmpl.Execute(w, nil)
		return
	}

	firstName := strings.Split(_user.Name, " ")[0]

	data := map[string]any{
		"Name": firstName,
	}

	tmpl.Execute(w, data)
}
