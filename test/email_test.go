package test

import (
	"cloud-disk/core/define"
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestEmailSend(t *testing.T) {
	e := email.NewEmail()
	e.From = "Tyson Hu <jongty@163.com>"
	e.To = []string{"jongty@foxmail.com"}
	e.Subject = "验证码"
	e.HTML = []byte("<h1>123456</h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "jongty@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{ServerName: "smtp.163.com", InsecureSkipVerify: true})
	if err != nil {
		t.Fatal(err)
	}

}
