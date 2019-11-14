package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"

)

type middleWareHandler struct{
	r *httprouter.Router
}
func NewMiddleWareHandler(){

}
func RegisterHandlers() *httprouter.Router{
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name",Login)
	router.GET("/user/:username", GetUserInfo)
	router.POST("/user/:username/videos", AddNewVideo)
	router.GET("/user/:username/videos", ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", DeleteVideo)
	router.POST("/videos/:vid-id/comments", PostComment)
	router.GET("/videos/:vid-id/comments", ShowComments)
	return router
}

func main(){
	r :=RegisterHandlers()
	http.ListenAndServe(":8000",r)

}