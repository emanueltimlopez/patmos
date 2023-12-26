package home

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
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

	/*currentTime := time.Now()
	currentMonth := currentTime.Month()
	numberOfDays := daysInMonth(currentTime.Year(), currentTime.Month())

	habitData := make([]int, numberOfDays)
	for i := 0; i < numberOfDays; i++ {
		habitData[i] = 0
	}

	month := map[string]any{
		"Current": currentMonth.String(),
		"Days":    numberOfDays,
		"Habit":   habitData,
	}
	tmpl.Execute(w, month)*/

	tmpl.Execute(w, nil)
}
