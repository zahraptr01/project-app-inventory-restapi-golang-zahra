package repository

import (
	"assignment5/models"
	"database/sql"
)

type ReportRepository interface {
	GenerateReport() (models.Report, error)
}

type reportRepo struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepo{db}
}

func (r *reportRepo) GenerateReport() (models.Report, error) {
	var report models.Report
	err := r.db.QueryRow(`
		SELECT 
			(SELECT COUNT(*) FROM items),
			(SELECT COUNT(*) FROM sales),
			(SELECT COALESCE(SUM(total), 0) FROM sales),
			(SELECT COUNT(*) FROM items WHERE stock < min_stock)
	`).Scan(&report.TotalItems, &report.TotalSales, &report.TotalRevenue, &report.LowStockAlerts)
	return report, err
}
