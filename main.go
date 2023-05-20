package main

import (
	"log"

	"github.com/AmmarAbouZor/journals_web_server/journals"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := journals.InitDB(); err != nil {
		log.Fatalf("Database error: %v", err)
	}

	router := gin.Default()

	journalGroup := router.Group("/journal")

	journalGroup.GET("", journals.GetJournals)
	journalGroup.POST("", journals.PostJournal)
	journalGroup.PUT("", journals.PutJournal)
	journalGroup.DELETE("", journals.DeleteJournal)

	//TODO: manage host and port with environment variables
	if routerErr := router.Run("localhost:8080"); routerErr != nil {
		log.Fatalf("Router err: %v", routerErr)
	}
}
