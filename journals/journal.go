package journals

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Journal struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
}

var testJournals = []Journal{
	{1, "title 1", time.Now().AddDate(-1, -1, -1), "content 1"},
	{2, "title 2", time.Now(), "content 2"},
	{3, "title 3", time.Now().AddDate(0, 2, -1), "content 3"},
	{4, "title 4", time.Now().AddDate(0, 5, -1), "content 4"},
	{5, "title 5", time.Now().AddDate(0, 5, 3), "content 5"},
}

func GetJournals(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, testJournals)
}

func PostJournal(c *gin.Context) {
	var journal Journal

	if err := c.BindJSON(&journal); err != nil {
		return
	}

	journal.ID = testJournals[len(testJournals)-1].ID + 1

	testJournals = append(testJournals, journal)

	c.IndentedJSON(http.StatusCreated, journal)
}

func PutJournal(c *gin.Context) {
	var putJournal Journal

	if err := c.BindJSON(&putJournal); err != nil {
		return
	}

	for idx := range testJournals {
		journal := &testJournals[idx]

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
	idString := c.Param("id")

	id, err := strconv.ParseInt(idString, 10, 0)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "no journal id in request"})
		return
	}

	delIdx := -1

	for index := range testJournals {
		if testJournals[index].ID == id {
			delIdx = index
			return
		}
	}

	testJournals = append(testJournals[:delIdx], testJournals[delIdx+1:]...)

	c.Status(http.StatusOK)
}
