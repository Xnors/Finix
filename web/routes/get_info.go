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

func RGetInfo(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	cmd := exec.Command(consts.CLI_PATH(), "get-info")
	// 创建一个Buffer来捕获输出
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println("ran command failed: ", cmd.String(), " \n|#!>  ", out.String())
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "error",
				"output": strings.Trim(out.String(), "\n"),
				"error":  err.Error(),
			},
		)
		return
	}
	fmt.Println("输出: ", out.String())
	c.JSON(http.StatusOK, str2ArrayTablesInfo(out.String()))
}