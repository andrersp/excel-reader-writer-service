package main

import (
	"excel/internal/adapters/excel"
	"excel/internal/domain/xlsx"
	"log"
	"time"
)

func main() {

	column := []xlsx.Column{
		{
			Title: "Nome",
			Id:    "nome",
			Type:  xlsx.STRING,
		},
		{
			Title: "Valor",
			Id:    "valor",
			Type:  xlsx.FLOAT,
		},
		{
			Title: "Nascimento",
			Id:    "nascimento",
			Type:  xlsx.DATE,
		},
		{
			Title: "Pessoas",
			Id:    "pessoas",
			Type:  xlsx.LIST,
		},
		{
			Title: "Data de Nascimento (PF)/Data de Abertura (PJ)",
			Id:    "nascimento",
			Type:  xlsx.DATE,
		},
	}

	datas := []map[string]any{
		{"nome": "Andre", "valor": 2.50},
		{"nome": "Andre", "nascimento": time.Now()},
		{"nome": "Andre", "pessoas": []string{"-André", "-Luis"}},
	}

	sheet := xlsx.Sheet{
		Name:    "Processos",
		Columns: column,
		Data:    datas,
	}

	adapter := excel.NewXlsxWriterAdapter()
	_, err := adapter.Write([]xlsx.Sheet{sheet})
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(r)

}
