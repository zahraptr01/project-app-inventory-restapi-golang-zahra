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

	// Inisialisasi Repository
	categoryRepo := repository.NewCategoryRepository(db)
	itemRepo := repository.NewItemRepository(db)
	rackRepo := repository.NewRackRepository(db)
	warehouseRepo := repository.NewWarehouseRepository(db)
	userRepo := repository.NewUserRepository(db)
	saleRepo := repository.NewSaleRepository(db)
	reportRepo := repository.NewReportRepository(db)

	// Inisialisasi Service
	categoryService := service.NewCategoryService(categoryRepo)
	itemService := service.NewItemService(itemRepo)
	rackService := service.NewRackService(rackRepo)
	warehouseService := service.NewWarehouseService(warehouseRepo)
	userService := service.NewUserService(userRepo)
	saleService := service.NewSaleService(saleRepo)
	reportService := service.NewReportService(reportRepo)

	// Inisialisasi Handler
	categoryHandler := handler.NewCategoryHandler(categoryService)
	itemHandler := handler.NewItemHandler(itemService)
	rackHandler := handler.NewRackHandler(rackService)
	warehouseHandler := handler.NewWarehouseHandler(warehouseService)
	userHandler := handler.NewUserHandler(userService)
	saleHandler := handler.NewSaleHandler(saleService)
	reportHandler := handler.NewReportHandler(reportService)

	// Routing

	// Category Routes
	r.Route("/categories", func(r chi.Router) {
		r.Get("/", categoryHandler.GetAll)
		r.Get("/{id}", categoryHandler.GetByID)
		r.Post("/", categoryHandler.Create)
		r.Put("/{id}", categoryHandler.Update)
		r.Delete("/{id}", categoryHandler.Delete)
	})

	// Item Routes
	r.Route("/items", func(r chi.Router) {
		r.Get("/", itemHandler.GetAll)
		r.Get("/{id}", itemHandler.GetByID)
		r.Post("/", itemHandler.Create)
		r.Put("/{id}", itemHandler.Update)
		r.Delete("/{id}", itemHandler.Delete)
	})

	// Rack Routes
	r.Route("/racks", func(r chi.Router) {
		r.Get("/", rackHandler.GetAll)
		r.Get("/{id}", rackHandler.GetByID)
		r.Post("/", rackHandler.Create)
		r.Put("/{id}", rackHandler.Update)
		r.Delete("/{id}", rackHandler.Delete)
	})

	// Warehouse Routes
	r.Route("/warehouses", func(r chi.Router) {
		r.Get("/", warehouseHandler.GetAll)
		r.Get("/{id}", warehouseHandler.GetByID)
		r.Post("/", warehouseHandler.Create)
		r.Put("/{id}", warehouseHandler.Update)
		r.Delete("/{id}", warehouseHandler.Delete)
	})

	// User Routes
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetAll)
		r.Get("/{id}", userHandler.GetByID)
		r.Post("/", userHandler.Create)
		r.Put("/{id}", userHandler.Update)
		r.Delete("/{id}", userHandler.Delete)
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
		http.ServeFile(w, r, "templates/index.html")
	})

	return r

}
