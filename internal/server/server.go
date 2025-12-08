package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/data/models"
	"github.com/MostafaSensei106/GoBookManagementAPI/internal/routes"
)

func Start() {
	models.Init()
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle(routes.Root, r)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
