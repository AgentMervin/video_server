package main

import (
	"net/http"
	// "html/template"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandler() *httprouter.Router {
	router:=httprouter.New()
	router.GET("/", homeHandler)
	router.POST("/", homeHandler)
	router.GET("/userhome", userHomeHandler)
	router.POST("/userhome", userHomeHandler)
	router.POST("/api", apiHandler)
	router.POST("/upload/:vid-id", proxyHandler)
	router.ServeFiles("/statics/*filepath", http.Dir("./template"))
	return router
}


func main() {
	r:=RegisterHandler()
	http.ListenAndServe(":8080", r)
}
