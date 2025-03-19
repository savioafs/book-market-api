package repository

import (
	"database/sql"
)

type ReviewRepositoryGorm struct {
	DB *sql.DB
}

func NewReviewRepositoryGorm(db *sql.DB) *ReviewRepositoryGorm {
	return &ReviewRepositoryGorm{DB: db}
}

/*
func (r *ReviewRepositoryGorm) CreateReview(review *entity.Review) error {
	return r.DB.Create(review).Error
}

func (r *ReviewRepositoryGorm) GetReviewByID(id string) (*entity.Review, error) {
	var review *entity.Review

	err := r.DB.First(&review, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return review, nil
}

func (r *ReviewRepositoryGorm) ExistsReview(saleID string) (bool, error) {
	var isReviewed bool

	err := r.DB.Model(&entity.Sale{}).Select("is_reviewed").Where("id = ?", saleID).Scan(&isReviewed).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return isReviewed, nil
}
*/
