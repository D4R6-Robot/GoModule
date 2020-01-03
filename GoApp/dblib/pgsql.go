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

// InitDbInfo struct
func (dbInfo *PostSQLDB) InitDbInfo() {
	dbInfo.host = "192.168.220.104"
	dbInfo.port = 5432
	dbInfo.user = "ctadmin"
	dbInfo.name = "D4R6"
	dbInfo.password = "ctadmin123$%^"
}

// https://golangkorea.github.io/post/go-start/object-oriented/

// ConnectDB struct
func (dbInfo *PostSQLDB) ConnectDB() {

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

// SelectDB struct
func (dbInfo *PostSQLDB) SelectDB() {
	rows, err := dbInfo.db.Query("SELECT \"MAL_INDEX\" FROM public.\"APT_DECISION_TB\"")

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer rows.Close()

	var malIndex string
	for rows.Next() {
		err := rows.Scan(&malIndex)
		if err != nil {
			fmt.Print(err)
			panic(err)
		}

		fmt.Println(malIndex)
	}
}

// CloseDB struct
func (dbInfo *PostSQLDB) CloseDB() {
	dbInfo.db.Close()
}
