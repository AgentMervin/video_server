package defs

type Err struct{
	Error string `json:"error"`
	ErrCode string `json:"err_code"`
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request body is not correct", ErrCode: "001"}}
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "User anthentication failed.", ErrCode: "002"}}
	ErrorDBError = ErrResponse{HttpSC: 500, Error: Err{Error: "DB ops failed", ErrCode: "003"}}
	ErrorInternalFaults = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrCode: "004"}}
)

