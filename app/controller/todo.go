package controller

import (
    "app/db"
    "app/model"
    "net/http"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)


func List(c *gin.Context) {
    db := db.Connection()
    defer db.Close()

    var todos []model.Todo
    db.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

func Create(c *gin.Context) {
    db := db.Connection()
    defer db.Close()

    todo := model.Todo{}
    err := c.BindJSON(&todo)
    if err != nil {
        c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
    }

    db.Create(&todo)
    if db.NewRecord(todo) == false {
        c.JSON(http.StatusCreated, todo)
    }
}

func Update(c *gin.Context) {
    db := db.Connection()
    defer db.Close()
    
    todo := model.Todo{}
    id := c.Param("id")
    db.Find(&todo, id)
    err := c.BindJSON(&todo)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
          "error": err.Error(),
        })
      }

    db.Save(&todo)
    c.JSON(http.StatusOK, todo)
}

func Delete(c *gin.Context) {
    db := db.Connection()
    defer db.Close()

    todo := model.Todo{}
    id := c.Param("id")
    db.First(&todo, id)
    err := db.First(&todo, id).Error
    if err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        return
    }
    db.Delete(&todo)
    c.JSON(http.StatusOK, gin.H{
        "status": "ok"})
}
