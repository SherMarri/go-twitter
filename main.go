package main

import (
	"log"
	"net/http"

	"./common"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db = common.Dependencies.DB

func main() {
	r := GetProjectRouter()
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("Server failed.")
	}
}
