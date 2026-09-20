package db

import (
	"fmt"
	"log"
	"renet/config/env"
	"renet/schemas"

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
func InitDB() (db *gorm.DB, err error) {

	cfg := DBConfig{
		DBUSER: env.GetString("DBUSER"),
		DBPASS: env.GetString("DBPASS"),
		DBHOST: env.GetString("DBHOST"),
		DBPORT: env.GetString("DBPORT"),
		DBName: env.GetString("DBName"),
		DBSSL:  env.GetString("DBSSL"),
	}

	if cfg.DBUSER == "" {
		return nil, fmt.Errorf("Error occured while parsing the USER cnfgs for DB ")
	}
	if cfg.DBHOST == "" {
		return nil, fmt.Errorf("Error occured while parsing the HOST cnfgs for DB ")
	}
	if cfg.DBName == "" {
		return nil, fmt.Errorf("Error occured while parsing the NAME cnfgs for DB ")
	}
	if cfg.DBPASS == "" {
		return nil, fmt.Errorf("Error occured while parsing the PASS cnfgs for DB ")
	}
	if cfg.DBPORT == "" {
		return nil, fmt.Errorf("Error occured while parsing the PORT cnfgs for DB ")
	}
	if cfg.DBSSL == "" {
		return nil, fmt.Errorf("Error occured while parsing the SSL cnfgs for DB ")
	}

	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DBHOST, cfg.DBUSER, cfg.DBPASS, cfg.DBName, cfg.DBPORT, cfg.DBSSL)

	DB, DBerr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if DBerr != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", DBerr)
	}

	pgsql, err := DB.DB()
	if err != nil {
		return nil, err
	}

	err = pgsql.Ping()
	if err != nil {
		log.Println("Failed to connected with the DB")
		return nil, err
	}

	if err := DB.AutoMigrate(&schemas.User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate user table: %w", err)
	}
	log.Println("Successfully connected with the DB")

	return DB, nil

}
