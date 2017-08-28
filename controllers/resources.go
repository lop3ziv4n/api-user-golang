package controllers

import (
	"github.com/lop3ziv4n/api-user-golang/models"
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
