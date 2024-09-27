package xlsx

type XlsxRequest struct {
	Sheets []Sheet `json:"sheets"`
} //@Name XlsxRequest

type Column struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Width float64 `json:"width"`
	Type  Type    `json:"type"`
} //@Name Column

type Sheet struct {
	Name    string                   `json:"name"`
	Columns []Column                 `json:"columns"`
	Data    []map[string]interface{} `json:"data"`
} //@Name Sheet

type Type int

// const ColumnType string
const (
	STRING Type = iota
	DATE
	FLOAT
	LIST
	MAP_BOOL
)
