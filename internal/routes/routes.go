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

const (
// authorRoute   = "/author/"
// authorRouteID = "/author/{id}"
// authoresRoute = "/authors"
)

var RegisterRoutes = func(r *mux.Router) {

	//------------------------------------------------
	// Book Routes
	//------------------------------------------------
	r.HandleFunc(bookRouteID, handlers.GetBookByID).Methods(GET)
	r.HandleFunc(bookesRoute, handlers.GetAllBooks).Methods(GET)
	r.HandleFunc(bookRoute, handlers.CreateBook).Methods(POST)
	r.HandleFunc(bookRouteID, handlers.UpdateBook).Methods(PUT)
	r.HandleFunc(bookRouteID, handlers.DeleteBook).Methods(DELETE)

	//------------------------------------------------
	// Author Routes
	//------------------------------------------------
	// r.HandleFunc(authorRoute, handlers.GetAuthor).Methods(GET)
	// r.HandleFunc(authorRouteID, handlers.GetAuthorByID).Methods(GET)
	// r.HandleFunc(authoresRoute, handlers.GetAllAuthors).Methods(GET)
	// r.HandleFunc(authorRoute, handlers.CreateAuthor).Methods(POST)
	// r.HandleFunc(authorRouteID, handlers.UpdateAuthor).Methods(PUT)
	// r.HandleFunc(authorRouteID, handlers.DeleteAuthor).Methods(DELETE)
}
