package model

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func InitFromSQLLite(dbPath string) (func(), error) {
	db, closeFunc, err := NewGormDB(dbPath)
	if err != nil {
		panic(err)
	}
	err = migrateTable(db)
	return closeFunc, err
}

func NewGormDB(dbPath string) (*gorm.DB, func(), error) {
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10000)

	cleanFunc := func() {
		err = sqlDB.Close()
		if err != nil {
			log.Fatalf("Gorm db close error: %s", err)
		}
	}
	return db, cleanFunc, err
}

func migrateTable(db *gorm.DB) error {
	return db.AutoMigrate(
		new(Wager),
		new(Buy),
	)
}
