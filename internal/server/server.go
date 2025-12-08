package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/routes"
)

func Start() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle(routes.Root, r)

	// db := config.GetDB()
	// bookRepository := repository.NewBookRepository(db)

	log.Fatal(http.ListenAndServe(":9010", r))

}
