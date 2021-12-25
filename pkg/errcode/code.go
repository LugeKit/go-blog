package errcode

var (
	Success                   = newError(0, "success")
	ServerError               = newError(10000000, "service internal error")
	InvalidParams             = newError(10000001, "invalid params")
	NotFound                  = newError(10000002, "path not found")
	UnauthorizedAuth          = newError(10000003, "invalid appkey and appSecret")
	UnauthorizedToken         = newError(10000004, "token not authorized")
	TokenTimeout              = newError(10000005, "error token timeout")
	UnauthorizedTokenGenerate = newError(10000006, "fail to generate token")
	TooManyRequest            = newError(10000007, "too many request")
)
