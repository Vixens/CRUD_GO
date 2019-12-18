package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
	Debug    bool
}

func InitDB() *gorm.DB {
	conf := DBConfig{
		User:     "root",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Database: "crud_go",
		Debug:    true,
	}

	param := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ":" + conf.Port + ")/" +
		conf.Database + "?parseTime=true&loc=UTC&charset=utf8mb4"
	db, err := gorm.Open("mysql", param)

	if err != nil {
		panic(err)
	}

	if conf.Debug {
		db = db.Debug()
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	// db.AutoMigrate(&Stores)
	return db
}
