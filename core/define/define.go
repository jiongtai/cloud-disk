package define

import "github.com/dgrijalva/jwt-go"

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
