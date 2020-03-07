package main

import (
    "app/db"
    "app/model"
    "fmt"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    db := db.Connection()
    defer db.Close()
    fmt.Printf("migration")
    
    db.AutoMigrate(&model.User{})
    db.AutoMigrate(&model.Todo{})
}
