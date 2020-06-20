package utils

import (
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
)

// AllowedOrigins ...
var AllowedOrigins = []string{
	"localhost",
	"0.0.0.0",
	"127.0.0.1",
}

// CORSOrigin ...
var CORSOrigin = handlers.AllowedOrigins(
	AllowedOrigins,
)

// CORSHeaders ...
var CORSHeaders = handlers.AllowedHeaders(
	[]string{"X-Requested-With", "X-Token-Auth", "Content-Type", "Authorization"},
)

// CORSMethods ...
var CORSMethods = handlers.AllowedMethods(
	[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
)

// SetCORSPolicy ...
var SetCORSPolicy = handlers.CORS(CORSMethods, CORSHeaders, CORSOrigin)

// SetCORSHeaders ...
func SetCORSHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Access-Control-Allow-Origin", strings.Join(AllowedOrigins, ","))
	writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, HEAD")
	writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Requested-With, X-Token-Auth, Authorization")
}
