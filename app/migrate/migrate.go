package main

import (
    "app/db"
    "app/model"

    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    db := db.Connection()
    defer db.Close()

    db.AutoMigrate(&model.User{})
    db.AutoMigrate(&model.Todo{})
}