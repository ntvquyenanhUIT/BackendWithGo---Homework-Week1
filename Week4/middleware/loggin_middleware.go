package middleware

import (
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {

	logDir := "log"
	logFilePath := path.Join(logDir, "logs.txt")
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	logger := log.New(file, "", log.LstdFlags)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logger.Printf("%s %s %s %s", r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))

	})
}
