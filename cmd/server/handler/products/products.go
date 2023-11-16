package products

import (
	"net/http"

	"github.com/burgosfacundo/ApiGo.git/internal/domain"
	"github.com/burgosfacundo/ApiGo.git/internal/products"
	"github.com/gin-gonic/gin"
)

// Controller is a struct that contains the service of Product objects
type Controller struct {
	service products.Service
}

// NewControllerProducts is a function that loads the service into the controller
func NewControllerProducts(service products.Service) *Controller {
	return &Controller{service: service}
}

// HandlerCreate is a function that calls the service for create a Product in the db
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var productRequest domain.Product

		// We receive the product
		err := ctx.Bind(&productRequest)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
		}

		// We call the service to create the product
		product, err := c.service.Create(ctx, productRequest)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}

		// We return the product that was created
		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})

	}
}

// HandlerGetAll is a function that calls the service for get all the products in the db
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// We call the service to get all the products
		listProducts, err := c.service.GetAll(ctx)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}

		// We return the list of products
		ctx.JSON(http.StatusOK, gin.H{
			"data": listProducts,
		})
	}
}

// HandlerGetByID is a function that calls the service for get a product by id
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// We receive the id of the product
		idParam := ctx.Param("id")

		// We call the service to update the product
		product, err := c.service.GetByID(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}

		// We return the product
		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}
}

// HandlerUpdate is a function that calls the service for update a product by id
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// We receive the id of the product
		idParam := ctx.Param("id")

		var productRequest domain.Product

		// We receive the new atributes of the product
		err := ctx.Bind(&productRequest)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
				"error":   err,
			})
		}

		// We call the service to update the product
		product, err := c.service.Update(ctx, productRequest, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}

		// We return the product that was updated
		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}
}

// HandlerDelete is a function that calls the service for delete a product by id
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// We receive the id of the product
		idParam := ctx.Param("id")

		// We call the service to delete the product by id
		err := c.service.Delete(ctx, idParam)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
			})
		}

		// We return the confirmation of the delete
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Product eliminated",
		})
	}
}
