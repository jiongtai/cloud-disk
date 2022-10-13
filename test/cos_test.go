package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestCosUploadFile(t *testing.T) {
	u, _ := url.Parse(define.CosURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CosSecretId),
			SecretKey: os.Getenv(define.CosSecretKey),
		},
	})
	key := "cloud-disk/emotion12.jpg"
	_, _, err := client.Object.Upload(
		context.Background(), key, "./images/emotion.png", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestCosUploadByReader(t *testing.T) {
	u, _ := url.Parse(define.CosURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CosSecretId),
			SecretKey: os.Getenv(define.CosSecretKey),
		},
	})
	key := "cloud-disk/emotion2222.jpg"
	file, err := os.ReadFile("./images/emotion.png")
	if err != nil {
		return
	}
	_, err = client.Object.Put(context.Background(), key, bytes.NewReader(file), nil)
	if err != nil {
		return
	}
}
