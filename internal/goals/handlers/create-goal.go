package goals

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/emanueltimlopez/books-motivation/internal/goals"
	"github.com/emanueltimlopez/books-motivation/internal/platform/supabase"
	"github.com/google/uuid"
	supa "github.com/nedpals/supabase-go"
)

func CreateGoalHandler(w http.ResponseWriter, r *http.Request, userSupa *supa.User) {
	ctx := context.Background()
	dbClient := supabase.InitClient()
	dbRepository := supabase.NewSupabaseRepository(dbClient)

	r.ParseForm()
	title := r.Form.Get("title")
	quantity, err := strconv.Atoi(r.Form.Get("quantity"))
	if err != nil {
		fmt.Println(err)
	}
	days, err := strconv.Atoi(r.Form.Get("days"))
	if err != nil {
		fmt.Println(err)
	}

	goalsService := goals.NewGoalService(dbRepository)
	err = goalsService.CreateGoal(ctx, goals.Goal{
		Title:    title,
		Quantity: quantity,
		Days:     days,
		ID:       uuid.New().String(),
		UserID:   userSupa.ID,
	})

	if err != nil {
		fmt.Println(err)
	}
}
