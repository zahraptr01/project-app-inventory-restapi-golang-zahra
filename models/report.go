package models

type Report struct {
	TotalItems     int `json:"total_items"`
	TotalSales     int `json:"total_sales"`
	TotalRevenue   int `json:"total_revenue"`
	LowStockAlerts int `json:"low_stock_alerts"`
}
