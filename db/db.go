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

var db *sql.DB

const migrationsPath = "file://db/db_migrations"

func InitDB() error {

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		fmt.Println("DB_PATH env varialbe not set. Defaulting to 'journals.db'")
		dbPath = "journals.db"
	}

	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("Oppining/creating database failed: %v", err)
	}

	driver, drivErr := migratSqlit.WithInstance(db, &migratSqlit.Config{})
	if drivErr != nil {
		return fmt.Errorf("Get database driver failed: %v", drivErr)
	}

	m, migrateErr := migrate.NewWithDatabaseInstance(migrationsPath, "sqlite3", driver)

	if migrateErr != nil {
		return fmt.Errorf("Migration error: %v", migrateErr)
	}

	m.Up()

	if pingErr := db.Ping(); pingErr != nil {
		return fmt.Errorf("Ping Error: %v", pingErr)
	}

	//TODO: remove test code when not needed anymore
	// for _, j := range m.TestJournals {
	// 	_, err := AddJournal(j)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func AddJournal(journal m.Journal) (int64, error) {
	result, err := db.Exec("INSERT INTO journals (title, date, content) VALUES (?, ?, ?)", journal.Title, journal.Date, journal.Content)
	if err != nil {
		return -1, fmt.Errorf("Add journal failed: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("AddJournal: LastInsertId error: %v", err)
	}

	return id, nil
}

func GetJournals() ([]m.Journal, error) {
	var journals []m.Journal

	rows, err := db.Query("SELECT * FROM journals ORDER BY date DESC")
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

func UpdateJournal(journal *m.Journal) (int64, error) {
	result, err := db.Exec("UPDATE journals SET title=? date=? content=? WHERE id=?",
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

func DeleteJournal(id int64) (int64, error) {
	result, err := db.Exec("DELETE FROM journals WHERE id=?", id)
	if err != nil {
		return -1, fmt.Errorf("DeleteJournal error: %v", err)
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("DeleteJournal error: %v", err)
	}

	return affect, nil
}
