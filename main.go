package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kntaka/go-sample-api/infrastructure/api/router"
	"github.com/kntaka/go-sample-api/infrastructure/persistence/datastore"
	"github.com/kntaka/go-sample-api/registry"
	"log"
)

func main() {
	// read .env
	envLoad()

	// DB取得
	conn := datastore.NewMySqlDB()
	// db stop
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	// handler
	h := registry.InitializeApp(conn)

	// gin
	g := gin.Default()

	// router
	router.NewRouter(g, h)
	g.Run(":8000")
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
