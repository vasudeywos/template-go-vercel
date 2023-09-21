package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/xuri/excelize/v2"
)
type Mess struct {
	Date      string   `json:"date"`
	Breakfast []string `json:"breakfast"`
	Lunch     []string `json:"lunch"`
	Dinner    []string `json:"dinner"`
}
func Json(w http.ResponseWriter, r *http.Request) {
	f, _ := excelize.OpenFile("ssms_test.xlsx")

	rows, _ := f.GetCols("Sheet1")

	var messArr []Mess

	for i, row := range rows {
		if i == 0 {
			continue
		}
		var b []string
		var l []string
		var d []string
		for j := 3; j < 12; j++ {
			b = append(b, row[j])
		}
		for n := 14; n <= 20; n++ {
			l = append(l, row[n])
		}
		for k := 24; k <= 30; k++ {
			d = append(d, row[k])
		}

		mess := Mess{
			Date:      row[1],
			Breakfast: b,
			Lunch:     l,
			Dinner:    d,
		}
		messArr = append(messArr, mess)
	}

	jsonResp, err := json.Marshal(messArr)
	if err != nil {
		fmt.Printf("Error happened in JSON marshal. Err: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
