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

var MailPassword = "PJMEKQJRMAKCOAIG"
var FromMailAddress = "jongty@163.com"
var FromMailName = "tyson"
var MailHost = "smtp.163.com"

// CodeLength 验证码长度
var CodeLength = 6
var CodeExpired = time.Second * 300
