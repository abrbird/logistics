package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddOrdersAvailability, downAddOrdersAvailability)
}

func upAddOrdersAvailability(tx *sql.Tx) error {
	_, err := tx.Exec(`
		CREATE TABLE logistics_orders_availability (
		    order_id bigint NOT NULL,
			issue_point_id bigint NOT NULL,
			status VARCHAR NOT NULL,
			updated_at TIMESTAMP NOT NULL default current_timestamp,
			PRIMARY KEY (order_id, issue_point_id),
			CONSTRAINT logistics_orders_availability_fk_logistics_issue_points
			    FOREIGN KEY(issue_point_id)
			        REFERENCES logistics_issue_points(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downAddOrdersAvailability(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE logistics_orders_availability;")
	if err != nil {
		return err
	}
	return nil
}
