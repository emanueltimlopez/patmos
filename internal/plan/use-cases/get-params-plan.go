package plan_usecases

import (
	"fmt"
	"net/http"
	"strconv"
)

type PlanParams struct {
	Minutes  int
	Sessions int
	Words    int
}

func GetParamsSearchBook(r *http.Request) PlanParams {
	r.ParseForm()
	minutes, err := strconv.Atoi(r.Form.Get("minutes"))
	if err != nil {
		fmt.Println("[GetParamsSearchBook:minutes]", err)
	}

	sessions, _err := strconv.Atoi(r.Form.Get("sessions"))
	if _err != nil {
		fmt.Println("[GetParamsSearchBook:sessions]", _err)
	}

	words, __err := strconv.Atoi(r.Form.Get("words"))
	if __err != nil {
		fmt.Println("[GetParamsSearchBook:words]", __err)
	}

	return PlanParams{
		Minutes:  minutes,
		Sessions: sessions,
		Words:    words,
	}
}
