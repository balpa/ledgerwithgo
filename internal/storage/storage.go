package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Balance struct {
	Name    string          `json:"name"`
	Surname string          `json:"surname"`
	Credit  sql.NullFloat64 `json:"credit"`
}

var db *sql.DB

func createUUID() uuid.UUID {
	return uuid.New()
}

func ConnectDB() error {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/ledgerapp")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to MySQL ledgerapp db!")
	return nil
}

func CreateUniqueUser(name string, surname string) {
	uuid := createUUID()

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

func AddCredit(name string, surname string, amount int) {
	addCredit, err := db.Prepare("UPDATE ledgerappuserdata SET credit = IFNULL(credit, 0) + ? WHERE name = ? AND surname = ?")
	if err != nil {
		panic(err.Error())
	}
	defer addCredit.Close()

	_, err = addCredit.Exec(amount, name, surname)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Credit amount changed!")
}

func GetAllBalances() []byte {
	allBalances, err := db.Prepare("SELECT name, surname, credit FROM ledgerappuserdata")
	if err != nil {
		panic(err.Error())
	}
	defer allBalances.Close()

	rows, err := allBalances.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var balances []Balance

	for rows.Next() {
		var balance Balance
		if err := rows.Scan(&balance.Name, &balance.Surname, &balance.Credit); err != nil {
			panic(err.Error())
		}
		balances = append(balances, balance)
	}

	if err := rows.Err(); err != nil {
		panic(err.Error())
	}

	jsonData, err := json.Marshal(balances)
	if err != nil {
		panic(err.Error())
	}

	return jsonData
}

func UserBalance(name string, surname string) ([]byte, error) {
	userBalance, err := db.Prepare("SELECT name, surname, credit FROM ledgerappuserdata WHERE name = ? AND surname = ?")
	if err != nil {
		return nil, err
	}
	defer userBalance.Close()

	rows, err := userBalance.Query(name, surname)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var balances []Balance

	for rows.Next() {
		var balance Balance
		if err := rows.Scan(&balance.Name, &balance.Surname, &balance.Credit); err != nil {
			return nil, err
		}
		balances = append(balances, balance)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(balances)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
