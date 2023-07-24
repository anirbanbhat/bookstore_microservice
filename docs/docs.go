// Package docs contains the Swagger documentation definitions.
package docs

import (
	_ "github.com/swaggo/echo-swagger"                               // Import echo-swagger to generate Swagger documentation.
	_ "github.com/your-username/bookstore_microservice/api/handlers" // Import API handlers to generate Swagger documentation.
)
