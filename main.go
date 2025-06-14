package main

import (
	handlerV1 "ginAPI/internal/api/v1/handler"
	handlerV2 "ginAPI/internal/api/v2/handler"
	"ginAPI/middleware"
	"ginAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	r := gin.Default()
	go middleware.CleanupClients() // dọn dẹp các client đã hết hạn
	if err := utils.RegisterValidators(); err != nil {
		panic("Failed to register custom validators: " + err.Error()) // stop the server if validators cannot be registered
	}
	err := godotenv.Load(".env") // Load environment variables from .env file
	if err != nil {
		log.Println("Error loading .env file")
	}
	r.Use(middleware.LoggerMiddleware(), middleware.APIKeyMiddleware(), middleware.RateLimitMiddleware())
	//verrsion 1
	v1 := r.Group("/api/v1")
	{
		userHandlerV1 := handlerV1.NewUserHandler()
		productHandler := handlerV1.NewProductHandler()
		//user
		v1.GET("/users", userHandlerV1.GetUsersV1)
		v1.GET("/users/:id", userHandlerV1.GetUsersByIdV1)
		v1.GET("/users/admin/:uid", userHandlerV1.GetUsersByUidV1)
		v1.GET("/users/admintest/:slug", userHandlerV1.GetUsersBySlugV1)
		v1.POST("/users", userHandlerV1.PostUsers)
		v1.PUT("/users/:id", userHandlerV1.PutUsers)
		v1.DELETE("/users/:id", userHandlerV1.DeleteUsers)
		// product
		v1.GET("/products", productHandler.GetProductsV1)
		v1.GET("/products/:slug", productHandler.GetProductsBySlugV1)
		v1.POST("/products", productHandler.PostProducts)
		v1.PUT("/products/:id", productHandler.PutProducts)
		v1.DELETE("/products/:id", productHandler.DeleteProducts)
		// category
		category := v1.Group("/category")
		{
			categoryHandlerV1 := handlerV1.NewCategoryHandler()
			category.GET("/:category", categoryHandlerV1.GetCategoriesV1)
			category.POST("/", categoryHandlerV1.PostCategoriesV1)
		}

		// new

		news := v1.Group("/news")
		{
			newsHandlerV1 := handlerV1.NewNewsHandler()
			news.GET("/:slug", newsHandlerV1.GetNewsBySlugV1)
			news.GET("", newsHandlerV1.GetNewsV1)
			news.POST("", newsHandlerV1.PostNewsV1)
			news.POST("/upload-file", newsHandlerV1.PostUploadFileNewsV1)
			news.POST("/upload-multiple-file", newsHandlerV1.PostUploadMultipleFileNewsV1)
		}

	}

	//verrsion 2
	v2 := r.Group("/api/v2")
	{
		userHandlerV2 := handlerV2.NewUserHandler()
		//user
		v2.GET("/users", userHandlerV2.GetUsersV2)
		v2.GET("/users/:id", userHandlerV2.GetUsersByIdV2)
		v2.POST("/users", userHandlerV2.PostUsers)
		v2.PUT("/users/:id", userHandlerV2.PutUsers)
		v2.DELETE("/users/:id", userHandlerV2.DeleteUsers)

	}
	r.StaticFS("/images", gin.Dir("./upload", false)) // Serve static files from the "upload" directory
	r.Run(":8080")                                    // listen and serve on
}
