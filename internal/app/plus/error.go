package plus

/*
	错误状态码[状态码错误]
*/
var (

	// ******** User ****************//
	MSC_PasswordError = NewWrapResponse(1002, "密码错误", 1002)
	MSC_NotUser       = NewWrapResponse(1003, "用户不存在", 1003)
	MSC_NotApp        = NewWrapResponse(1004, "应用不存在", 1004)
	MSC_NotApi        = NewWrapResponse(1007, "没有当前api", 1007)
	//************** Token *************/
	MSC_Login        = NewWrapResponse(1300, "请先登录", 1300)
	MSC_TokenExpired = NewWrapResponse(1301, "token已过期", 1301)
	MSC_InvalidToken = NewWrapResponse(1302, "token不合法", 1302)
	MSC_NotToken     = NewWrapResponse(1303, "请求未携带token，无权限访问", 1303)
	// ********** AcessToken ************//
	MSC_NotAcessToken     = NewWrapResponse(1401, "access_token为空", 1401)
	MSC_InvalidAcessToken = NewWrapResponse(1402, "access_token不合法", 1402)
	MSC_AcessTokenExpired = NewWrapResponse(1401, "access_token过期", 1401)

	MSC_UnKnownError = NewWrapResponse(1500, "未知错误", 1500)

	MSC_NotAuthorization = NewWrapResponse(401, "无访问权限", 401)
	MSC_MethodNotAllow   = NewWrapResponse(405, "请求方法不允许", 405)

	MSC_NotFound    = NewWrapResponse(404, "资源不存在", 404)
	MSC_RequstFrecy = NewWrapResponse(429, "请求过于频繁", 429)
	MSC_ServerError = NewWrapResponse(500, "服务器发生错误", 500)

	MSC_RequestError      = NewWrap400Response("请求发生错误")
	MSC_ValidateFormError = NewWrap500Response("表单验证码失败")
)
