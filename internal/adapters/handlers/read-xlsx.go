package handlers

import (
	"escrituras/internal/domain/xlsx"
	"net/http"

	"github.com/labstack/echo/v4"
)

type readExcelFile struct {
	excel xlsx.Reader
}

// @tags Reader
// @title Excel Reader
// @description Reader excel file and return data
// @accept multipart/form-data
// @param file formData file true "this is a excel file"
// @Success 200	 {array} []string "ok"
// @router /reader [post]
func (re *readExcelFile) Execute(c echo.Context) error {

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	response, err := re.excel.Read(src)
	if err != nil {

		return err
	}
	return c.JSON(http.StatusOK, response)

}

func ReadExcelFileHandler(excel xlsx.Reader) readExcelFile {
	return readExcelFile{excel: excel}
}
