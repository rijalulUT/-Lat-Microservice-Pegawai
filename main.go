package main

import (
	"fmt"
	"os"
	"pegawaimicroservice/config"
	"pegawaimicroservice/pegawai"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	r := gin.Default()
	db := config.DBInit()
	pegawai := pegawai.Pegawai{DB: db}

	port := os.Getenv("PORT")
	config.RegisterConsul()
	config.RegisterZipkin()

	r.GET("/pegawai", pegawai.GetPegawai)
	r.GET("/pegawaidetail", pegawai.GetPegawaiUnit)
	r.GET("/healthcheck", config.Healthcheck)
	r.Run(":" + port)
}
