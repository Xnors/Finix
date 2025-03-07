package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func getdata(tb_name string) error {
	table_path := TablePath(tb_name)

	if tb_name == "" {
		return errors.New("表名不能为空")
	}
	if _, err := os.Stat(table_path); os.IsNotExist(err) {
		return fmt.Errorf("表 %s 不存在", tb_name)
	}

	if file_text, err := os.ReadFile(table_path); err == nil {
		fmt.Println(string(file_text))
	} else {
		return errors.New("读取文件失败")
	}
	return nil
}

var GetdataCmd = &cobra.Command{
	Use:   "getdata",
	Short: "获取数据",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("请指定表名")
			os.Exit(1)
		}
		err := getdata(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
