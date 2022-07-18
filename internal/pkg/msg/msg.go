package msg

var MsgFlags = map[int]string{
	SUCCESS:        "Success",
	ERROR:          "Fail",
	INVALID_PARAMS: "Invalid param",

	ERROR_AUTH:                     "Authentication failed",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Access token is invalid",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Access token is expired",
	ERROR_AUTH_TOKEN:               "Validate token failed",

	ERROR_ADD_ACCOUNT_FAIL: "Create account failed",
	ERROR_EXIST_ACCOUNT:    "Email account already exists!",

	ERROR_NOT_EXIST:   "Resource is not exist",
	ERROR_GET_FAIL:    "Get data failed",
	ERROR_ADD_FAIL:    "Add failed",
	ERROR_UPDATE_FAIL: "Update failed",
	ERROR_DELETE_FAIL: "Delete failed",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
