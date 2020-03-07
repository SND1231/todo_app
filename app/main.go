package main

import (
    "app/controller"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    r := gin.Default()
    r.Use(cors.New(cors.Config{
        // 許可したいHTTPメソッドの一覧
        AllowMethods: []string{
            "POST",
            "GET",
            "OPTIONS",
            "PUT",
            "DELETE",
        },
        // 許可したいHTTPリクエストヘッダの一覧
        AllowHeaders: []string{
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "X-CSRF-Token",
            "Authorization",
        },
        AllowOrigins: []string{
            "*",
        },
    }))
    r.GET("/todo", controller.List)
    r.POST("/todo", controller.Create)
    r.PUT("/todo/:id", controller.Update)
    r.DELETE("/todo/:id", controller.Delete)

    r.Run()
}
