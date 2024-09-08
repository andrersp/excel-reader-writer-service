package handlers

import (
	"excel/internal/domain/xlsx"

	"github.com/labstack/echo/v4"
)

type writer struct {
	serviceWriter xlsx.Writer
}

// @tags Writer
// @title Excel Writer
// @description Create excel file
// @accept application/json
// @param Payload body XlsxRequest true "Payload"
// @Success 200 {file} file
// @produce application/vnd.ms-excel
// @router /writer [post]
func (wr *writer) Execute(c echo.Context) error {
	var payload xlsx.XlsxRequest

	if err := c.Bind(&payload); err != nil {
		return err
	}

	wb, err := wr.serviceWriter.Write(payload.Sheets)
	if err != nil {
		return err
	}

	return c.Blob(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", wb.Bytes())

}

func NewXlsxWriterHandler(service xlsx.Writer) *writer {
	return &writer{serviceWriter: service}
}
