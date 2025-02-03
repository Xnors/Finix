package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func deleteTable(tb_name string) error {
	table_path := TablePath(tb_name)

	if tb_name == "" {
		return errors.New("表名不能为空")
	}
	if _, err := os.Stat(table_path); os.IsNotExist(err) {
		return fmt.Errorf("表 %s 不存在", tb_name)
	}

	if os.Remove(table_path) != nil {
		return fmt.Errorf("删除表 %s 失败", tb_name)
	}

	fmt.Printf("表 %s 删除成功\n", tb_name)
	return nil
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "删除表",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("请指定表名")
			os.Exit(1)
		}
		err := deleteTable(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
