package main

import (
	"github.com/celio001/product_api.git/controller"
	"github.com/celio001/product_api.git/db"
	"github.com/celio001/product_api.git/repository"
	"github.com/celio001/product_api.git/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada Usecase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	// Camada de controller
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/products", ProductController.GetProduct)
	server.GET("/product/:productId", ProductController.GetProductById)
	server.POST("/product", ProductController.CreateProduct)
	server.PUT("/product/", ProductController.UpdateProduct)
	server.DELETE("/product/:productId", ProductController.DeleteProcuctById)

	server.Run(":8080")
}
