package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SimpleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// trước khi bắt đầu vào handler before
		log.Println("start func check middleware")
		c.Writer.Write([]byte("Start func check middleware")) // Ghi dữ liệu vào response trước khi tiếp tục xử lý
		c.Next()                                              // Call the next handler in the chain

		// sau khi đã xử lý xong handler after
		log.Println("end func check middleware")
		c.Writer.Write([]byte("Start func check middleware"))
	}
}
