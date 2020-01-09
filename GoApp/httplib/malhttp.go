package httplib

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"malwaresystem/dblib"
	"net/http"
	"os"
)

// MalwareInfo struct
// 일단  struct 의 변수가 대문자가 아니면 template 에서 먹히지가 않는다.
// 이 data struct 를 따로 빼서 관리하려고 하니, 실제 요소에 접근이 안되서 데려왔는데, 사실
//
// func NewMalwareInfo()(*MalwareInfo){
//    return &( MalwareInfo{} );
// 이렇게 데리고 와도, 요소에 접근이 안된다. 캡슐화가 잘 되어 있는 것 같다. 이 분리를 어떻게 시킬지는 계속 업데이트
type MalwareInfo struct {
	MalIndex         string
	MalName          string
	Y                int
	ProcTree         int
	ProcXSLCnt       int
	ProcJSCnt        int
	ProcMSFormsCnt   int
	ProcMacrosCnt    int
	NetHTTPCnt       int
	NtQueryKey       int
	NtOpenKey        int
	NtOpenFile       int
	NtWriteFile      int
	NtQueryValueKey  int
	NtDelayExecution int
	NtCreateFile     int
	FileCreated      int
	FileReCreated    int
	FileDllLoaded    int
	FileOpened       int
	FileWritten      int
	FileFailed       int
}

// ListHandler struct
func ListHandler(w http.ResponseWriter, r *http.Request) {

	cDBInfo := dblib.PostSQLDB{}
	cDBInfo.MalDBInfo()
	cDBInfo.MalDBConnect()

	var rows *sql.Rows
	var count int

	rows = cDBInfo.MalDBSelect()
	count = cDBInfo.MalDBCount()

	var MalIndex string
	var MalName string
	var Y int
	var ProcTree int
	var ProcXSLCnt int
	var ProcJSCnt int
	var ProcMSFormsCnt int
	var ProcMacrosCnt int
	var NetHTTPCnt int
	var NtQueryKey int
	var NtOpenKey int
	var NtOpenFile int
	var NtWriteFile int
	var NtQueryValueKey int
	var NtDelayExecution int
	var NtCreateFile int
	var FileCreated int
	var FileReCreated int
	var FileDllLoaded int
	var FileOpened int
	var FileWritten int
	var FileFailed int

	malData := make([]MalwareInfo, count, count)
	defer rows.Close()

	i := 0
	for rows.Next() {

		err := rows.Scan(&MalIndex,
			&MalName,
			&Y,
			&ProcTree,
			&ProcXSLCnt,
			&ProcJSCnt,
			&ProcMSFormsCnt,
			&ProcMacrosCnt,
			&NetHTTPCnt,
			&NtQueryKey,
			&NtOpenKey,
			&NtOpenFile,
			&NtWriteFile,
			&NtQueryValueKey,
			&NtDelayExecution,
			&NtCreateFile,
			&FileCreated,
			&FileReCreated,
			&FileDllLoaded,
			&FileOpened,
			&FileWritten,
			&FileFailed)

		if err != nil {
			fmt.Print(err)
			panic(err)
		}

		malData[i].MalIndex = MalIndex
		malData[i].MalName = MalName
		malData[i].Y = Y
		malData[i].ProcTree = ProcTree
		malData[i].ProcXSLCnt = ProcXSLCnt
		malData[i].ProcJSCnt = ProcJSCnt
		malData[i].ProcMSFormsCnt = ProcMSFormsCnt
		malData[i].ProcMacrosCnt = ProcMacrosCnt
		malData[i].NetHTTPCnt = NetHTTPCnt
		malData[i].NtQueryKey = NtQueryKey
		malData[i].NtOpenKey = NtOpenKey
		malData[i].NtOpenFile = NtOpenFile
		malData[i].NtWriteFile = NtWriteFile
		malData[i].NtQueryValueKey = NtQueryValueKey
		malData[i].NtDelayExecution = NtDelayExecution
		malData[i].NtCreateFile = NtCreateFile
		malData[i].FileCreated = FileCreated
		malData[i].FileReCreated = FileReCreated
		malData[i].FileDllLoaded = FileDllLoaded
		malData[i].FileOpened = FileOpened
		malData[i].FileWritten = FileWritten
		malData[i].FileFailed = FileFailed

		i = i + 1
	}

	cDBInfo.CloseDB()

	if err := os.Chdir("C:/Users/D4R6/Go/src"); err != nil {
		panic(err)
	}

	// Working Directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var templates = template.Must(template.ParseFiles(wd + "/malwaresystem/templates/mallist.html"))

	templates.Execute(w, malData)
}
