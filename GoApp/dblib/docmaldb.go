package dblib

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // here
)

// PostSQLDB struct
type PostSQLDB struct {
	user     string
	password string
	name     string
	host     string
	port     int
	db       *sql.DB
}

// MalDBInfo struct
func (dbInfo *PostSQLDB) MalDBInfo() {
	dbInfo.host = "192.168.220.104"
	dbInfo.port = 5432
	dbInfo.user = "ctadmin"
	dbInfo.name = "D4R6"
	dbInfo.password = "ctadmin123$%^"
}

// https://golangkorea.github.io/post/go-start/object-oriented/

// MalDBConnect struct
func (dbInfo *PostSQLDB) MalDBConnect() {

	dbInfoCon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbInfo.host, dbInfo.port, dbInfo.user, dbInfo.password, dbInfo.name)
	dbcon, err := sql.Open("postgres", dbInfoCon)

	if err != nil {
		panic(err)
	}

	err = dbcon.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	dbInfo.db = dbcon
}

// MalDBCount struct
func (dbInfo *PostSQLDB) MalDBCount() int {
	var count int
	dbInfo.db.QueryRow("SELECT count(*) FROM public.\"APT_DECISION_TB\"").Scan(&count)

	return count
}

// MalDBSelect struct
func (dbInfo *PostSQLDB) MalDBSelect() *sql.Rows {

	rows, err := dbInfo.db.Query("SELECT * FROM public.\"APT_DECISION_TB\"")

	var count int
	dbInfo.db.QueryRow("SELECT count(*) FROM public.\"APT_DECISION_TB\"").Scan(&count)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return rows
}

// CloseDB struct
func (dbInfo *PostSQLDB) CloseDB() {
	dbInfo.db.Close()
}
