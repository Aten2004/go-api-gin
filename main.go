package main

import (
	"github.com/Aten2004/go-api-gin/handlers" 
	"github.com/Aten2004/go-api-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("./students.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Student{})
	
	// ส่งค่า db ไปให้ handlers ใช้งาน
	handlers.SetDB(db)

	r := gin.Default()

	// เรียกใช้งานฟังก์ชันผ่าน handlers.
	r.GET("/students", handlers.GetStudents)
	r.POST("/students", handlers.CreateStudent)
	r.PUT("/students/:id", handlers.UpdateStudent)
	r.DELETE("/students/:id", handlers.DeleteStudent)

	r.Run(":8080")
}