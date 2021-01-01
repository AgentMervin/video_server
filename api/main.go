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
	router.POST("/user", CreateUser)//function作为变量传入
	return router
}

func main(){
	r :=RegisterHandlers()
	http.ListenAndServe(":8000",r)

}

//handler -> validation{1.request, 2.user}->business logic -> response.