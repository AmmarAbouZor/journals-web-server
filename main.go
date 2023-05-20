package main

import (
	"log"

	c "github.com/AmmarAbouZor/journals_web_server/controller"
	"github.com/AmmarAbouZor/journals_web_server/db"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Database error: %v", err)
	}

	router := gin.Default()

	journalGroup := router.Group("/journal")

	journalGroup.GET("", c.GetJournals)
	journalGroup.POST("", c.PostJournal)
	journalGroup.PUT("", c.PutJournal)
	journalGroup.DELETE("", c.DeleteJournal)

	//TODO: manage host and port with environment variables
	if routerErr := router.Run("localhost:8080"); routerErr != nil {
		log.Fatalf("Router err: %v", routerErr)
	}
}
