package config

import (
	"log"
	"os"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
				)

				// DB
				// ده المتغير اللي هنستخدمه في باقي المشروع عشان نكلم الداتابيز
				// لازم يكون بحروف كبيرة عشان يتشاف بره الفولدر ده
				var DB *gorm.DB

				func ConnectDatabase() {
					// بنقرا الرابط من ملف الـ .env
						// لاحظ اننا بنفترض ان ملف الـ env اتحمل خلاص في الـ main
							dbPath := os.Getenv("DB_URL")

								var err error
									// الاتصال وتخزين النتيجة في المتغير DB اللي فوق
										DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

											if err != nil {
													log.Fatal("فشل الاتصال بقاعدة البيانات: ", err)
														}

															log.Println("تم الاتصال بقاعدة البيانات بنجاح")
															}
															