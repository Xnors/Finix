package routes

import (
	"encoding/json"
	"log"
	"fmt"
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

type TablesInfo struct {
	// 表信息
	Name      string    `json:"name"`
	TableData TableJSON `json:"table_data"`
}

func Str2TableJSON(jsonData string) TableJSON {
	var data TableJSON
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	return data
}


func str2ArrayTablesInfo(jsonData string) TablesInfo {
	fmt.Println("解析json数据 ", jsonData)
	var data []TablesInfo
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	fmt.Println("解析到json数据 ", data)
	return data[0]
}

