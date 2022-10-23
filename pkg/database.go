package pkg

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	usr  = os.Getenv("DB_USERNAME")
	pas  = os.Getenv("DB_PASSWORD")
	url  = os.Getenv("DB_URL")
	port = os.Getenv("DB_PORT")
	name = os.Getenv("DB_NAME")
)

func InitDB() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=" + url + " user=" + usr + " password=" + pas + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//db.AutoMigrate(&models.Book{})

	return db, err
}
