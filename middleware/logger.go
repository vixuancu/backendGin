package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	// nơi chứa log request
	logPath := "logs/htttp.log"
	// Tạo thư mục cha nếu chưa tồn tại
	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		panic(err)
	}
	//os.O_APPEND: Ghi dữ liệu vào cuối file (append mode).
	//
	//os.O_CREATE: Nếu file chưa tồn tại, nó sẽ được tạo mới.
	//
	//	os.O_WRONLY: Mở file ở chế độ chỉ ghi (write-only).
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	// Tạo logger với zerolog
	logger := zerolog.New(logFile).With().Timestamp().Logger()
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// Tính toán thời gian xử lý request
		duration := time.Since(start)
		statusCode := c.Writer.Status()
		logEvent := logger.Info() // Mặc định ghi log ở mức Info
		if statusCode >= 500 {
			// Nếu mã trạng thái là 500 trở lên, ghi log ở mức Error
			logEvent = logger.WithLevel(zerolog.ErrorLevel)
		} else if statusCode >= 400 {
			// Nếu mã trạng thái là 400 trở lên, ghi log ở mức Warn
			logEvent = logger.WithLevel(zerolog.WarnLevel)
		} else {
			// Mã trạng thái thành công (2xx), ghi log ở mức Info
			logEvent = logger.WithLevel(zerolog.InfoLevel)
		}

		logEvent.
			Str("method", c.Request.Method).                  // Ghi phương thức HTTP(GET, POST, PUT, DELETE, v.v.)
			Str("path", c.Request.URL.Path).                  // Ghi đường dẫn của request(ví dụ: /api/v1/users)
			Str("query", c.Request.URL.RawQuery).             // Ghi query string nếu có (ví dụ: ?page=1&limit=10)
			Str("client_ip", c.ClientIP()).                   // Ghi địa chỉ IP của client
			Str("user_agent", c.Request.UserAgent()).         // Ghi user agent của client (trình duyệt, ứng dụng, v.v.)
			Str("referer", c.Request.Referer()).              // Ghi referer nếu có (trang trước đó mà client đã truy cập)
			Str("protocol", c.Request.Proto).                 // Ghi giao thức HTTP (HTTP/1.1, HTTP/2, v.v.)
			Str("host", c.Request.Host).                      // Ghi host của request (ví dụ: example.com)
			Str("remote_address", c.Request.RemoteAddr).      // nếu địa chỉ IP của client không được cung cấp bởi c.ClientIP()
			Str("request_uri", c.Request.RequestURI).         // Ghi toàn bộ URI của request (bao gồm query string)
			Int64("content_length", c.Request.ContentLength). // Ghi độ dài của nội dung request  (nếu có)
			Interface("headers", c.Request.Header).           // Ghi tất cả các header của request
			Int("status_code", statusCode).                   // Ghi mã trạng thái HTTP của response (ví dụ: 200, 404, 500, v.v.)
			Int64("duration_ms", duration.Milliseconds()).    // Ghi thời gian xử lý request tính bằng mili giây
			Msg("HTTP Request Log")

	}
}
