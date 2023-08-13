package main

import (
	"go-server-part3/internal/routes"

	"github.com/labstack/echo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	// Khởi tạo kết nối cơ sở dữ liệu MySQL
	dsn := "root:root@tcp(127.0.0.1:3307)/todo_list?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Khởi tạo router và đăng ký các route
	routes.RegisterMainRoutes(e, db)

	// Khởi chạy server
	e.Logger.Fatal(e.Start(":8080"))
}
