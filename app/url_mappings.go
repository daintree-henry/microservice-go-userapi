package app

import (
	"github.com/daintree-henry/microservice-go-userapi/controller/health"
	"github.com/daintree-henry/microservice-go-userapi/controller/users"
)

func mapUrls() {
	router.GET("/health", health.Health)

	router.POST("/users", users.CtrCreateUser)
	router.GET("/users/:id", users.CtrGetUserById)
	router.PUT("/users/:id", users.CtrUpdateUser)
	router.PATCH("/users/:id", users.CtrUpdateUser)
	router.DELETE("/users", users.CtrDeleteUserById)

	router.POST("/login", users.CtrLoginUserByIdAndPw)

	router.PUT("/search/phone", users.CtrFindIdByPhoneNumber)
	router.GET("/search/status", users.CtrGetUsersByStatus)

}
