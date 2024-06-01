package excel

import (
	"escrituras/internal/domain/xlsx"
	"io"

	"github.com/xuri/excelize/v2"
)

type xlsxReader struct{}

// Read implements xlsx.Reader.
func (x *xlsxReader) Read(r io.Reader) (result [][]string, err error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	for i := 0; i < f.SheetCount; i++ {
		sheetName := f.GetSheetName(i)
		rows, err := f.GetRows(sheetName)
		if err != nil {
			return result, err
		}
		result = append(result, rows...)

	}
	return
}

func NewXlsxReaderAdapter() xlsx.Reader {
	return &xlsxReader{}
}
