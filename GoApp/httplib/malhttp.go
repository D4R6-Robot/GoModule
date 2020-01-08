package httplib

import (
	"fmt"
	"malwaresystem/dblib"
	"malwaresystem/mallist"
	"net/http"
)

// ListHandler struct
func ListHandler(w http.ResponseWriter, r *http.Request) {

	cDBInfo := dblib.PostSQLDB{}
	cDBInfo.MalDBInfo()
	cDBInfo.MalDBConnect()

	var malData []mallist.MalwareInfo
	malData = cDBInfo.MalDBSelect()

	cDBInfo.CloseDB()

	for index, element := range malData {
		fmt.Fprintln(w, index, "\t", element.GetMalIndex(), "\t", element.GetMalName(), "\t", element.GetY(), "\t", element.GetProcTree(), "\t",
			element.GetProcXSLCnt(), "\t", element.GetProcJSCnt(), "\t", element.GetProcMSFormsCnt(), "\t", element.GetProcMacrosCnt())
	}
}
