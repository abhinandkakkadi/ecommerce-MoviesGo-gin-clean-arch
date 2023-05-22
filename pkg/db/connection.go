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

	db = db.Session(&gorm.Session{
    AllowGlobalUpdate: true, // This allows updating the zero value of the DeletedAt field
	})
	db = db.Set("gorm:delete_option", "OPTION(CASCADE)")	

	return db, dbErr
}
