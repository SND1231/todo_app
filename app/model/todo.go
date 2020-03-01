package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Todo struct {
    gorm.Model
    Content string
    Status  int
}