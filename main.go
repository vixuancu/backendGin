package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/demo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": []string{"Alice", "Bob", "Charlie"},
		})
	})
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id":   id,
			"data": []string{"Alice", "Bob", "Charlie"},
		})
	})
	r.GET("/products", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": []string{"Laptop", "Smartphone", "Tablet"},
		})
	})
	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		price := c.Query("price")
		c.JSON(200, gin.H{
			"id":    id,
			"price": price,
			"data":  []string{"Laptop", "Smartphone", "Tablet"},
		})
	})
	r.Run(":8080") // listen and serve on
}
