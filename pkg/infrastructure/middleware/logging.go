package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// Logging returns a handlers.CombinedLoggingHandler with the http.Handler.
func Logging(next http.Handler) http.Handler {
	return handlers.CombinedLoggingHandler(os.Stdout, next)
}
