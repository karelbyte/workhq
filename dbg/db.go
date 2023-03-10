package dbg

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"sync"
)


var lock = &sync.Mutex{}
var singleInstance *sql.DB
var err error

func DB() *sql.DB {

	dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")

    strConnection := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName
	
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance, err = sql.Open("mysql", strConnection)

		// defer singleInstance.Close()

		if err != nil {
			panic(err.Error())
		}
		fmt.Println("---------------------")
		fmt.Println("DB conection success!")

	}

	return singleInstance
}

func InitDB() {

}
