package admin

import (
	"gindemo02/controllers/middleware"
	"gindemo02/dto"
	"gindemo02/models"
	"gindemo02/util"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"strconv"
	"time"
)

const (
	TOKEN_PREFIX = "dual_token_"
	TOKEN_EXPIRE = 7 * 24 * time.Hour //一次登录7天有效
)

type LoginResponse struct {
	Code  int    `json:"code"` //前后端分离，前端根据code向用户展示对应的话术。如果需要改话术，后端代码不用动
	Msg   string `json:"msg"`  //msg仅用于研发人员内部排查问题，不会展示给用户
	Uid   int    `json:"uid"`
	Token string `json:"token"`
}

type LoginController struct {
	BaseController
}

func (this LoginController) Login(ctx *gin.Context) {

	req := dto.ReqGetAccount{}

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.String(http.StatusBadRequest, "invalid parameter")
		return
	}

	// 驗證帳密是否正確
	userModel := models.User{}
	user, err := userModel.GetUsersByAccAndPwd(req.Account, util.Md5(req.Password))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, LoginResponse{5, "token生成失败", 0, ""})
		return
	}

	//登录成功,给前端返回一个Set-Cookie
	//uid := user.ID
	header := util.DefautHeader
	payload := util.JwtPayload{ //payload以明文形式编码在token中，server用自己的密钥可以校验该信息是否被篡改过
		Issue:       "blog",
		IssueAt:     time.Now().Unix(),                                       //因为每次的IssueAt不同，所以每次生成的token也不同
		Expiration:  time.Now().Add(TOKEN_EXPIRE).Add(24 * time.Hour).Unix(), //(7+1)天后过期，需要重新登录，假设24小时内用户肯定要重启浏览器
		UserDefined: map[string]any{middleware.UID_IN_TOKEN: user.ID},        //用户自定义字段。如果token里包含敏感信息，请结合https使用
	}
	jwtKey := util.ConfigMap["key"]["jwtKey"]
	if token, _ := util.GenJWT(header, payload, jwtKey); err != nil {
		util.LogRus.Errorf("生成token失败: %s", err)
		ctx.JSON(http.StatusInternalServerError, LoginResponse{5, "token生成失败", 0, ""})
		return
	} else {
		refreshToken := strconv.Itoa(user.ID) //生成长度为20的随机字符串，作为refresh_token
		SetToken(refreshToken, token)         //把<refreshToken, authToken>写入redis
		//response header里会有一条 Set-Cookie: auth_token=xxx; other_key=other_value，浏览器后续请求会自动把同域名下的cookie再放到request header里来，即request header里会有一条Cookie: auth_token=xxx; other_key=other_value
		ctx.SetCookie("refresh_token", refreshToken, //注意：受cookie本身的限制，这里的token不能超过4K
			int(TOKEN_EXPIRE.Seconds()), //maxAge，cookie的有效时间，时间单位秒。如果不设置过期时间，默认情况下关闭浏览器后cookie被删除
			"/",                         //path，cookie存放目录
			"localhost",                 //cookie从属的域名,不区分协议和端口。如果不指定domain则默认为本host(如b.a.com)，如果指定的domain是一级域名(如a.com)，则二级域名(b.a.com)下也可以访问
			false,                       //是否只能通过https访问
			true,                        //是否允许别人通过js获取自己的cookie，设为false防止XSS攻击
		)
		ctx.JSON(http.StatusOK, LoginResponse{0, "success", user.ID, token})
		return
	}
}

// get auth_token by refresh_token
func (this LoginController) GetAuthToken(ctx *gin.Context) {
	refreshToken := ctx.Query("refresh_token")
	authToken := GetToken(refreshToken)
	ctx.String(http.StatusOK, authToken)
}

// 把<refreshToken, authToken>写入redis
func SetToken(refreshToken, authToken string) {
	if err := models.REDIS.Set(TOKEN_PREFIX+refreshToken, authToken, TOKEN_EXPIRE).Err(); err != nil { //7天之后就拿不到authToken了
		util.LogRus.Errorf("write token pair(%s, %s) to redis failed: %s", refreshToken, authToken, err)
	}
}

// 根据refreshToken获取authToken
func GetToken(refreshToken string) (authToken string) {
	var err error
	if authToken, err = models.REDIS.Get(TOKEN_PREFIX + refreshToken).Result(); err != nil {
		if err != redis.Nil {
			util.LogRus.Errorf("get auth token of refresh token %s failed: %s", refreshToken, err)
		}
	}
	return
}
