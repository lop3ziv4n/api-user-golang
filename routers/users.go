package routers

import (
	"workspaces/microservices-docker-go-mongodb-master/users/controllers"

	"github.com/gorilla/mux"
)

func SetUsersRouters(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
	return router
}
