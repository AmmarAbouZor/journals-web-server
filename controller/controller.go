package controller

import (
	"net/http"
	"strconv"

	m "github.com/AmmarAbouZor/journals_web_server/models"
	"github.com/gin-gonic/gin"
)

func GetJournals(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, m.TestJournals)
}

func PostJournal(c *gin.Context) {
	var journal m.Journal

	if err := c.BindJSON(&journal); err != nil {
		return
	}

	journal.ID = m.TestJournals[len(m.TestJournals)-1].ID + 1

	m.TestJournals = append(m.TestJournals, journal)

	c.IndentedJSON(http.StatusCreated, journal)
}

func PutJournal(c *gin.Context) {
	var putJournal m.Journal

	if err := c.BindJSON(&putJournal); err != nil {
		return
	}

	for idx := range m.TestJournals {
		journal := &m.TestJournals[idx]

		if putJournal.ID == journal.ID {
			journal.Title = putJournal.Title
			journal.Date = putJournal.Date
			journal.Content = putJournal.Content

			c.IndentedJSON(http.StatusOK, journal)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "journal not found"})
}

func DeleteJournal(c *gin.Context) {
	idString := c.Query("id")

	id, err := strconv.ParseInt(idString, 10, 0)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no journal id in request"})
		return
	}

	delIdx := -1

	for index := range m.TestJournals {
		if m.TestJournals[index].ID == id {
			delIdx = index
			break
		}
	}

	m.TestJournals = append(m.TestJournals[:delIdx], m.TestJournals[delIdx+1:]...)

	c.Status(http.StatusOK)
}
