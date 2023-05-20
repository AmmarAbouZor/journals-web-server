package main

import (
	"github.com/AmmarAbouZor/journals_backend/journals"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	journalGroup := router.Group("/journal")

	journalGroup.GET("", journals.GetJournals)
	journalGroup.POST("", journals.PostJournal)
	journalGroup.PUT("", journals.PutJournal)
	journalGroup.DELETE("", journals.DeleteJournal)

	//TODO: manage host and port with environment variables
	router.Run("localhost:8080")
}
