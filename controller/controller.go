package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AmmarAbouZor/journals_web_server/db"
	m "github.com/AmmarAbouZor/journals_web_server/models"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	DB db.DB
}

func (ct *Controller) GetJournals(c *gin.Context) {
	journals, err := ct.DB.GetJournals()
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while retrieving journals"})
		return
	}
	c.IndentedJSON(http.StatusOK, journals)
}

func (ct *Controller) PostJournal(c *gin.Context) {
	var journal m.Journal

	if err := c.BindJSON(&journal); err != nil {
		return
	}

	id, err := ct.DB.AddJournal(journal)
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while creating journal"})
		return
	}

	journal.ID = id

	c.IndentedJSON(http.StatusCreated, journal)
}

func (ct *Controller) PutJournal(c *gin.Context) {
	var putJournal m.Journal

	if err := c.BindJSON(&putJournal); err != nil {
		return
	}

	_, err := ct.DB.UpdateJournal(&putJournal)
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while updating journal"})
		return
	}

	c.IndentedJSON(http.StatusOK, putJournal)
}

func (ct *Controller) DeleteJournal(c *gin.Context) {
	idString := c.Query("id")

	id, err := strconv.ParseInt(idString, 10, 0)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no journal id in request"})
		return
	}

	_, err = ct.DB.DeleteJournal(id)
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while deleting journal"})
		return
	}

	c.Status(http.StatusOK)
}
