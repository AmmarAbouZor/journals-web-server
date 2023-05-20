package models

import "time"

type Journal struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
}

var TestJournals = []Journal{
	{1, "title 1", time.Now().AddDate(-1, -1, -1), "content 1"},
	{2, "title 2", time.Now(), "content 2"},
	{3, "title 3", time.Now().AddDate(0, 2, -1), "content 3"},
	{4, "title 4", time.Now().AddDate(0, 5, -1), "content 4"},
	{5, "title 5", time.Now().AddDate(0, 5, 3), "content 5"},
}
