package main

import (
	"log"
	"time"

	handlerPing "github.com/burgosfacundo/ApiGo.git/cmd/server/handler/ping"
	handlerProduct "github.com/burgosfacundo/ApiGo.git/cmd/server/handler/products"
	"github.com/burgosfacundo/ApiGo.git/internal/domain"
	"github.com/burgosfacundo/ApiGo.git/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {

	// Loads the db into dinamyc memory
	db := LoadStore()

	// Ping.
	controllerPing := handlerPing.NewControllerPing()

	// Products.
	repository := products.NewMemoryRepository(db)
	service := products.NewServiceProduct(repository)
	controllerProduct := handlerProduct.NewControllerProducts(service)

	engine := gin.Default()

	// /api/v1 Group
	group := engine.Group("/api/v1")
	{
		group.GET("/ping", controllerPing.HandlerPing())

		// /product group
		grupoProduct := group.Group("/product")
		{
			// POST /product 	for create a new product
			grupoProduct.POST("", controllerProduct.HandlerCreate())

			// GET /product 	for get all the products
			grupoProduct.GET("", controllerProduct.HandlerGetAll())

			// GET /product/:id 	for get a single product for id
			grupoProduct.GET("/:id", controllerProduct.HandlerGetByID())

			// PUT /product/:id 	for edit a single product for id
			grupoProduct.PUT("/:id", controllerProduct.HandlerUpdate())

			// DELETE /product/:id 	for delete a single product for id
			grupoProduct.DELETE("/:id", controllerProduct.HandlerDelete())

		}

	}

	// Run the engine in the port 8080
	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// LoadStore loads the db into dinamic memory
func LoadStore() []domain.Product {
	return []domain.Product{
		{
			Id:          "1",
			Name:        "Coco Cola",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.5,
		},
		{
			Id:          "2",
			Name:        "Pepsito",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.5,
		},
		{
			Id:          "3",
			Name:        "Fantastica",
			Quantity:    10,
			CodeValue:   "123456789",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.5,
		},
	}
}
