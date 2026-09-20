package db

import (
	"fmt"
	"log"
	"renet/config/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT string
	DBName string
	DBSSL  string
}

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
func InitDB() (db *gorm.DB, err error){

	cfg := DBConfig{
		DBUSER: env.GetString("DBUSER"),
		DBPASS: env.GetString("DBPASS"),
		DBHOST: env.GetString("DBHOST"),
		DBPORT: env.GetString("DBPORT"),
		DBName: env.GetString("DBName"),
		DBSSL: env.GetString("DBSSL"),
	}

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",cfg.DBHOST,cfg.DBUSER,cfg.DBPASS,cfg.DBName,cfg.DBPORT,cfg.DBSSL)

	DB,DBerr := gorm.Open(postgres.Open(dsn),&gorm.Config{})

	if DBerr != nil{
		return nil,fmt.Errorf("failed to connect to database: %w",DBerr)
	}
	
	pgsql,err := DB.DB()
	if err!=nil{
		return nil,err
	}

	err = pgsql.Ping()
	if err != nil{
		log.Println("Failed to connected with the DB")
		return nil,err
	}
	log.Println("Successfully connected with the DB")

	return DB,nil

}