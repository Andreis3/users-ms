package main

import (
	memoryRepository "github.com/andreis3/users-ms/src/infra/factory"
	httpConfig "github.com/andreis3/users-ms/src/infra/http"
)

var memoryRepositoryFactory memoryRepository.MemoryRepositoryFactory

func main() {
	http := httpConfig.NewGinHttp()
	routes := httpConfig.NewRouterConfig(http, &memoryRepositoryFactory)
	routes.Build()
	http.Listen("8080")
}
