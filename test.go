package main

import (
	"fmt"
	"log"

	// استدعاء الفولدر الجديد
	// "اسم-الموديول/اسم-الفولدر"
	"Banking-System-Golang/config"
	"Banking-System-Golang/models" 
	"github.com/joho/godotenv"
)


func main() {
	// 1. تحميل ملف .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. استدعاء دالة الاتصال من الفولدر الجديد
	config.ConnectDatabase()

	// 3. استخدام الداتابيز
	// config.DB لاحظ اننا بنستخدم
	config.DB.AutoMigrate(&models.Product{})

	// تجربة إضافة منتج
	config.DB.Create(&models.Product{Code: "Laptop cor i 3", Price: 15000})
	fmt.Println("تمت العملية بنجاح!")
}