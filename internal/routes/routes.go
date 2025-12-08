package routes

import (
	"github.com/gorilla/mux"

	"github.com/MostafaSensei106/GoBookManagementAPI/internal/handlers"
)

const (
	Root        = "/"
	bookRoute   = "/book/"
	bookRouteID = "/book/{id}"
	bookesRoute = "/books"

	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

var RegisterRoutes = func(r *mux.Router) {

	r.HandleFunc(bookRoute, handlers.GetBook).Methods(GET)

	r.HandleFunc(bookRouteID, handlers.GetBookByID).Methods(GET)

	r.HandleFunc(bookesRoute, handlers.GetAllBooks).Methods(GET)
	r.HandleFunc(bookRoute, handlers.CreateBook).Methods(POST)

	r.HandleFunc(bookRouteID, handlers.UpdateBook).Methods(PUT)

	r.HandleFunc(bookRouteID, handlers.DeleteBook).Methods(DELETE)

}
