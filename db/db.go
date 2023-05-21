package db

import (
	"database/sql"
	"fmt"
	"os"

	m "github.com/AmmarAbouZor/journals_web_server/models"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	migratSqlit "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DB interface {
	CloseDB() error
	AddJournal(journal m.Journal) (int64, error)
	GetJournals() ([]m.Journal, error)
	UpdateJournal(journal *m.Journal) (int64, error)
	DeleteJournal(id int64) (int64, error)
}

type SqliteDB struct {
	dbConnection *sql.DB
}

// migratoinsPath can't be const since it will be overwriten in unit tests
var migrationsPath = "file://db/db_migrations"

func InitDB() (DB, error) {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		fmt.Println("DB_PATH env varialbe not set. Defaulting to 'journals.db'")
		dbPath = "journals.db"
	}

	db := &SqliteDB{}

	var err error
	db.dbConnection, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("Oppining/creating database failed: %v", err)
	}

	driver, drivErr := migratSqlit.WithInstance(db.dbConnection, &migratSqlit.Config{})
	if drivErr != nil {
		return nil, fmt.Errorf("Get database driver failed: %v", drivErr)
	}

	m, migrateErr := migrate.NewWithDatabaseInstance(migrationsPath, "sqlite3", driver)

	if migrateErr != nil {
		return nil, fmt.Errorf("Migration error: %v", migrateErr)
	}

	m.Up()

	if pingErr := db.dbConnection.Ping(); pingErr != nil {
		return nil, fmt.Errorf("Ping Error: %v", pingErr)
	}

	return db, nil
}

func (db *SqliteDB) CloseDB() error {
	err := db.dbConnection.Close()
	if err != nil {
		return fmt.Errorf("Error while closing database: %v", err)
	}
	return nil
}

func (db *SqliteDB) AddJournal(journal m.Journal) (int64, error) {
	result, err := db.dbConnection.Exec("INSERT INTO journals (title, date, content) VALUES (?, ?, ?)", journal.Title, journal.Date, journal.Content)
	if err != nil {
		return -1, fmt.Errorf("Add journal failed: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("AddJournal: LastInsertId error: %v", err)
	}

	return id, nil
}

func (db *SqliteDB) GetJournals() ([]m.Journal, error) {
	var journals []m.Journal

	rows, err := db.dbConnection.Query("SELECT * FROM journals ORDER BY date DESC")
	if err != nil {
		return nil, fmt.Errorf("Get Journals faild: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var j m.Journal
		if err := rows.Scan(&j.ID, &j.Title, &j.Date, &j.Content); err != nil {
			return nil, fmt.Errorf("GetJournals: Scan results error: %v", err)
		}

		journals = append(journals, j)
	}

	return journals, nil
}

func (db *SqliteDB) UpdateJournal(journal *m.Journal) (int64, error) {
	result, err := db.dbConnection.Exec("UPDATE journals SET title=?, date=?, content=? WHERE id=?",
		journal.Title, journal.Date, journal.Content, journal.ID)

	if err != nil {
		return -1, fmt.Errorf("UpdateJornal error: %v", err)
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("UpdateJornal error: %v", err)
	}

	return affect, nil
}

func (db *SqliteDB) DeleteJournal(id int64) (int64, error) {
	result, err := db.dbConnection.Exec("DELETE FROM journals WHERE id=?", id)
	if err != nil {
		return -1, fmt.Errorf("DeleteJournal error: %v", err)
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DeleteJournal error: %v", err)
	}

	return affect, nil
}
