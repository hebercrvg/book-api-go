package database

import (
	"books/infra/database/migrations"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	config := BuildDatabaseConfig()
	// str := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v",
	// 	config.Host, config.Port, config.User, config.Password, config.Database)
	str := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database)
	fmt.Printf(str)
	// str = "host=localhost port=5432 user=postgres password=thugstools dbname=my_books"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect in database: " + err.Error())
	}

	db = database

	migrations.RunMigrations(db)

}

func GetDatabase() *gorm.DB {
	return db
}
