package gorm

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/NovanHsiu/go-demo-api-server/internal/adapter/gorm/model"
	"github.com/NovanHsiu/go-demo-api-server/internal/domain/common"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(dbConfig common.ConfigDB) (*gorm.DB, error) {
	host := dbConfig.Host
	dbname := dbConfig.Dbname
	port := dbConfig.Port
	sslmode := dbConfig.Sslmode
	user := dbConfig.User
	dbType := strings.ToLower(dbConfig.Type)
	password := dbConfig.Passwd
	timeZone := "Asia/Taipei"
	if dbConfig.Timezone != "" {
		timeZone = dbConfig.Timezone
	}
	if dbType == "sqlite" {
		return gorm.Open(sqlite.Open("local.db"), &gorm.Config{})
	} else if dbType == "mysql" {
		// default "3306"
		dsn := user + ":" + password + "@" + fmt.Sprintf("(%s:%s)/%s", host, port, dbname) + "?charset=utf8mb4&parseTime=True&loc=Local"
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if dbType == "postgres" {
		host := dbConfig.Host
		dbname := dbConfig.Dbname
		// default port "5432"
		dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s TimeZone=%s", host, port, user, dbname, password, sslmode, timeZone)
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		return nil, errors.New(dbType + " is an invalid type")
	}
}

func CreateDefaultTable(db *gorm.DB) {
	db.AutoMigrate(&model.UserRole{})
	db.AutoMigrate(&model.User{})
	// create default user_role
	count := int64(0)
	if db.Model(model.UserRole{}).Where("code=?", 1).Count(&count); count == 0 {
		if err := db.Create(&model.DefaultUserRole[0]).Error; err != nil {
			log.Println("create default admin user_role error", err)
		}
		if err := db.Create(&model.DefaultUserRole[1]).Error; err != nil {
			log.Println("create default regular user_role error", err)
		}
	}
	// create default user
	if db.Model(model.User{}).Where("account=?", "admin").Count(&count); count == 0 {
		if err := db.Create(&model.User{
			Account:    "admin",
			Password:   common.Cipher.EncodePassword("admin"),
			Name:       "管理者帳號",
			Email:      "admin@testmail.com",
			UserRoleID: 1,
		}).Error; err != nil {
			log.Println("create default admin user error", err)
		}
	}
	if db.Model(model.User{}).Where("account=?", "testuser").Count(&count); count == 0 {
		if err := db.Create(&model.User{
			Account:    "testuser",
			Password:   common.Cipher.EncodePassword("testuser"),
			Name:       "測試一般使用者",
			Email:      "testuser@testmail.com",
			UserRoleID: 2,
		}).Error; err != nil {
			log.Println("create default regular user error", err)
		}
	}
}
