package main

import (
	"fmt"
        "github.com/jinzhu/gorm"
        _ "github.com/jinzhu/gorm/dialects/mysql"
)

const DRIVER = "mysql"
const DSN = "root:a1b2V4d@tcp(mysql-container:3306)/gin_app?charset=utf8&parseTime=True&loc=Local"

func main() {
	db, err := gorm.Open(DRIVER, DSN)
	if err != nil {
                fmt.Println(err)
		fmt.Println("Open fail")
	} else {
		fmt.Println("OpenOK")
	}


	db.Close()
}
