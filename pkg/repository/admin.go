package repository

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/domain"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/helper"
	interfaces "github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/repository/interface"
	"github.com/abhinandkakkadi/ecommerce-MoviesGo-gin-clean-arch/pkg/utils/models"
	"gorm.io/gorm"
)

type adminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminRepository{
		DB: DB,
	}
}

func (ad *adminRepository) LoginHandler(adminDetails models.AdminLogin) (domain.Admin, error) {

	var adminCompareDetails domain.Admin
	if err := ad.DB.Raw("select * from admins where email = ? ", adminDetails.Email).Scan(&adminCompareDetails).Error; err != nil {
		return domain.Admin{}, err
	}

	return adminCompareDetails, nil
}

// check if an admin with specified email already exist
func (ad *adminRepository) CheckAdminAvailability(admin models.AdminSignUp) bool {

	var count int
	if err := ad.DB.Raw("select count(*) from admins where email = ?", admin.Email).Scan(&count).Error; err != nil {
		return false
	}

	return count > 0

}

func (ad *adminRepository) CreateAdmin(admin models.AdminSignUp) (models.AdminDetailsResponse, error) {

	var adminDetails models.AdminDetailsResponse
	if err := ad.DB.Raw("insert into admins (name,email,password) values (?, ?, ?) RETURNING id, name, email", admin.Name, admin.Email, admin.Password).Scan(&adminDetails).Error; err != nil {
		return models.AdminDetailsResponse{}, err
	}

	return adminDetails, nil

}

// Get users details for authenticated admins
func (ad *adminRepository) GetUsers(page int, count int) ([]models.UserDetailsAtAdmin, error) {

	if page == 0 {
		page = 1
	}
	offset := (page - 1) * count
	var userDetails []models.UserDetailsAtAdmin

	if err := ad.DB.Raw("select id,name,email,phone,blocked from users limit ? offset ?", count, offset).Scan(&userDetails).Error; err != nil {
		return []models.UserDetailsAtAdmin{}, err
	}

	return userDetails, nil

}

func (ad *adminRepository) GetGenres() ([]domain.Genre, error) {

	var genres []domain.Genre
	if err := ad.DB.Raw("select * from genres").Scan(&genres).Error; err != nil {
		return []domain.Genre{}, err
	}

	return genres, nil

}

// CATEGORY MANAGEMENT
func (ad *adminRepository) AddGenre(genre models.CategoryUpdate) error {

	var count int
	err := ad.DB.Raw("select count(*) from genres where genre_name = ?", genre.Genre).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("the genre already exist")
	}

	if err := ad.DB.Exec("insert into genres (genre_name) values (?) ", genre.Genre).Error; err != nil {
		return err
	}
	return nil

}

func (ad *adminRepository) Delete(genre_id string) error {

	id, err := strconv.Atoi(genre_id)
	if err != nil {
		return err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from genres where id = ?").Scan(&count).Error; err != nil {
		return err
	}
	if count < 1 {
		return errors.New("genre for given id does not exist")
	}

	query := fmt.Sprintf("delete from genres where id = '%d'", id)
	if err := ad.DB.Exec(query).Error; err != nil {
		return err
	}

	return nil

}

func (ad *adminRepository) GetUserByID(id string) (domain.Users, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		return domain.Users{}, err
	}

	var count int
	if err := ad.DB.Raw("select count(*) from users where id = ?").Scan(&count).Error; err != nil {
		return domain.Users{}, err
	}
	if count < 1 {
		return domain.Users{}, errors.New("user for the given id does not exist")
	}

	query := fmt.Sprintf("select * from users where id = '%d'", user_id)
	var userDetails domain.Users

	if err := ad.DB.Raw(query).Scan(&userDetails).Error; err != nil {
		return domain.Users{}, err
	}

	return userDetails, nil

}

// function which will both block and unblock a user
func (ad *adminRepository) UpdateBlockUserByID(user domain.Users) error {

	err := ad.DB.Exec("update users set blocked = ? where id = ?", user.Blocked, user.ID).Error
	if err != nil {
		fmt.Println("Error updating user:", err)
		return err
	}

	return nil

}

func (ad *adminRepository) FilteredSalesReport(startTime time.Time, endTime time.Time) (models.SalesReport, error) {

	var salesReport models.SalesReport

	result := ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.TotalSales)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = ad.DB.Raw("select count(*) from orders").Scan(&salesReport.TotalOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = ad.DB.Raw("select count(*) from orders where payment_status = 'paid' and approval = true and  created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.CompletedOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = ad.DB.Raw("select count(*) from orders where shipment_status = 'processing' and approval = false and  created_at >= ? and created_at <= ?", startTime, endTime).Scan(&salesReport.PendingOrders)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	var productID int
	result = ad.DB.Raw("select product_id from order_items group by product_id order by sum(quantity) desc limit 1").Scan(&productID)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}

	result = ad.DB.Raw("select movie_name from products where id = ?", productID).Scan(&salesReport.TrendingProduct)
	if result.Error != nil {
		return models.SalesReport{}, result.Error
	}
	fmt.Println(salesReport.TrendingProduct)

	return salesReport, nil
}

func (ad *adminRepository) TotalRevenue() (models.DashboardRevenue, error) {

	var revenueDetails models.DashboardRevenue
	startTime := time.Now().AddDate(0, 0, -1)
	endTime := time.Now()
	err := ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.TodayRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}

	startTime, endTime = helper.GetTimeFromPeriod("month")
	err = ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.MonthRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}

	startTime, endTime = helper.GetTimeFromPeriod("year")
	err = ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true and created_at >= ? and created_at <= ?", startTime, endTime).Scan(&revenueDetails.YearRevenue).Error
	if err != nil {
		return models.DashboardRevenue{}, nil
	}

	return revenueDetails, nil
}

func (ad *adminRepository) DashBoardOrder() (models.DashboardOrder, error) {

	var orderDetails models.DashboardOrder
	err := ad.DB.Raw("select count(*) from orders where payment_status = 'paid' and approval = true ").Scan(&orderDetails.CompletedOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = ad.DB.Raw("select count(*) from orders where shipment_status = 'pending' or shipment_status = 'processing'").Scan(&orderDetails.PendingOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = ad.DB.Raw("select count(*) from orders where shipment_status = 'cancelled'").Scan(&orderDetails.CancelledOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = ad.DB.Raw("select count(*) from orders").Scan(&orderDetails.TotalOrder).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	err = ad.DB.Raw("select sum(quantity) from order_items").Scan(&orderDetails.TotalOrderItem).Error
	if err != nil {
		return models.DashboardOrder{}, nil
	}

	return orderDetails, nil

}

func (ad *adminRepository) AmountDetails() (models.DashboardAmount, error) {

	var amountDetails models.DashboardAmount
	err := ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'paid' and approval = true ").Scan(&amountDetails.CreditedAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}

	err = ad.DB.Raw("select coalesce(sum(final_price),0) from orders where payment_status = 'not paid' and shipment_status = 'processing' or shipment_status = 'pending' or shipment_status = 'order placed' ").Scan(&amountDetails.PendingAmount).Error
	if err != nil {
		return models.DashboardAmount{}, nil
	}

	return amountDetails, nil

}

func (ad *adminRepository) DashboardUserDetails() (models.DashboardUser, error) {

	var userDetails models.DashboardUser
	err := ad.DB.Raw("select count(*) from users").Scan(&userDetails.TotalUsers).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	err = ad.DB.Raw("select count(distinct user_id) from orders").Scan(&userDetails.OrderedUsers).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	err = ad.DB.Raw("select count(*) from users where blocked = true").Scan(&userDetails.BlockedUser).Error
	if err != nil {
		return models.DashboardUser{}, nil
	}

	return userDetails, nil
}

func (ad *adminRepository) DashBoardProductDetails() (models.DashBoardProduct, error) {

	var productDetails models.DashBoardProduct
	err := ad.DB.Raw("select count(*) from products").Scan(&productDetails.TotalProducts).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	err = ad.DB.Raw("select count(*) from products where quantity = 0").Scan(&productDetails.OutOfStockProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	var productID int
	err = ad.DB.Raw("select product_id from order_items group by product_id order by sum(quantity) desc limit 1").Scan(&productID).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	err = ad.DB.Raw("select movie_name from products where id = ?", productID).Scan(&productDetails.TopSellingProduct).Error
	if err != nil {
		return models.DashBoardProduct{}, nil
	}

	return productDetails, nil
}
