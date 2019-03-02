package models

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"vote/utils/setting"
)

var db *sql.DB

func Setup() {
	var (
		err                          error
		dbName, user, password, host string
		port                         int
	)

	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	port = setting.DatabaseSetting.Port
	connString := fmt.Sprintf("server=%s;port%d;database=%s;user id=%s;password=%s",
		host, port, dbName, user, password)
	db, err = sql.Open("mssql", connString)

	if err != nil {
		log.Println(err)
	}

}

func CloseDB() {
	defer db.Close()
}
