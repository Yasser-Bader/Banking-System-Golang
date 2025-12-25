package main

import (
	"fmt"
	"log"

			// استدعاء الفولدر الجديد
				// "اسم-الموديول/اسم-الفولدر"
	"Banking-System-Golang/config" 

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	)

							// المودل ممكن تخليه هنا أو تنقله فولدر models بعدين
							type Product struct {
								gorm.Model
									Code  string
										Price uint
										}

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
																					config.DB.AutoMigrate(&Product{})

																						// تجربة إضافة منتج
																							config.DB.Create(&Product{Code: "Laptop", Price: 15000})
																								fmt.Println("تمت العملية بنجاح!")
																								}
