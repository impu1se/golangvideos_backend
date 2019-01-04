package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20190104070548, Down20190104070548)
}

func Up20190104070548(tx *sql.Tx) error {
	_, err := tx.Exec(`
        create table if not exists videos (
            id serial primary key,
            name varchar(255) not null,
            created_at timestamp default NOW(),
            updated_at timestamp default NOW()
        )
    `)
	if err != nil {
		return err
	}
	// This code is executed when the migration is applied.
	return nil
}

func Down20190104070548(tx *sql.Tx) error {
	_, err := tx.Exec(`drop table if exists videos`)
	if err != nil {
		return err
	}
	return nil
}
