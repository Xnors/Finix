package routes

import (
	"encoding/json"
	"log"
)

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

func Str2TableJSON(jsonData string) TableJSON {
	var data TableJSON
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	return data
}
