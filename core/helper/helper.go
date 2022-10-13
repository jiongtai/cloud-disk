package helper

import (
	"cloud-disk/core/define"
	"context"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"path"
	"time"
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

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func GetUUID() string {
	return uuid.NewV4().String()
}

func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CosSecretId),
			SecretKey: os.Getenv(define.CosSecretKey),
		},
	})
	formFile, fileHeader, err := r.FormFile("file")
	if err != nil {
		return "", err
	}

	key := "cloud-disk/" + GetUUID() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(context.Background(), key, formFile, nil)
	if err != nil {
		panic(err)
	}

	return define.CosURL + "/" + key, nil
}
