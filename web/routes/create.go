package routes

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"strings"

	"finix-web/consts"

	"github.com/gin-gonic/gin"
)

func RCreate(c *gin.Context) {
	table_name := c.Query("table_name")
	if table_name == "" {
		c.JSON(http.StatusOK, gin.H{"status": "err_empty_name"})
		return
	}

	cmd := exec.Command(consts.CLI_PATH(), "create", c.Query("table_name"))
	// 创建一个Buffer来捕获输出
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println("ran command failed: ", cmd.String(), out.String())
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "error",
				"table_name": table_name,
				"output":     strings.Trim(out.String(), "\n"),
				"error":      err.Error(),
			},
		)
		return
	}
	fmt.Println(out.String())

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
