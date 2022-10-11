package helper

import (
	"cloud-disk/core/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func GenerateToken(id int, identity, name string) (string, error) {
	uc := define.UserClaim{
		Id:       id,
		Identity: identity,
		Name:     name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func MailSendCode(emailAddress, code string) error {
	e := email.NewEmail()
	e.From = define.FromMailName + "<" + define.FromMailAddress + ">"
	e.To = []string{emailAddress}
	e.Subject = "验证码"
	e.HTML = []byte("您的验证码为：<h1>" + code + "</h1>")
	err := e.SendWithTLS(define.MailHost+":465", smtp.PlainAuth("", define.FromMailAddress, define.MailPassword, define.MailHost),
		&tls.Config{ServerName: define.MailHost, InsecureSkipVerify: true})
	if err != nil {
		return err
	}
	return nil
}
