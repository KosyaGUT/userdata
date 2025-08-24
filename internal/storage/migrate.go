package storage

import "github.com/jmoiron/sqlx"

func InitSchema(db *sqlx.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY,
        login TEXT NOT NULL,
        fcs TEXT NOT NULL,
        sex TEXT NOT NULL,
        age SMALLINT NOT NULL,
        contacts TEXT[],
        avatar TEXT,
        date_reg TIMESTAMP,
        status BOOLEAN
    );`
	_, err := db.Exec(schema)
	return err
}
