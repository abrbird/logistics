package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddAddresses, downAddAddresses)
}

func upAddAddresses(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE logistics_addresses (
		    id bigint PRIMARY KEY,
			address VARCHAR NOT NULL 
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downAddAddresses(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE logistics_addresses;")
	if err != nil {
		return err
	}
	return nil
}
