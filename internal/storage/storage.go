package storage

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Balance struct {
	Name    string          `json:"name"`
	Surname string          `json:"surname"`
	Credit  sql.NullFloat64 `json:"credit"`
}

type BalanceLog struct {
	Name    string    `json:"Name"`
	Surname string    `json:"Surname"`
	Credit  int       `json:"Credit"`
	UTCTime time.Time `json:"UTCTime"`
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

func StoreBalanceLogs() {
	allBalances, err := db.Query("SELECT name, surname, credit FROM ledgerappuserdata")
	if err != nil {
		panic(err.Error())
	}
	defer allBalances.Close()

	currentTime := time.Now().UTC()

	for allBalances.Next() {
		var name, surname string
		var credit int

		err := allBalances.Scan(&name, &surname, &credit)
		if err != nil {
			panic(err.Error())
		}

		insertQuery := "INSERT INTO balance_logs (name, surname, credit, timestamp_utc) VALUES (?, ?, ?, ?)"
		_, err = db.Exec(insertQuery, name, surname, credit, currentTime)
		if err != nil {
			panic(err.Error())
		}
	}

	if err = allBalances.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println("store balance log!")
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

	StoreBalanceLogs()

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

func UserBalance(name string, surname string, token string) ([]byte, error) {
	userBalance, err := db.Prepare(
		"SELECT name, surname, credit FROM ledgerappuserdata WHERE name = ? AND surname = ? AND Token = ?")
	if err != nil {
		return nil, err
	}
	defer userBalance.Close()

	rows, err := userBalance.Query(name, surname, token)
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

func userExists(name string, surname string, token string) bool {
	query := "SELECT COUNT(*) > 0 AS userExists FROM ledgerappuserdata WHERE name = ? AND surname = ? AND Token = ?"
	row := db.QueryRow(query, name, surname, token)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}

func TransferCredit(
	SenderName string,
	SenderSurname string,
	SenderToken string,
	ReceiverName string,
	ReceiverSurname string,
	TransferAmount int) {
	addToReceiver, err := db.Prepare("UPDATE ledgerappuserdata SET credit = IFNULL(credit, 0) + ? WHERE name = ? AND surname = ?")
	removeFromSender, err := db.Prepare("UPDATE ledgerappuserdata SET credit = IFNULL(credit, 0) - ? WHERE name = ? AND surname = ?")

	if userExists(SenderName, SenderSurname, SenderToken) {
		if err != nil {
			panic(err.Error())
		}
		defer addToReceiver.Close()
		defer removeFromSender.Close()

		_, err = addToReceiver.Exec(TransferAmount, ReceiverName, ReceiverSurname)
		if err != nil {
			panic(err.Error())
		}

		_, err = removeFromSender.Exec(TransferAmount, SenderName, SenderSurname)
		if err != nil {
			panic(err.Error())
		}

		StoreBalanceLogs()

		fmt.Println("Transfer done")
	}
}

func GetBalanceLog(Name string, Surname string, startDate time.Time, endDate time.Time) ([]byte, error) {
	query := "SELECT name, surname, credit, timestamp_utc FROM balance_logs WHERE name = ? AND surname = ? AND timestamp_utc >= ? AND timestamp_utc <= ?"

	rows, err := db.Query(query, Name, Surname, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []BalanceLog

	for rows.Next() {
		var log BalanceLog

		err := rows.Scan(&log.Name, &log.Surname, &log.Credit, &log.UTCTime)
		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(logs)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
