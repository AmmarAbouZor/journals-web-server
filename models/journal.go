package models

import "time"

type Journal struct {
	ID      int64     `json:"id"`
	Title   string    `json:"title"`
	Date    time.Time `json:"date"`
	Content string    `json:"content"`
}
