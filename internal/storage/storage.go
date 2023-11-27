package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ledgerapp")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL ledgerapp db!")

	return db, nil
}

func QueryDatabase() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ledgerappuserdata")
	if err != nil {
		fmt.Println("Error querying database:", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var col1, col2 string
		if err := rows.Scan(&col1, &col2); err != nil {
			fmt.Println("Error scanning rows:", err)
			return
		}

		fmt.Println(col1, col2)
	}
}

func CreateUniqueUser(name string, surname string) {
	uuid := createUUID()

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ledgerapp")
	if err != nil {
		panic(err.Error())
	}

	insertUniqueUser, err := db.Prepare("INSERT IGNORE INTO ledgerappuserdata (uuid, name, surname) VALUES (?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer insertUniqueUser.Close()

	_, err = insertUniqueUser.Exec(uuid, name, surname)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Unique user created.")
}

func createUUID() uuid.UUID {
	return uuid.New()
}
