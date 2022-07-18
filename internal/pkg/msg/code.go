package msg

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	// Authentication
	ERROR_AUTH                     = 10001
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 10002
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 10003
	ERROR_AUTH_TOKEN               = 10004
	// Account
	ERROR_ADD_ACCOUNT_FAIL = 10005
	ERROR_EXIST_ACCOUNT    = 10006
	// common object
	ERROR_NOT_EXIST   = 10007
	ERROR_GET_FAIL    = 10008
	ERROR_ADD_FAIL    = 10009
	ERROR_UPDATE_FAIL = 10010
	ERROR_DELETE_FAIL = 10011
)
