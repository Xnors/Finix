package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func createTable(tb_name string) error {
	CheckDataDirExistsOrCreate()

	if _, err := os.Stat(TABLES_DIR + "/" + tb_name + ".toml"); !os.IsNotExist(err) {
		return errors.New("表已存在")
	}

	_, err := os.Create(TABLES_DIR + "/" + tb_name + ".toml")
	if err != nil {
		return err
	}
	return nil
}

var CreateCmd = &cobra.Command{
	Use:   "create",
	Short: "创建新表",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("请指定表名")
			return
		}
		err := createTable(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}
