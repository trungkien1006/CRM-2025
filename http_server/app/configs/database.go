package configs

import (
	"admin-v1/app/helpers"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func GormConnection() *gorm.DB {
	var db *gorm.DB

	if helpers.GormDB != nil {
		return helpers.GormDB
	}

	var (
		devHostName = os.Getenv("MYSQL_HOST")
		devDbName   = os.Getenv("MYSQL_NAME")
		devUser     = os.Getenv("MYSQL_USER")
		devPassword = os.Getenv("MYSQL_PASSWORD")
		devPort     = os.Getenv("MYSQL_PORT")
	)

	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		devUser,
		devPassword,
		devHostName,
		devPort,
		devDbName,
	)

	// Retry tối đa 10 lần, mỗi lần cách nhau 5s
	for i := 0; i < 10; i++ {
		db, err := gorm.Open(
			mysql.Open(dsn),
			&gorm.Config{
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true,
				},
			},
		)
		if err == nil {
			fmt.Println("Kết nối MySQL thành công!")
			return db
		}
		fmt.Printf("Không thể kết nối MySQL, thử lại sau 5s (%d/10)...\n", i+1)
		fmt.Println("DSN:", dsn)
		time.Sleep(5 * time.Second)
	}
	// db, err := gorm.Open(
	// 	mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", devUser, devPassword, devHostName, devPort, devDbName)+"?parseTime=true&charset=utf8mb4&loc=Local"),
	// 	&gorm.Config{
	// 		NamingStrategy: schema.NamingStrategy{
	// 			SingularTable: true,
	// 		},
	// 	},
	// )

	if err != nil {
		panic(err)
	}

	return db
}