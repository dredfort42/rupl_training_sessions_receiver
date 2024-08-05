package db

import (
	"database/sql"

	loger "github.com/dredfort42/tools/logprinter"
)

// TrainingResultExistsCheck checks if a device exists in the database based on the email and device ID provided
func TrainingResultExistsCheck(sessionStartTime int64, sessionEndTime int64, email string) bool {
	query := `
		SELECT session_start_time, session_end_time, email 
		FROM ` + DB.TableTrainingResults + ` 
		WHERE session_start_time = $1 AND session_end_time = $2 AND email = $3;`

	var existingStartTime int64
	var existingEndTime int64
	var existingEmail string

	err := DB.Database.QueryRow(query, sessionStartTime, sessionEndTime, email).Scan(&existingStartTime, &existingEndTime, &existingEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}

		loger.Error("Error checking if device exists", err)
		return false
	}

	return true
}
