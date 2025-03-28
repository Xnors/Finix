package routes

import (
	"bytes"
	"finix-web/consts"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func RGet(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	table_name := c.Query("table_name")
	if table_name == "" {
		c.JSON(http.StatusOK, gin.H{"status": "err_empty_table_name"})
		return
	}

	cmd := exec.Command(consts.CLI_PATH(), "getdata", c.Query("table_name"))
	// 创建一个Buffer来捕获输出
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println("ran command failed: ", cmd.String(), out.String())
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":     "error",
				"table_name": table_name,
				"output":     strings.Trim(out.String(), "\n"),
				"error":      err.Error(),
			},
		)
		return
	}
	fmt.Println("读取了表: ", table_name)
	fmt.Println("输出: ", out.String())
	c.JSON(http.StatusOK, Str2TableJSON(out.String()))
}
