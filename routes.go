// Define all API routes here
package main

func initializeRoutes() {

  productRoutes := router.Group("/product")
  {
      // Get all Products
      productRoutes.GET("/",getAllProducts)

      // Get specific Product with ID
      productRoutes.GET("/:id",productById)

      // Update specific Product with ID
      productRoutes.PUT("/:id",updateProduct)

      // Create Product
      productRoutes.POST("/",createProduct)

  }

}
