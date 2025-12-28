package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	// fmt.Printf(dsn)
	// return sql.Open("mysql", dsn)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 🔴 PENTING: cek koneksi nyata
	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("✅ MySQL connected successfully")
	return db, nil
}
