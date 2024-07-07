package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhaohaihang/k8s-manage/cmd/app"
	"os"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
