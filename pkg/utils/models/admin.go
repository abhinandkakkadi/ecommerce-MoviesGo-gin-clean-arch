package models

type AdminLogin struct {
	Email    string `json:"email" binding:"required" validate:"required"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}

type AdminDetails struct {
	ID    uint   `json:"id" gorm:"uniquekey; not null"`
	Name  string `json:"name"  gorm:"validate:required"`
	Email string `json:"email"  gorm:"validate:required"`
}

type AdminSignUp struct {
	Name            string `json:"name" binding:"required" gorm:"validate:required"`
	Email           string `json:"email" binding:"required" gorm:"validate:required"`
	Password        string `json:"password" binding:"required" gorm:"validate:required"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
}

type AdminDetailsResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email" `
}

type SalesReport struct {
	TotalSales      float64
	TotalOrders     int
	CompletedOrders int
	PendingOrders   int
	TrendingProduct string
}

// ADMIN DASHBOARD COMPLETE DETAILS

type DashboardRevenue struct {
	TodayRevenue float64
	MonthRevenue float64
	YearRevenue  float64
}

type DashboardOrder struct {
	CompletedOrder int
	PendingOrder   int
	CancelledOrder int
	TotalOrder     int
	TotalOrderItem int
}

type DashboardAmount struct {
	CreditedAmount float64
	PendingAmount  float64
}

type DashboardUser struct {
	TotalUsers   int
	BlockedUser  int
	OrderedUsers int
}

type DashBoardProduct struct {
	TotalProducts     int
	OutOfStockProduct int
	TopSellingProduct string
}

type CompleteAdminDashboard struct {
	DashboardRevenue DashboardRevenue
	DashboardOrder   DashboardOrder
	DashboardAmount  DashboardAmount
	DashboardUser    DashboardUser
	DashBoardProduct DashBoardProduct
}
