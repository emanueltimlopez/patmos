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

func GetParamsPlan(r *http.Request) PlanParams {
	r.ParseForm()
	minutes, err := strconv.Atoi(r.Form.Get("minutes"))
	if err != nil {
		fmt.Println("[GetParamsPlan:minutes]", err)
		minutes = 0
	}

	sessions, _err := strconv.Atoi(r.Form.Get("sessions"))
	if _err != nil {
		fmt.Println("[GetParamsPlan:sessions]", _err)
		sessions = 0
	}

	words, __err := strconv.Atoi(r.Form.Get("words"))
	if __err != nil {
		fmt.Println("[GetParamsPlan:words]", __err)
		words = 0
	}

	return PlanParams{
		Minutes:  minutes,
		Sessions: sessions,
		Words:    words,
	}
}
