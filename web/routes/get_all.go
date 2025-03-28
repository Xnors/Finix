package routes

import (
	"bytes"
	"finix-web/consts"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

type TablesInfo struct {
	// 表信息
	Name      string    `json:"name"`
	TableData TableJSON `json:"table_data"`
}

func str2ArrayTablesInfo(jsonData string) TablesInfo {
	fmt.Println("解析json数据 ",jsonData)
	var data []TablesInfo
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}
	fmt.Println("解析到json数据 ",data)
	return data[0]
}


func RGetAll(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	cmd := exec.Command(consts.CLI_PATH(), "get-all-tables")
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
