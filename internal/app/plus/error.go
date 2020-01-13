package plus

/*
	错误状态码[状态码错误]
*/
var (
	MSC_NotAuthorization =NewWrapResponse(401, "无访问权限", 401)
	MSC_MethodNotAllow = NewWrapResponse(405, "请求方法不允许", 405)
	MSC_InvalidToken    =NewWrapResponse(9999, "令牌失效", 9999)
	MSC_NotFound 		=NewWrapResponse(404, "资源不存在", 404)
	MSC_RequstFrecy	   = NewWrapResponse(429, "请求过于频繁", 429)
	MSC_ServerError    = NewWrapResponse(500, "服务器发生错误", 500)


	MSC_RequestError   = NewWrap400Response("请求发生错误")
)
