package httplib

import (
	"html/template"
	"log"
	"malwaresystem/dblib"
	"malwaresystem/mallist"
	"net/http"
	"os"
)

// ListHandler struct
func ListHandler(w http.ResponseWriter, r *http.Request) {

	cDBInfo := dblib.PostSQLDB{}
	cDBInfo.MalDBInfo()
	cDBInfo.MalDBConnect()

	var malData []mallist.MalwareInfo
	malData = cDBInfo.MalDBSelect()

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

	/*
		for index, element := range malData {
			fmt.Fprintln(w, index, "\t", element.GetMalIndex(), "\t", element.GetMalName(), "\t", element.GetY(), "\t", element.GetProcTree(), "\t",
				element.GetProcXSLCnt(), "\t", element.GetProcJSCnt(), "\t", element.GetProcMSFormsCnt(), "\t", element.GetProcMacrosCnt())
		}
	*/
}
