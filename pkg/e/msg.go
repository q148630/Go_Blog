package e

var MsgFlags = map[int]string {
	SUCCESS: "ok",
	ERROR: "fail",
	INVALID_PARAMS : "請求參數錯誤",
	ERROR_EXIST_TAG : "已存在該標籤名稱",
	ERROR_NOT_EXIST_TAG : "該標籤不存在",
	ERROR_NOT_EXIST_ARTICLE : "該文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token驗證失敗",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超時",
	ERROR_AUTH_TOKEN : "Token產生失敗",
	ERROR_AUTH : "Token錯誤",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}