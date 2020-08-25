package main

import (
	"io"
	"encoding/json"
	"net/http"
	"video_server/api/defs"
)

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.WriteHeader(errResp.HttpSC)
	resStr, _:=json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func   malResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}