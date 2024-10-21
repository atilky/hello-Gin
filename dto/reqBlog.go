package dto

// 参数向结构体映射，并执行校验
type UpdateRequest struct {
	Id      int    `form:"Id" binding:"gt=0"`
	Title   string `form:"Title" binding:"gt=0"`   //字符串长度大于0
	Article string `form:"Article" binding:"gt=0"` //字符串长度大于0
}

type ReqGetAccount struct {
	Account  string `form:"Account"`  //帳號
	Password string `form:"Password"` //密碼
}
