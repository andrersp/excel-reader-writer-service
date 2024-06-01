package excel

import (
	"bufio"
	"bytes"
	"escrituras/internal/domain/xlsx"
	"fmt"
	"io"
	"log"
	"log/slog"

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

type excelAdapter struct{}

func (e *excelAdapter) Read(r io.Reader) (result [][]string, err error) {

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

func (e *excelAdapter) Write(sheets []xlsx.Sheet) (bytes.Buffer, error) {

	var b bytes.Buffer
	f := excelize.NewFile()

	e.createSheets(f, sheets)

	writer := bufio.NewWriter(&b)
	if err := f.Write(writer); err != nil {
		return b, err
	}

	return b, nil
}

func (e *excelAdapter) createSheets(f *excelize.File, sheets []xlsx.Sheet) {
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
		e.createSheet(f, sheet, &styles)
	}

}

func (e *excelAdapter) createSheet(f *excelize.File, sheet xlsx.Sheet, styles *fileStyles) {
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
	currentRowIndex = e.setTitle(sw, sheet.Columns, currentRowIndex, styles.titleStyle)
	e.setTable(sw, currentRowIndex, sheet.Columns, sheet.Data, *styles)
	sw.Flush()

}

func (e *excelAdapter) setTitle(sw *excelize.StreamWriter, columns []xlsx.Column, currentRowIndex int, styleId int) int {
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
	cell := e.cell(currentRowIndex, 0)
	if err := sw.SetRow(cell, titles); err != nil {
		slog.Error("Erro ao criar titulo", "err", err.Error())
	}

	currentRowIndex++
	return currentRowIndex
}

func (e *excelAdapter) setTable(sw *excelize.StreamWriter, currentRowIndex int, columns []xlsx.Column, data []xlsx.Data, styles fileStyles) int {
	for i := 0; i < len(data); i++ {
		cell := e.cell(currentRowIndex, 0)
		rows := make([]interface{}, 0)
		for _, col := range columns {
			var cel interface{}
			value := data[i][col.Id]
			switch col.Type {
			case "moeda":
				cel = e.getCell(value, styles.floatStyle)
			case "data":
				cel = e.getCell(value, styles.dateStyle)
			case "mapBool":
				cel = e.getRichTextCell(value)
			default:
				cel = e.getCell(value, styles.defaultStyle)

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

func (e *excelAdapter) getCell(value interface{}, styleId int) excelize.Cell {

	return excelize.Cell{
		Value:   value,
		StyleID: styleId,
	}
}

func (*excelAdapter) getRichTextCell(value interface{}) []excelize.RichTextRun {
	data := value.(map[string]bool)
	richTexts := make([]excelize.RichTextRun, 0)
	for key, value := range data {
		richText := excelize.RichTextRun{
			Text: fmt.Sprintf("%s\n", key),
			Font: &excelize.Font{
				Family: "Century Gothic",
				Size:   8,
				Bold:   value,
			},
		}
		richTexts = append(richTexts, richText)

	}

	return richTexts

}

func (*excelAdapter) SheetName(f *excelize.File, sheetIndex int) string {
	return f.GetSheetName(sheetIndex)
}

func (e *excelAdapter) cell(currentRowIndex, coll int) string {

	return fmt.Sprintf("%s%d", collLetters[coll], currentRowIndex)
}

func NewExcelAdapter() xlsx.Reader {
	return &excelAdapter{}
}

func NewXlsxWriterAdapter() xlsx.Writer {
	return &excelAdapter{}
}
