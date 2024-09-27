package excel

import (
	"bufio"
	"bytes"
	"excel/internal/domain/xlsx"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/xuri/excelize/v2"
)

var collLetters = [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"AA", "AB", "AC", "AD", "AE", "AF", "AG", "AH", "AI", "AJ", "AK", "AL", "AM", "AN", "AO", "AP", "AQ", "AR", "AS", "AT", "AU", "AV", "AW", "AX", "AY", "AZ",
	"BA", "BB", "BC", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BK", "BL", "BM", "BN", "BO", "BP", "BQ", "BR", "BS", "BT", "BU", "BV", "BW", "BX", "BY", "BZ",
	"CA", "CB", "CC", "CD", "CE", "CF", "CG", "CH", "CI", "CJ", "CK", "CL", "CM", "CN", "CO", "CP", "CQ", "CR", "CS", "CT", "CU", "CV", "CW", "CX", "CY", "CZ",
	"DA", "DB", "DC", "DD", "DE", "DF", "DG", "DH", "DI", "DJ", "DK", "DL", "DM", "DN", "DO", "DP", "DQ", "DR", "DS", "DT", "DU", "DV", "DW", "DX", "DY", "DZ"}

type fileStyles struct {
	defaultStyle int
	dateStyle    int
	floatStyle   int
	titleStyle   int
}

type xlsxWriter struct{}

// Write implements xlsx.Writer.
func (x *xlsxWriter) Write(sheets []xlsx.Sheet) (bytes.Buffer, error) {

	var b bytes.Buffer
	f := excelize.NewFile()

	x.createSheets(f, sheets)
	f.SaveAs("file.xlsx")

	writer := bufio.NewWriter(&b)
	if err := f.Write(writer); err != nil {
		return b, err
	}

	return b, nil
}

func (x *xlsxWriter) createSheets(f *excelize.File, sheets []xlsx.Sheet) {
	styles := fileStyles{}
	titleStyle, _ := f.NewStyle(TitleStyle)
	defaultStyle, _ := f.NewStyle(DefaultStyle)
	dateStyle, _ := f.NewStyle(DateStyle)
	floatStyle, _ := f.NewStyle(FloatStyle)
	styles.titleStyle = titleStyle
	styles.defaultStyle = defaultStyle
	styles.dateStyle = dateStyle
	styles.floatStyle = floatStyle

	for _, sheet := range sheets {
		x.createSheet(f, sheet, &styles)
	}

}

func (x *xlsxWriter) createSheet(f *excelize.File, sheet xlsx.Sheet, styles *fileStyles) {
	currentRowIndex := 1
	sheetName := sheet.Name
	if len(sheetName) > 31 {
		sheetName = sheetName[0:30]
	}
	f.NewSheet(sheetName)
	f.DeleteSheet("Sheet1")

	sw, err := f.NewStreamWriter(sheetName)
	if err != nil {
		log.Fatal(err)
	}
	currentRowIndex = x.setTitle(sw, sheet.Columns, currentRowIndex, styles.titleStyle)
	x.setTable(sw, currentRowIndex, sheet.Columns, sheet.Data, *styles)
	sw.Flush()

}

func (x *xlsxWriter) setTitle(sw *excelize.StreamWriter, columns []xlsx.Column, currentRowIndex int, styleId int) int {
	titles := []interface{}{}

	for index, title := range columns {
		col := index + 1
		if err := sw.SetColWidth(col, col, title.Width); err != nil {
			log.Fatal(err)
		}
		cel := excelize.Cell{
			Value:   title.Title,
			StyleID: styleId,
		}
		titles = append(titles, cel)
	}
	cell := x.cell(currentRowIndex, 0)
	if err := sw.SetRow(cell, titles); err != nil {
		slog.Error("Erro ao criar titulo", "err", err.Error())
	}

	currentRowIndex++
	return currentRowIndex
}

func (x *xlsxWriter) setTable(sw *excelize.StreamWriter, currentRowIndex int, columns []xlsx.Column, data []map[string]interface{}, styles fileStyles) int {
	for i := 0; i < len(data); i++ {
		cell := x.cell(currentRowIndex, 0)
		rows := make([]interface{}, 0)
		for _, col := range columns {
			var cel interface{}
			value := data[i][col.Id]
			switch col.Type {
			case xlsx.FLOAT:
				cel = x.getCell(value, styles.floatStyle)
			case xlsx.DATE:
				value, err := time.Parse(time.RFC3339, fmt.Sprintf("%s", value))
				if err == nil {
					cel = x.getCell(value, styles.dateStyle)
				}
			case xlsx.LIST, xlsx.MAP_BOOL:
				cel = x.getListRichText(value)
			default:
				cel = x.getCell(value, styles.defaultStyle)

			}
			rows = append(rows, cel)
		}
		if err := sw.SetRow(cell, rows); err != nil {
			slog.Error("Erro ao criar tabelas", "err", err.Error())
		}
		currentRowIndex++

	}

	return currentRowIndex
}

func (*xlsxWriter) getCell(value interface{}, styleId int) excelize.Cell {

	return excelize.Cell{
		Value:   value,
		StyleID: styleId,
	}
}

func (*xlsxWriter) getListRichText(value interface{}) []excelize.RichTextRun {
	richTexts := make([]excelize.RichTextRun, 0)

	switch output := value.(type) {
	case map[string]interface{}:
		for key, value := range output {
			richText := excelize.RichTextRun{
				Text: fmt.Sprintf("%s\n", key),
				Font: &excelize.Font{
					Family: "Century Gothic",
					Size:   8,
				},
			}
			if bold, ok := value.(bool); ok {
				richText.Font.Bold = bold
			}
			richTexts = append(richTexts, richText)

		}
	case []interface{}:
		{
			for _, value := range output {
				richText := excelize.RichTextRun{
					Text: fmt.Sprintf("%s\n", value),
					Font: &excelize.Font{
						Family: "Century Gothic",
						Size:   8,
					},
				}
				richTexts = append(richTexts, richText)

			}
		}

	}

	return richTexts

}

func (*xlsxWriter) getMapBoolRichText(value interface{}) []excelize.RichTextRun {
	richTexts := make([]excelize.RichTextRun, 0)

	switch output := value.(type) {
	case map[string]interface{}:
		for key, value := range output {
			richText := excelize.RichTextRun{
				Text: fmt.Sprintf("%s\n", key),
				Font: &excelize.Font{
					Family: "Century Gothic",
					Size:   8,
				},
			}
			if bold, ok := value.(bool); ok {
				richText.Font.Bold = bold
			}
			richTexts = append(richTexts, richText)

		}

	}

	return richTexts

}

func (*xlsxWriter) SheetName(f *excelize.File, sheetIndex int) string {
	return f.GetSheetName(sheetIndex)
}

func (x *xlsxWriter) cell(currentRowIndex, coll int) string {

	return fmt.Sprintf("%s%d", collLetters[coll], currentRowIndex)
}
func NewXlsxWriterAdapter() xlsx.Writer {
	return &xlsxWriter{}
}
