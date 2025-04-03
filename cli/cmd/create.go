package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func createTable(tb_name string) error {
	CheckDataDirExistsOrCreate()
	table_path := TablePath(tb_name)

	if _, err := os.Stat(table_path); !os.IsNotExist(err) {
		return errors.New("表已存在")
	}

	file, err := os.Create(table_path)
	if err != nil {
		return err
	}
	defer file.Close()

	init_struct := TableJSON{
		TableInfo{
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
		[]Record{},
	}

	init_string, err := json.Marshal(init_struct)
	if err != nil {
		return err
	}

	_, err = file.WriteString(string(init_string))
	if err != nil {
		return err
	}

	fmt.Printf("表 %s 创建成功\n", tb_name)
	return nil
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建新表",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("请指定表名")
			os.Exit(1)
		}
		err := createTable(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
