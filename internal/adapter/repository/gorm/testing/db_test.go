package testing

import (
	"context"
	"fmt"
	"testing"

	adapterGorm "github.com/NovanHsiu/go-demo-api-server/internal/adapter/repository/gorm"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func buildTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("local-test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// truncate all data
	// for sqlite only
	tables := []string{"users", "user_roles"}
	for _, table := range tables {
		db.Exec("DELETE FROM ?", table)
	}
	db.Exec("VACUUM")
	// create default data
	adapterGorm.CreateDefaultTable(db)
	return db, err
}

func initRepository(db *gorm.DB) *adapterGorm.GormRepository {
	return adapterGorm.NewGormRepository(context.Background(), db)
}

// nolint
func TestMain(m *testing.M) {
	// To avoid violating table constraints
	db, err := buildTestDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	testDB = db
	m.Run()
}
