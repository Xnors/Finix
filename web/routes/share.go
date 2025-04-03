package routes

import (
	"encoding/json"
	"fmt"
	"log"
)

type TableJSON struct {
	Info    TableInfo `json:"table-info"`
	Records []Record  `json:"records"`
}
type TableInfo struct {
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
}

type Record struct {
	Title     string  `json:"title"`
	Comment   string  `json:"comment"`
	CreatedAt string  `json:"created_at"`
	Change    float32 `json:"change"`
}

type TablesData struct {
	// 表信息
	Name      string    `json:"name"`
	TableData TableJSON `json:"table_data"`
}

type TablesInfo struct {
	// 表信息
	Name    string    `json:"name"`
	Info 	TableInfo `json:"table_info"`
}

func Str2TableJSON(jsonData string) TableJSON {
	var data TableJSON
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	return data
}

func str2ArrayTablesData(jsonData string) TablesData {
	fmt.Println("解析json数据 ", jsonData)
	var data []TablesData
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	fmt.Println("解析到json数据 ", data)
	return data[0]
}

func str2ArrayTablesInfo(jsonData string) []TablesData {
	fmt.Println("解析json数据 ", jsonData)
	var data []TablesData
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	fmt.Println("解析到json数据 ", data)
	return data
}