package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

type Tweet struct {
    ID      uint   `gorm:"primaryKey"`
    Content string `json:"content"`
}

var DB *gorm.DB

func initDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    DB.AutoMigrate(&Tweet{})
}

func main() {
    r := gin.Default()
    initDatabase()

    r.GET("/tweets", func(c *gin.Context) {
        var tweets []Tweet
        DB.Find(&tweets)
        c.JSON(200, tweets)
    })

    r.POST("/tweets", func(c *gin.Context) {
        var tweet Tweet
        if err := c.ShouldBindJSON(&tweet); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        DB.Create(&tweet)
        c.JSON(200, tweet)
    })

    r.Run(":8080")
}

