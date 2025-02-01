package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func createTable(tb_name string) error {
	CheckDataDirExistsOrCreate()

	if _, err := os.Stat(TABLES_DIR + "/" + tb_name + ".toml"); !os.IsNotExist(err) {
		return errors.New("表已存在")
	}

	file, err := os.Create(TABLES_DIR + "/" + tb_name + ".toml")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf(`
[table-info]
created_at = %s

[records]
	`, time.Now().Format("2006-01-02 15:04:05")))
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
