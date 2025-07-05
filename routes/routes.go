package routes

import (
	"assignment5/handler"
	"assignment5/repository"
	"assignment5/service"
	"net/http"

	"database/sql"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func SetupRoutes(db *sql.DB, logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	// --- Inisialisasi Repository ---
	categoryRepo := repository.NewCategoryRepository(db)
	itemRepo := repository.NewItemRepository(db)
	rackRepo := repository.NewRackRepository(db)
	warehouseRepo := repository.NewWarehouseRepository(db)
	userRepo := repository.NewUserRepository(db)
	saleRepo := repository.NewSaleRepository(db)
	reportRepo := repository.NewReportRepository(db)

	// --- Inisialisasi Service ---
	categoryService := service.NewCategoryService(categoryRepo)
	itemService := service.NewItemService(itemRepo)
	rackService := service.NewRackService(rackRepo)
	warehouseService := service.NewWarehouseService(warehouseRepo)
	userService := service.NewUserService(userRepo)
	saleService := service.NewSaleService(saleRepo)
	reportService := service.NewReportService(reportRepo)

	// --- Inisialisasi Handler ---
	categoryHandler := handler.NewCategoryHandler(categoryService)
	itemHandler := handler.NewItemHandler(itemService)
	rackHandler := handler.NewRackHandler(rackService)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	userHandler := handler.NewUserHandler(userService)
	saleHandler := handler.NewSaleHandler(saleService)
	reportHandler := handler.NewReportHandler(reportService)

	// --- Routing ---

	// Category Routes
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", categoryHandler.GetAll)
		r.Post("/", categoryHandler.Create)
	})

	// Item Routes
	r.Route("/items", func(r chi.Router) {
		r.Get("/", itemHandler.GetAll)
		r.Post("/", itemHandler.Create)
	})

	// Rack Routes
	r.Route("/racks", func(r chi.Router) {
		r.Get("/", rackHandler.GetAll)
		r.Post("/", rackHandler.Create)
	})

	// Warehouse Routes
	r.Route("/warehouses", func(r chi.Router) {
		r.Get("/", warehouseHandler.GetAll)
		r.Post("/", warehouseHandler.Create)
	})

	// User Routes (registration only, since no login/middleware)
	r.Route("/users", func(r chi.Router) {
		r.Post("/register", userHandler.Register)
	})

	// Sale Routes
	r.Route("/sales", func(r chi.Router) {
		r.Get("/", saleHandler.GetAll)
		r.Post("/", saleHandler.Create)
	})

	// Report Route (check manually for role "admin" in handler)
	r.Route("/report", func(r chi.Router) {
		r.Get("/", reportHandler.Generate)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Inventory API Menu</title>
			<style>
				body { font-family: Arial; padding: 20px; }
				h1 { color: #333; }
				ul { list-style: none; padding: 0; }
				li { margin-bottom: 10px; }
				a { text-decoration: none;}
			</style>
		</head>
		<body>
			<h1>Welcome to Inventory API</h1>
			<ul>
				<li><a href="/categories">Categories</a></li>
				<li><a href="/items">Items</a></li>
				<li><a href="/racks">Racks</a></li>
				<li><a href="/warehouses">Warehouses</a></li>
				<li><a href="/sales">Sales</a></li>
				<li><a href="/report">Report</a></li>
			</ul>
		</body>
		</html>
	`))
	})

	return r

}
