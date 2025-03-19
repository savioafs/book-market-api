package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitMySQL(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitGorm(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&entity.DiscountCoupon{},
		&entity.Seller{},
		&entity.Sale{},
		&entity.Review{},
		&entity.Client{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
