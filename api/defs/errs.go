package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrResponse{HttpSC: 400,
		Error: Err{Error: "RequestBody is error", ErrorCode: "001"}}
	ErrorNotAuthUser = ErrResponse{HttpSC: 401,
		Error: Err{Error: "User authentication is error", ErrorCode: "002"}}
	ErrorDBError = ErrResponse{HttpSC: 500,
		Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFault = ErrResponse{HttpSC: 500,
		Error: Err{Error: "Internal error", ErrorCode: "004"}}
)
