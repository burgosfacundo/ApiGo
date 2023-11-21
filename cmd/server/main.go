package main

import (
	"log"
	"time"

	handlerPing "github.com/burgosfacundo/ApiGo.git/cmd/server/handler/ping"
	handlerProduct "github.com/burgosfacundo/ApiGo.git/cmd/server/handler/products"
	"github.com/burgosfacundo/ApiGo.git/internal/domain"
	"github.com/burgosfacundo/ApiGo.git/internal/products"
	"github.com/burgosfacundo/ApiGo.git/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/burgosfacundo/ApiGo.git/cmd/server/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Products API
// @version         1.0
// @description     This is a api for Products.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Loads the db into dinamyc memory
	db := LoadStore()

	// Ping.
	controllerPing := handlerPing.NewControllerPing()

	// Products.
	repository := products.NewMemoryRepository(db)
	service := products.NewServiceProduct(repository)
	controllerProduct := handlerProduct.NewControllerProducts(service)

	engine := gin.Default()
	/*
		engine := gin.New()
		engine.Use(gin.Recovery())

		//Use the logger in the middleware and not gin's default logger
		engine.Use(middleware.Logger())
	*/

	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// /api/v1 Group
	group := engine.Group("/api/v1")
	{
		// /ping for testing
		group.GET("/ping", controllerPing.HandlerPing())

		// /product group
		grupoProduct := group.Group("/product")
		{
			// POST /product 	for create a new product
			grupoProduct.POST("", middleware.Auth(), controllerProduct.HandlerCreate())

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
