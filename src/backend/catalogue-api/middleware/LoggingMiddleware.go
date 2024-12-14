package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/yourusername/yourproject/logging" // Import the logging package
)

// LoggingMiddleware logs details of incoming HTTP requests and responses.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get or generate a correlation ID
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String() // Generate a new correlation ID if missing
		}

		// Get the singleton logger instance
		logger := logging.GetLogger()

		// Log the incoming request details with the correlation ID
		logger.LogWithCorrelationID(correlationID, fmt.Sprintf("Request: %s %s", r.Method, r.URL.Path))

		// Record the time taken to process the request
		startTime := time.Now()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Log the response duration with the correlation ID
		duration := time.Since(startTime)
		logger.LogWithCorrelationID(correlationID, fmt.Sprintf("Response: %s %s, Duration: %v", r.Method, r.URL.Path, duration))
	})
}
