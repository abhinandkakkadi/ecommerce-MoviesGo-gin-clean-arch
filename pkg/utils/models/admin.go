package models

type AdminLogin struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password" validate:"min=8,max=20"`
}

type AdminDetails struct {
	ID    uint   `json:"id" gorm:"uniquekey; not null"`
	Name  string `json:"name" gorm:"validate:required"`
	Email string `json:"email" gorm:"validate:required"`
}

type AdminSignUp struct {
	Name            string `json:"name" gorm:"validate:required"`
	Email           string `json:"email" gorm:"validate:required"`
	Password        string `json:"password" gorm:"validate:required"`
	ConfirmPassword string `json:"confirmpassword"`
}

type AdminDetailsResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email" `
}

// ADMIN DASHBOARD OVERVIEW

type SalesReport struct {
	TotalSales      float64
	TotalOrders     int
	CompletedOrders int
	PendingOrders   int
	TrendingProduct string
}
