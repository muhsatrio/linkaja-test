package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	DBName   string `yaml:"db_name"`
}

// Open is to open a connection to mysql server
func Open(cf Config) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cf.Username, cf.Password, cf.Host, cf.Port, cf.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Customer{}, &Account{})

	// init data

	customers := []Customer{
		{
			CustomerNumber: 1001,
			Name:           "Bob Martion",
		},
		{
			CustomerNumber: 1002,
			Name:           "Linus Torvalds",
		},
	}

	accounts := []Account{
		{
			AccountNumber:  555001,
			CustomerNumber: 1001,
			Balance:        10000,
		},
		{
			AccountNumber:  555002,
			CustomerNumber: 1002,
			Balance:        15000,
		},
	}

	for _, eachData := range customers {
		db.Where(Customer{CustomerNumber: eachData.CustomerNumber}).FirstOrCreate(&eachData)
	}

	for _, eachData := range accounts {
		db.Where(Account{AccountNumber: eachData.AccountNumber}).FirstOrCreate(&eachData)
	}

	return db, nil
}
