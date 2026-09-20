package db

import "renet/config/env"

type DBConfig struct {
	DBUSER string
	DBPASS string
	DBHOST string
	DBPORT int
	DBName string
	DBSSL  string
}

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
func InitDB() {

	cfg := DBConfig{
		DBUSER: env.GetString("DBUSER"),
		DBPASS: env.GetString("DBPASS"),
		DBHOST: env.GetString("DBHOST"),
		DBPORT: env.GetInt("DBPORT"),
		DBName: env.GetString("DBName"),
		DBSSL: env.GetString("DBSSL"),
	}
	

}