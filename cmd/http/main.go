package main

import (
	_ "excel/docs"
	"excel/internal/adapters/api"
	"excel/internal/adapters/config"
	"excel/internal/adapters/excel"
	"excel/internal/adapters/handlers"
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

}
