package routes

import (
	"net/http"

	"github.com.br/rodrygo262/hackaton_golang/controllers"
)

func LoadRoutes() {
	http.HandleFunc("/insert_provisioning", controllers.Insert)
	http.HandleFunc("/list_provisionings", controllers.List)
}
