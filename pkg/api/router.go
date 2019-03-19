package api

import (
	"github.com/julienschmidt/httprouter"
)

func GetRouter(commit string) *httprouter.Router {
	router := httprouter.New()
	router.GET("/healthz", constructHealthz(commit))

	// PET
	router.POST("/pet", handleAddPet())
	router.GET("/pet/:petID", handleFindPet()) //also handles findByStatus
	router.DELETE("/pet/:petID", handleDeletePet())

	//router.GET("/pet/findByTags", handleNotImplemented()) // conflicts with :petID
	router.PUT("/pet", handleNotImplemented())
	router.POST("/pet/:petID", handleNotImplemented())
	router.POST("/pet/:petID/uploadImage", handleNotImplemented())

	// STORE
	router.GET("/store/inventory", handleNotImplemented())
	router.POST("/store/order", handleNotImplemented())
	router.GET("/store/order/:orderID", handleNotImplemented())
	router.DELETE("/store/order/:orderID", handleNotImplemented())

	// USER
	router.POST("/user", handleNotImplemented())
	router.POST("/user/createWithArray", handleNotImplemented())
	router.POST("/user/createWithList", handleNotImplemented())
	//router.GET("/user/login", handleNotImplemented()) // conflicts with :userName
	//router.GET("/user/logout", handleNotImplemented()) // conflicts with :userName
	router.GET("/user/:userName", handleNotImplemented())
	router.PUT("/user/:userName", handleNotImplemented())
	router.DELETE("/user/:userName", handleNotImplemented())

	return router
}
