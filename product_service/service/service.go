package service

import (
	"fmt"
	"product_service/db"
	"product_service/utils"

	"github.com/gin-gonic/gin"
)

func Run() {
	config, err := utils.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("config loading error: %v\n", err)
		return
	}

	err = db.Connect(&config.Database)
	defer db.Close()
	if err != nil {
		fmt.Printf("database connection error: %v\n", err)
		return
	}

	r := gin.Default()
	r.GET("/", func(c *gin.Context) { c.JSON(200, gin.H{"text": "monkey"}) })
	r.Run(config.Service.Host + ":" + config.Service.Port)
}
