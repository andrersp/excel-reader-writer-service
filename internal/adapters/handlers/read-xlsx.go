package handlers

import (
	"encoding/json"
	"escrituras/internal/domain/xlsx"
	"net/http"
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
func (re *readExcelFile) Execute(w http.ResponseWriter, r *http.Request) {

	f, fh, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer f.Close()

	if fh.Size <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("err"))
		return

	}

	response, err := re.excel.Read(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func ReadExcelFileHandler(excel xlsx.Reader) readExcelFile {
	return readExcelFile{excel: excel}
}
