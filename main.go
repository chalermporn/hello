package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "Chalermporn"})

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		var u User
		db.First(&u)
		c.JSON(200, u)
	})

	r.Run()

}
