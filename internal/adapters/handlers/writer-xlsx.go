package handlers

import (
	"encoding/json"
	"escrituras/internal/domain/xlsx"
	"fmt"
	"net/http"
)

type writer struct {
	serviceWriter xlsx.Writer
}

// @tags Writer
// @title Excel Writer
// @description Create excel file
// @accept application/json
// @param Payload body XlsxRequest true "Payload"
// @Success 200	 {array} []string "ok"
// @router /writer [post]
func (wr *writer) Execute(w http.ResponseWriter, r *http.Request) {
	var payload xlsx.XlsxRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	wb, err := wr.serviceWriter.Write(payload.Sheets)
	if err != nil {
		w.Write([]byte(err.Error()))
		return

	}

	w.Header().Set("Content-Disposition", "attachment; filename=file.xlsx")
	w.Write(wb.Bytes())
	fmt.Println(payload)

}

func NewXlsxWriterHandler(service xlsx.Writer) *writer {
	return &writer{serviceWriter: service}
}
