package common

import "errors"

var (
	ErrBookIsRequired            = errors.New("book is required")
	ErrSellerIsRequired          = errors.New("seller is required")
	ErrBuyerIsRequired           = errors.New("buyer is required")
	ErrDiscountCouponIsNotActive = errors.New("discount coupon is not active")

	ErrTitleIsRequired         = errors.New("title is required")
	ErrImageIsRequired         = errors.New("image is required")
	ErrAuthorIsRequired        = errors.New("author is required")
	ErrPublisherIsRequired     = errors.New("publisher is required")
	ErrIsbnIsRequired          = errors.New("isbn is required")
	ErrCategoryIsRequired      = errors.New("category is required")
	ErrDescriptionIsRequired   = errors.New("description is required")
	ErrPriceIsRequired         = errors.New("price is required")
	ErrPublishedYearIsRequired = errors.New("published year is required")

	ErrCodeIsRequired               = errors.New("code is required")
	ErrDiscountPercentageIsRequired = errors.New("discount percentage is required")
	ErrExpirationDateIsRequired     = errors.New("expiration date is required")
	ErrUsageLimitIsRequired         = errors.New("usage limit is required")

	ErrSaleIsRequired       = errors.New("sale is required")
	ErrRatingIsRequired     = errors.New("rating is required")
	ErrCommissionIsRequired = errors.New("commission is required")

	ErrNameIsRequired  = errors.New("name is required")
	ErrEmailIsRequired = errors.New("email is required")
	ErrPhoneIsRequired = errors.New("phone is required")
)
