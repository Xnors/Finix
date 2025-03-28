package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func get_all_tables() error {
	/*
		返回的是一个json文本格式的表信息
		格式:
		[
			{
				"name": "table1",
				"table_data": {...}
			},
			...
		]
	*/

	type TableInfo struct {
		// 表信息
		Name      string    `json:"name"`
		TableData TableJSON `json:"table_data"`
	}

	var totalTables []TableInfo

	// 读取 TABLES_DIR 目录下所有文件名，并返回一个包含所有文件名的切片
	files, err := os.ReadDir(TABLES_DIR)
	if err != nil {
		return err
	}

	for _, file := range files {
		fileName := file.Name()
		tableName := fileName[:len(fileName)-5] // 去掉.json后缀
		if (!file.IsDir()) && filepath.Ext(fileName) == ".json" {
			// 这里应该是读取文件内容，然后解析json，然后添加到totalTables中
			// 读取文件内容
			fileContent, err := os.ReadFile(TABLES_DIR + "/" + fileName)
			if err != nil {
				return err
			}

			// 解析json
			var tableJSON TableJSON
			err = json.Unmarshal(fileContent, &tableJSON)
			if err != nil {
				return err
			}

			// 添加到totalTables中
			totalTables = append(totalTables, TableInfo{
				Name:      tableName,
				TableData: tableJSON,
			})
		}

	}
	// 输出json格式的表信息
	jsonBytes, err := json.Marshal(totalTables)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))

	return nil
}

var GetAllCmd = &cobra.Command{
	Use:   "get-all-tables",
	Short: "获取所有表信息",
	Run: func(cmd *cobra.Command, args []string) {
		err := get_all_tables()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
