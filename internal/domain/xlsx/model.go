package xlsx

type XlsxRequest struct {
	Sheets []Sheet `json:"sheets"`
} //@Name XlsxRequest

type Data map[string]interface{}

type Column struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Width float64 `json:"width"`
	Type  string  `json:"type"`
} //@Name Column

// type AdditionalData struct {
// 	Title string `json:"title"`
// 	Value string `json:"value"`
// 	Type  string `json:"type"`
// }

type Sheet struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
	// AdditionalInfo *[]AdditionalData `json:"additioanlData"`
	Data []Data `json:"data"`
} //@Name Sheet
