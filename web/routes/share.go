package routes

type TableJSON struct {
	Info    TableInfo `json:"table-info"`
	Records []Record  `json:"records"`
}
type TableInfo struct {
	CreatedAt string `json:"created_at"`
}
type Record struct {
	Title     string  `json:"title"`
	Comment   string  `json:"comment"`
	CreatedAt string  `json:"created_at"`
	Change    float32 `json:"change"`
}

