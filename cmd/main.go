package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"caffinity/config"
	"caffinity/internal/db"
	"caffinity/internal/handler"
	"caffinity/internal/repository"
	"caffinity/internal/service"
)

func main() {
	// conn, err := sql.Open("postgres", "postgres://user:password@localhost:5432/dbname?sslmode=disable")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := config.Application.InitConfig(); err != nil {
		fmt.Printf("Error : %v\n", err.Error())
	}

	var connectionString = config.Application.DB

	queries := db.New(connectionString)

	repo := repository.NewCafeRepository(queries)
	service := service.NewCafeService(repo)
	handler := handler.NewCafeHandler(service)

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Welcome to " + config.Application.NAME + " API",
			"version": config.Application.VERSION,
		})
	})

	r.GET("/cafes", handler.GetCafes)

	log.Printf("Starting server on %s:%s", config.Application.HOST, config.Application.PORT)
	r.Run(config.Application.HOST + ":" + config.Application.PORT)
}
