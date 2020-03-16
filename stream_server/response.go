package main

import (
	"io"
	//"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, sc int, errMsg string){
	w.WriteHeader(sc)
	io.WriteString(w, errMsg)
}