package database

import (
	"github.com/savioafs/book-market/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitGorm(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.Book{}, &entity.DiscountCoupon{}, &entity.Seller{}, &entity.Sale{}, &entity.Review{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
