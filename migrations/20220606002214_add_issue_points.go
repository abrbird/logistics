package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddIssuePoints, downAddIssuePoints)
}

func upAddIssuePoints(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE logistics_issue_points (
		    id bigint PRIMARY KEY,
			address_id bigint NOT NULL UNIQUE ,
			is_available boolean NOT NULL,
			CONSTRAINT logistics_issue_points_fk_logistics_addresses
			    FOREIGN KEY(address_id)
			        REFERENCES logistics_addresses(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downAddIssuePoints(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE logistics_issue_points;")
	if err != nil {
		return err
	}
	return nil
}
