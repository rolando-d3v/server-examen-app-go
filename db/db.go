package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func ConexionBD() (db *sql.DB) {

	godotenv.Load()

	var user = os.Getenv("DBUSER")
	var pass = os.Getenv("DBPASS")
	var dbhost = os.Getenv("DBHOST")
	var dbPort = os.Getenv("DBPORT")
	var dbName = os.Getenv("DBNAME")

	url := string(user + ":" + pass + "@tcp(" + dbhost + ":" + dbPort + ")/" + dbName)
	// var url string = user + ":" + pass + "@tcp(" + dbhost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", url)
	// db, err := sql.Open("mysql", "root:fXaYofuS1RXXXD8R0FlP@tcp(containers-us-west-105.railway.app:6565)/railway")
	// db, err := sql.Open("mysql", "root:Rolando@tcp(localhost)/db_personal")
	if err != nil {
		fmt.Println("Error DB ", err.Error())
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println("Connection x success DB 😃  ✔️")
	return db

}
