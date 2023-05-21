package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/AmmarAbouZor/journals_web_server/controller"
	m "github.com/AmmarAbouZor/journals_web_server/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockDB is a mock implementation of the DB interface
type MockDB struct {
}

func (m *MockDB) CloseDB() error {
	return nil
}

func (m *MockDB) AddJournal(journal m.Journal) (int64, error) {
	return 2, nil
}

func (mock *MockDB) GetJournals() ([]m.Journal, error) {
	testJournals := []m.Journal{
		{ID: 1, Title: "title 1", Date: time.Now().AddDate(-1, -1, -1), Content: "content 1"},
		{ID: 2, Title: "title 2", Date: time.Now(), Content: "content 2"},
	}
	return testJournals, nil
}

func (m *MockDB) UpdateJournal(journal *m.Journal) (int64, error) {
	return 1, nil
}

func (m *MockDB) DeleteJournal(id int64) (int64, error) {
	return 1, nil
}

func TestController_GetJournals(t *testing.T) {
	mockDB := &MockDB{}
	ct := &controller.Controller{
		DB: mockDB,
	}

	router := gin.Default()
	router.GET("/journals", ct.GetJournals)

	// Perform a GET request to /journals
	req, _ := http.NewRequest("GET", "/journals", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Decode the response body
	var journals []m.Journal
	err := json.Unmarshal(rec.Body.Bytes(), &journals)
	assert.NoError(t, err)

	// Check the number of journals returned
	assert.Len(t, journals, 2)

	// Check the content of the first journal
	assert.Equal(t, int64(1), journals[0].ID)
	assert.Equal(t, "title 1", journals[0].Title)
	assert.Equal(t, "content 1", journals[0].Content)

	// Check the content of the second journal
	assert.Equal(t, int64(2), journals[1].ID)
	assert.Equal(t, "title 2", journals[1].Title)
	assert.Equal(t, "content 2", journals[1].Content)
}

func TestController_PostJournal(t *testing.T) {
	mockDB := &MockDB{}
	ct := &controller.Controller{
		DB: mockDB,
	}

	router := gin.Default()
	router.POST("/journals", ct.PostJournal)

	journal := m.Journal{
		Title:   "New Journal",
		Date:    time.Date(2023, 5, 22, 0, 0, 0, 0, time.UTC),
		Content: "New content",
	}
	body, _ := json.Marshal(journal)

	// Perform a POST request to /journals
	req, _ := http.NewRequest("POST", "/journals", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)

	// Decode the response body
	var createdJournal m.Journal
	err := json.Unmarshal(rec.Body.Bytes(), &createdJournal)
	assert.NoError(t, err)

	// Check the ID of the created journal
	assert.Equal(t, int64(2), createdJournal.ID)
	assert.Equal(t, "New Journal", createdJournal.Title)
	assert.Equal(t, "New content", createdJournal.Content)
}

func TestController_PutJournal(t *testing.T) {
	mockDB := &MockDB{}
	ct := &controller.Controller{
		DB: mockDB,
	}

	router := gin.Default()
	router.PUT("/journals", ct.PutJournal)

	// Create a request body
	journal := m.Journal{
		ID:      1,
		Title:   "Updated Journal",
		Date:    time.Date(2023, 5, 22, 0, 0, 0, 0, time.UTC),
		Content: "Updated content",
	}
	body, _ := json.Marshal(journal)

	// Perform a PUT request to /journals
	req, _ := http.NewRequest("PUT", "/journals", bytes.NewBuffer(body))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var updatedJournal m.Journal
	err := json.Unmarshal(rec.Body.Bytes(), &updatedJournal)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), updatedJournal.ID)
	assert.Equal(t, "Updated Journal", updatedJournal.Title)
	assert.Equal(t, "Updated content", updatedJournal.Content)
}

func TestController_DeleteJournal(t *testing.T) {
	mockDB := &MockDB{}
	ct := &controller.Controller{
		DB: mockDB,
	}

	router := gin.Default()
	router.DELETE("/journals", ct.DeleteJournal)

	// Perform a DELETE request to /journals with the ID query parameter
	req, _ := http.NewRequest("DELETE", "/journals?id=1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
