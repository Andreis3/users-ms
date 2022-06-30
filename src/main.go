package main

import (
	memoryRepository "github.com/andreis3/users-ms/src/infra/factory"
	httpConfig "github.com/andreis3/users-ms/src/infra/http"
	"github.com/andreis3/users-ms/src/infra/middleware/auth"
)

var memoryRepositoryFactory memoryRepository.MemoryRepositoryFactory

func main() {
	http := httpConfig.NewGinHttp()
	middleware := auth.NewAuthMiddleware("12345")
	routes := httpConfig.NewRouterConfig(http, &memoryRepositoryFactory, middleware)
	routes.Build()
	http.Listen("8080")
}
