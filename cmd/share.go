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
