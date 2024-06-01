package ato

import "escrituras/internal/domain/parte"

type Ato struct {
	Tipo        string
	Natureza    string
	Data        string
	Livro       string
	Folha       string
	Complemento string
	Cartorio    string
	Comarca     string
	Uf          string
	Partes      []parte.Parte
}
