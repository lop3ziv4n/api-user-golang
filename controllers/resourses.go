package controllers

import (
	"workspaces/microservices-docker-go-mongodb-master/users/models"
)

type (
	// For Get - /users
	UsersResource struct {
		Data []models.User `json:"data"`
	}
	// For Post/Put - /users
	UserResource struct {
		Data models.User `json:"data"`
	}
)
