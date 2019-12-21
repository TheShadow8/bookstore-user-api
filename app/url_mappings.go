package app

import (
	"github.com/TheShadow8/bookstore-user-api/controllers/ping"
	"github.com/TheShadow8/bookstore-user-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users/:user_id", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Create)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/users/:user_id", users.Search)
}
