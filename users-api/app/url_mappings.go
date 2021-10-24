package app

import (
	"github.com/sijanstha/controllers/ping"
	"github.com/sijanstha/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/internal/users/search", users.Search)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.PUT("/users", users.UpdateUser)
	router.PATCH("/users", users.UpdateUser)
	router.POST("/users/login", users.Login)
}
