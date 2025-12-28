package db

import (
	"database/sql"
	"fmt"
)

func GetFinanceStatement(db *sql.DB) (string, error) {
	rows, err := db.Query("SELECT date, nominal, status FROM finance")
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var context string
	for rows.Next() {
		var q, a, b string
		rows.Scan(&q, &a, &b)
		context += fmt.Sprintf("- %s %s : %s\n", q, a, b)
	}
	return context, nil
}
