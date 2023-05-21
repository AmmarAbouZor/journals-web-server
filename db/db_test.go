package db

import (
	"os"
	"testing"
	"time"

	"github.com/AmmarAbouZor/journals_web_server/models"
	"github.com/stretchr/testify/assert"
)

var TestJournals = []models.Journal{
	{ID: 1, Title: "title 1", Date: time.Now().AddDate(-1, -1, -1), Content: "content 1"},
	{ID: 2, Title: "title 2", Date: time.Now(), Content: "content 2"},
	{ID: 3, Title: "title 3", Date: time.Now().AddDate(0, 2, -1), Content: "content 3"},
	{ID: 4, Title: "title 4", Date: time.Now().AddDate(0, 5, -1), Content: "content 4"},
	{ID: 5, Title: "title 5", Date: time.Now().AddDate(0, 5, 3), Content: "content 5"},
}

func setup() {
	// Root directory change between test start point and the normal start point
	migrationsPath = "file://../db/db_migrations"
	os.Setenv("DB_PATH", ":memory:")
}

func initWithDefaultValues() error {
	if err := InitDB(); err != nil {
		return err
	}

	for _, jnl := range TestJournals {
		_, err := AddJournal(jnl)
		if err != nil {
			return err
		}
	}

	return nil
}

func TestAddJournal(t *testing.T) {
	setup()

	err := initWithDefaultValues()
	assert.NoError(t, err, "initWithDefaultValues should not return an error")

	defer CloseDB()

	journal := models.Journal{
		Title:   "title 1",
		Date:    time.Now(),
		Content: "content 1",
	}

	id, err := AddJournal(journal)
	assert.NoError(t, err, "AddJournal should not return an error")
	assert.NotEqual(t, int64(-1), id, "AddJournal should return a valid ID")
}

func TestGetJournals(t *testing.T) {
	setup()

	err := initWithDefaultValues()
	assert.NoError(t, err, "initWithDefaultValues should not return an error")

	defer CloseDB()

	journals, err := GetJournals()
	assert.NoError(t, err, "GetJournals should not return an error")
	assert.Len(t, journals, len(TestJournals), "GetJournals should return the length of TestJournals")
}

func TestUpdateJournal(t *testing.T) {
	setup()

	err := initWithDefaultValues()
	assert.NoError(t, err, "initWithDefaultValues should not return an error")

	defer CloseDB()

	var jrnl = TestJournals[len(TestJournals)-1]
	jrnl.Content = "content changed"
	jrnl.Title = "title changed"

	affectedRows, err := UpdateJournal(&jrnl)
	assert.NoError(t, err, "UpdateJournal should not return an error")
	assert.Equal(t, int64(1), affectedRows, "UpdateJournal should update one row")

	journals, err := GetJournals()
	assert.NoError(t, err, "GetJournals should not return an error")

	for _, updatedJrnl := range journals {
		if updatedJrnl.ID == jrnl.ID {
			assert.Equal(t, updatedJrnl.Title, jrnl.Title, "Title should be updated")
			assert.Equal(t, updatedJrnl.Content, jrnl.Content, "Content should be updated")
			return
		}
	}
}

func TestDeleteJournal(t *testing.T) {
	setup()

	err := initWithDefaultValues()
	assert.NoError(t, err, "initWithDefaultValues should not return an error")

	defer CloseDB()

	id := TestJournals[0].ID

	affectedRows, err := DeleteJournal(id)
	assert.NoError(t, err, "DeleteJournal should not return an error")
	assert.Equal(t, int64(1), affectedRows, "DeleteJournal should delete one row")

	journals, err := GetJournals()
	assert.NoError(t, err, "GetJournals should not return an error")
	assert.Len(t, journals, len(TestJournals)-1, "GetJournals should return the length of TestJournals minus the deleted journal")
}
