package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video_server/api/defs"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	//从request里面读出create user相关的body
	res, _:=ioutil.ReadAll(r.Body)
	user_body:=&defs.UserCrenditial{}
	if err:=json.Unmarshal(res,user_body);err!=nil{
		return
	}
	id:=session.Generate
	su:=&defs.SignedUp{Success:true, SessionId:id}


	io.WriteString(w, "Create User")

}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	uname := p.ByName("user_name")
	io.WriteString(w,uname)
}