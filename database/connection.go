package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"multi-messenger-server/config"
	"multi-messenger-server/tools"
)

var DB *gorm.DB

func ConnectDB() error {
	dsnTemplate := "host={{.host}} user={{.user}} password={{.password}} dbname={{.database}} port={{.port}}"
	data := map[string]interface{}{
		"host":     config.DbConfig.Host,
		"user":     config.DbConfig.User,
		"password": config.DbConfig.Password,
		"database": config.DbConfig.Database,
		"port":     config.DbConfig.Port,
	}
	dsn, err := tools.GetTextFromTemplate(dsnTemplate, data)
	if err != nil {
		return err
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}
