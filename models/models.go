package models

import "database/sql"

var (
	db *sql.DB = nil
)

// Globally sets the database connection for the models package
func UseDB(DBConnection *sql.DB) {
	db = DBConnection
}
