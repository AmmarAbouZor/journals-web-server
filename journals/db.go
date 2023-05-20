package journals

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	migratSqlit "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var db *sql.DB

// TODO: get db path from environment variable
const dbFile = "journals.db"
const migrationsPath = "file://journals/db_migrations"

func InitDB() error {

	var err error
	db, err = sql.Open("sqlite3", dbFile)
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

	return nil
}
