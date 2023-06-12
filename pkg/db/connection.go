package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/config"
	domain "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	db.AutoMigrate(&domain.Users{})
	db.AutoMigrate(&domain.Products{})
	db.AutoMigrate(&domain.Admin{})
	db.AutoMigrate(&domain.Cart{})
	db.AutoMigrate(&domain.Address{})
	db.AutoMigrate(&domain.PaymentMethod{})
	db.AutoMigrate(&domain.Order{})
	db.AutoMigrate(&domain.OrderItem{})
	db.AutoMigrate(&domain.Charge{})
	db.AutoMigrate(&domain.Coupons{})
	db.AutoMigrate(&domain.UsedCoupon{})
	db.AutoMigrate(&domain.OrderCoupon{})
	db.AutoMigrate(&domain.WishList{})
	db.AutoMigrate(&domain.Wallet{})
	db.AutoMigrate(&domain.RazerPay{})
	db.AutoMigrate(&domain.ProductOffer{})
	db.AutoMigrate(&domain.CategoryOffer{})
	db.AutoMigrate(&domain.Referral{})
	db.AutoMigrate(&domain.ProductOfferUsed{})
	db.AutoMigrate(&domain.CategoryOfferUsed{})

	return db, dbErr

}
