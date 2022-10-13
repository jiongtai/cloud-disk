package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "cloud-disk-key"

var MailPassword =
var FromMailAddress =
var FromMailName =
var MailHost =

// CodeLength 验证码长度
var CodeLength = 6
var CodeExpired = time.Second * 300

// CosURL 腾讯云COS存储地址
var CosURL =

// CosSecretId 腾讯云COS，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
var CosSecretId =

// CosSecretKey 腾讯云COS，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
var CosSecretKey =
