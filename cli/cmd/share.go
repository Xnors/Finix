package cmd

import "os"

const DATA_DIR = "./.data"
const TABLES_DIR = DATA_DIR + "/tables"

func CheckDataDirExistsOrCreate() {
	if _, err := os.Stat(DATA_DIR); os.IsNotExist(err) {
		os.Mkdir(DATA_DIR, 0755)
	}
	if _, err := os.Stat(TABLES_DIR); os.IsNotExist(err) {
		os.Mkdir(TABLES_DIR, 0755)
	}
}

func TablePath(tableName string) string {
	return TABLES_DIR + "/" + tableName + ".json"
}

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


