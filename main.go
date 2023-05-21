package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AmmarAbouZor/journals_web_server/controller"
	"github.com/AmmarAbouZor/journals_web_server/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Database error: %v", err)
	}
	defer db.CloseDB()

	router := gin.Default()

	journalGroup := router.Group("/journal")

	c := controller.Controller{DB: db}

	journalGroup.GET("", c.GetJournals)
	journalGroup.POST("", c.PostJournal)
	journalGroup.PUT("", c.PutJournal)
	journalGroup.DELETE("", c.DeleteJournal)

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("No PORT environment variable. Defaulting to 8080")
		port = "8080"
	}

	if routerErr := router.Run(fmt.Sprintf(":%v", port)); routerErr != nil {
		log.Fatalf("Router err: %v", routerErr)
	}
}
