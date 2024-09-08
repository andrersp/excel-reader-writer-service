package main

import (
	_ "escrituras/docs"
	"escrituras/internal/adapters/api"
	"escrituras/internal/adapters/config"
	"escrituras/internal/adapters/excel"
	"escrituras/internal/adapters/handlers"
	"log/slog"
	"net/http"
	"os"

	"log"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)
	if err := config.LoadConfig(); err != nil {
		log.Fatal(err)
	}
}

// @title Excel Reader and Writer
// @version 1.0
// @description Service to read and writer xlsx.
// @termsOfService http://swagger.io/terms/
func main() {

	xlsxReader := excel.NewXlsxReaderAdapter()

	excelReadHandler := handlers.ReadExcelFileHandler(xlsxReader)
	excelWriterAdapter := excel.NewXlsxWriterAdapter()
	xlsxWriterHandler := handlers.NewXlsxWriterHandler(excelWriterAdapter)

	api := api.NewApi()
	api.Add(http.MethodPost, "/reader", excelReadHandler.Execute)
	api.Add(http.MethodPost, "/writer", xlsxWriterHandler.Execute)

	if err := api.Start(); err != nil {
		log.Fatal(err)
	}
	// http.HandleFunc("POST /read", ReadFile)

	// if err := startServer(); err != nil {
	// 	log.Fatal(err)
	// }

	// atos := make(map[string]ato.Ato)

	// // atos := make([]ato.Ato, 0)
	// f, err := excelize.OpenFile("planilha.xlsx")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer f.Close()

	// shhetName := f.GetSheetName(0)

	// rows, err := f.GetRows(shhetName)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// validRows := [][]string{}

	// for _, row := range rows {
	// 	// fmt.Println(len(row))

	// 	if (len(row)) >= 12 {
	// 		validRows = append(validRows, row)
	// 	}

	// }
	// for i := 1; i < len(validRows); i++ {

	// 	row := validRows[i]

	// 	ato := ato.Ato{
	// 		Tipo:        row[0],
	// 		Natureza:    row[1],
	// 		Data:        row[2],
	// 		Livro:       row[3],
	// 		Folha:       row[4],
	// 		Complemento: row[5],
	// 		Cartorio:    row[9],
	// 		Comarca:     row[10],
	// 		Uf:          row[11],
	// 		Partes:      make([]parte.Parte, 0),
	// 	}
	// 	part := parte.Parte{

	// 		Name:      row[6],
	// 		CpfCnpj:   row[7],
	// 		Qualidade: row[8],
	// 	}

	// 	key := fmt.Sprintf("%s%s", ato.Livro, ato.Folha)

	// 	if data, ok := atos[key]; ok {

	// 		ato := data
	// 		ato.Partes = append(ato.Partes, part)
	// 		atos[key] = ato

	// 	} else {
	// 		ato.Partes = append(ato.Partes, part)

	// 		atos[key] = ato

	// 	}

	// }

	// fmt.Println(len(validRows))
	// fmt.Println(len(atos))

}
