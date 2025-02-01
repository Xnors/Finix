package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "finix",
	Short: "这是个资金流动管理程序",
	Long: `这是一个资金流动管理程序

制作方: 异或科技工作室 https://xnors.us.kg`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("感谢使用 Finix! 使用 finix -h 查看帮助信息")
	},
}
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Finix 版本号: 0.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(CreateCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
