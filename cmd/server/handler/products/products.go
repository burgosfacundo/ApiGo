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
// @Summary Post new product
// @Description Create a new product in the db
// @Tags Products
// @Accept json
// @Produce json
// @Param token header string true "TOKEN_ENV"
// @Param product body domain.Product true "Product"
// @Success 201 {object} domain.Product
// @Failure 400 "Bad Request"
// @Failure 500 "Internal Server Error"
// @Router /product [post]
func (c *Controller) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var productRequest domain.Product

		// We receive the product
		err := ctx.Bind(&productRequest)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
			return
		}

		// We call the service to create the product
		product, err := c.service.Create(ctx, productRequest)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		// We return the product that was created
		ctx.JSON(http.StatusCreated, product)

	}
}

// HandlerGetAll is a function that calls the service for get all the products in the db
// @Summary Get all the products
// @Description Return list of all the products in the db
// @Tags Products
// @Produces json
// @Success 200 {object} []domain.Product
// @Failure 500 "Internal Server Error"
// @Router /product [get]
func (c *Controller) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// We call the service to get all the products
		listProducts, err := c.service.GetAll(ctx)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Internal server error")
			return
		}

		// We return the list of products
		ctx.JSON(http.StatusOK, listProducts)
	}
}

// HandlerGetByID is a function that calls the service for get a product by id
// @Summary Get product by id
// @Description Return a product in the db
// @Tags Products
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} domain.Product
// @Failure 404 "Product Not Found"
// @Router /product/{id} [get]
func (c *Controller) HandlerGetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// We receive the id of the product
		idParam := ctx.Param("id")

		// We call the service to update the product
		product, err := c.service.GetByID(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Product not found")
			return
		}

		// We return the product
		ctx.JSON(http.StatusOK, product)
	}
}

// HandlerUpdate is a function that calls the service for update a product by id
// @Summary Update product
// @Description Update a product in the db
// @Tags Products
// @Produce json
// @Param id path string true "id"
// @Param product body domain.Product true "product"
// @Success 200 {object} domain.Product
// @Failure 400 "Bad Request"
// @Failure 404 "Product Not Found"
// @Router /product/{id} [put]
func (c *Controller) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// We receive the id of the product
		idParam := ctx.Param("id")

		var productRequest domain.Product

		// We receive the new atributes of the product
		err := ctx.Bind(&productRequest)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
			return
		}

		// We call the service to update the product
		product, err := c.service.Update(ctx, productRequest, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Product not found")
			return
		}

		// We return the product that was updated
		ctx.JSON(http.StatusOK, product)
	}
}

// HandlerDelete is a function that calls the service for delete a product by id
// @Summary Delete product
// @Description Delete a product in the db
// @Tags Products
// @Param id path string true "id"
// @Success 200 "OK"
// @Failure 404 "Product Not Found"
// @Router /product/{id} [delete]
func (c *Controller) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// We receive the id of the product
		idParam := ctx.Param("id")

		// We call the service to delete the product by id
		err := c.service.Delete(ctx, idParam)

		// If we have an error return it
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, "Product not found")
		}

		// We return the confirmation of the delete
		ctx.JSON(http.StatusOK, "Product eliminated")
	}
}
