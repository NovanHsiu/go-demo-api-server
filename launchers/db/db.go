package db

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/NovanHsiu/go-demo-api-server/models"
	"github.com/NovanHsiu/go-demo-api-server/utils"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(dbConfig map[string]string) (*gorm.DB, error) {
	host := dbConfig["host"]
	dbname := dbConfig["dbname"]
	port := dbConfig["port"]
	sslmode := dbConfig["sslmode"]
	user := dbConfig["user"]
	dbType := strings.ToLower(dbConfig["type"])
	password := dbConfig["passwd"]
	timeZone := "Asia/Taipei"
	if dbConfig["timezone"] != "" {
		timeZone = dbConfig["timezone"]
	}
	if dbType == "sqlite" {
		return gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	} else if dbType == "mysql" {
		// default "3306"
		dsn := user + ":" + password + "@" + fmt.Sprintf("(%s:%s)/%s", host, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=Local"
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if dbType == "postgres" {
		host := dbConfig["host"]
		dbname := dbConfig["dbname"]
		// default port "5432"
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s", host, port, user, dbname, password, sslmode, timeZone)
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		return nil, errors.New(dbType + " is an invalid type")
	}
}

func CreateDefaultTable(db *gorm.DB) {
	db.AutoMigrate(&models.UserRole{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserSession{})
	// create default user_role
	count := int64(0)
	if db.Model(models.UserRole{}).Where("code=?", 1).Count(&count); count == 0 {
		if err := db.Create(&models.DefaultUserRole[0]).Error; err != nil {
			log.Println("create default admin user_role error", err)
		}
		if err := db.Create(&models.DefaultUserRole[1]).Error; err != nil {
			log.Println("create default regular user_role error", err)
		}
	}
	// create default user
	if db.Model(models.User{}).Where("account=?", "admin").Count(&count); count == 0 {
		if err := db.Create(&models.User{
			Account:    "admin",
			Password:   utils.Cipher.EncodePassword("admin"),
			Name:       "管理者帳號",
			Email:      "admin@testmail.com",
			UserRoleID: 1,
		}).Error; err != nil {
			log.Println("create default admin user error", err)
		}
	}
	if db.Model(models.User{}).Where("account=?", "testuser").Count(&count); count == 0 {
		if err := db.Create(&models.User{
			Account:    "testuser",
			Password:   utils.Cipher.EncodePassword("testuser"),
			Name:       "測試一般使用者",
			Email:      "testuser@testmail.com",
			UserRoleID: 2,
		}).Error; err != nil {
			log.Println("create default regular user error", err)
		}
	}
}
