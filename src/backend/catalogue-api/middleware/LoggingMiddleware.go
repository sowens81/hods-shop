package middleware

import (
	"fmt"
	"time"

	"catalogue-api/logging" // Updated import path

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LoggingMiddleware is a middleware that logs incoming HTTP requests and their responses
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate or retrieve the correlation ID from request headers
		correlationID := c.GetHeader("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String() // Generate a new correlation ID if missing
		}

		// Get the singleton logger instance
		logger := logging.GetLogger()

		// Log the incoming request with the correlation ID
		logger.LogWithCorrelationID(correlationID, fmt.Sprintf("Request: %s %s", c.Request.Method, c.Request.URL.Path))

		// Record the time taken to process the request
		startTime := time.Now()

		// Process the request
		c.Next()

		// Log the response and the time taken
		duration := time.Since(startTime)
		logger.LogWithCorrelationID(correlationID, fmt.Sprintf("Response: %s %s, Duration: %v", c.Request.Method, c.Request.URL.Path, duration))
	}
}
