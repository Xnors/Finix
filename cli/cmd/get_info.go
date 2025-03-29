package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

func get_info() error {
	/*
		返回的是一个json文本格式的表信息
		格式:
		[
			{
				"name": "table1",
				"table_info": {created_at: "2021-01-01 12:00:00"},
			},
			...
		]
	*/

	type tablesInfo struct {
		// 表信息
		Name string    `json:"name"`
		Info TableInfo `json:"table_info"`
	}

	var totalTables []tablesInfo

	// 读取 TABLES_DIR 目录下所有文件
	files, err := os.ReadDir(TABLES_DIR)
	if err != nil {
		return err
	}
	// 遍历所有文件
	for _, file := range files {
		// 跳过非json文件
		if !file.IsDir() && filepath.Ext(file.Name()) != ".json" {
			continue
		}
		// 读取文件内容
		content, err := os.ReadFile(TABLES_DIR + "/" + file.Name())
		if err != nil {
			return err
		}
		// 解析json内容
		var tableJSON TableJSON
		err = json.Unmarshal(content, &tableJSON)
		if err != nil {
			return err
		}
		// 构造表信息
		tableInfo := tablesInfo{
			Name: file.Name(),
			Info: TableInfo{
				CreatedAt: tableJSON.Info.CreatedAt,
				Description: tableJSON.Info.Description,
			},
		}
		// 追加到总表信息
		totalTables = append(totalTables, tableInfo)
	}
	// 输出json格式的表信息
	jsonBytes, err := json.Marshal(totalTables)
	if err != nil {
		return err
	}
	fmt.Println(string(jsonBytes))

	return nil
}

var GetInfoCmd = &cobra.Command{
	Use:   "get-info",
	Short: "获取所有表名与表的信息(创建时间等)",
	Run: func(cmd *cobra.Command, args []string) {
		err := get_info()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
