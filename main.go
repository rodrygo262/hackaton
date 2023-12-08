package main

import (
	"net/http"

	"github.com.br/rodrygo262/hackaton_golang/db/migration"
	"github.com.br/rodrygo262/hackaton_golang/routes"
)

func main() {
	routes.LoadRoutes()
	migration.AutoMigration()

	http.ListenAndServe(":3003", nil)
}
