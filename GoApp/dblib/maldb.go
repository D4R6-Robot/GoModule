package dblib

import (
	"database/sql"
	"fmt"

	"malwaresystem/mallist"

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

// MalDBSelect struct
func (dbInfo *PostSQLDB) MalDBSelect() []mallist.MalwareInfo {

	var malIndex string
	var malName string
	var y int
	var procTree int
	var procXSLCnt int
	var procJSCnt int
	var procMSFormsCnt int
	var procMacrosCnt int
	var netHTTPCnt int
	var ntQueryKey int
	var ntOpenKey int
	var ntOpenFile int
	var ntWriteFile int
	var ntQueryValueKey int
	var ntDelayExecution int
	var ntCreateFile int
	var fileCreated int
	var fileReCreated int
	var fileDllLoaded int
	var fileOpened int
	var fileWritten int
	var fileFailed int

	rows, err := dbInfo.db.Query("SELECT * FROM public.\"APT_DECISION_TB\"")

	var count int
	dbInfo.db.QueryRow("SELECT count(*) FROM public.\"APT_DECISION_TB\"").Scan(&count)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	malData := make([]mallist.MalwareInfo, count, count)
	defer rows.Close()

	i := 0

	for rows.Next() {

		err := rows.Scan(&malIndex,
			&malName,
			&y,
			&procTree,
			&procXSLCnt,
			&procJSCnt,
			&procMSFormsCnt,
			&procMacrosCnt,
			&netHTTPCnt,
			&ntQueryKey,
			&ntOpenKey,
			&ntOpenFile,
			&ntWriteFile,
			&ntQueryValueKey,
			&ntDelayExecution,
			&ntCreateFile,
			&fileCreated,
			&fileReCreated,
			&fileDllLoaded,
			&fileOpened,
			&fileWritten,
			&fileFailed)

		if err != nil {
			fmt.Print(err)
			panic(err)
		}

		malData[i].SetMalIndex(malIndex)
		malData[i].SetMalName(malName)
		malData[i].SetY(y)
		malData[i].SetProcTree(procTree)
		malData[i].SetProcXSLCnt(procXSLCnt)
		malData[i].SetProcJSCnt(procJSCnt)
		malData[i].SetProcMSFormsCnt(procMSFormsCnt)
		malData[i].SetProcMacrosCnt(procMacrosCnt)
		malData[i].SetNetHTTPCnt(netHTTPCnt)
		malData[i].SetNtQueryKey(ntQueryValueKey)
		malData[i].SetNtOpenKey(ntOpenKey)
		malData[i].SetNtOpenFile(ntOpenFile)
		malData[i].SetNtWriteFile(ntWriteFile)
		malData[i].SetNtQueryValueKey(ntQueryValueKey)
		malData[i].SetNtDelayExecution(ntDelayExecution)
		malData[i].SetNtCreateFile(ntCreateFile)
		malData[i].SetFileCreated(fileCreated)
		malData[i].SetFileReCreated(fileReCreated)
		malData[i].SetFileDllLoaded(fileDllLoaded)
		malData[i].SetFileOpened(fileOpened)
		malData[i].SetFileWritten(fileWritten)
		malData[i].SetFileFailed(fileFailed)

		i = i + 1
	}

	return malData
}

// CloseDB struct
func (dbInfo *PostSQLDB) CloseDB() {
	dbInfo.db.Close()
}
