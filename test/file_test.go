package test

import (
	"bytes"
	"cloud-disk/core/define"
	"context"
	"crypto/md5"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"testing"
)

// 分片大小
const chunkSize = 1024 * 1024 // 1MB

// 文件的分片
func TestGenerateChunkFile(t *testing.T) {
	fileInfo, err := os.Stat("test.mp4")
	if err != nil {
		return
	}
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	myFile, err := os.OpenFile("test.mp4", os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	b := make([]byte, chunkSize)
	for i := 0; i < int(chunkNum); i++ {
		// 起始位置
		myFile.Seek(int64(i*chunkSize), 0)
		if chunkSize > fileInfo.Size()-int64(i*chunkSize) {
			b = make([]byte, fileInfo.Size()-int64(i*chunkSize))
		}
		// 写入
		myFile.Read(b)
		f, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return
		}
		f.Write(b)
		f.Close()
	}
	myFile.Close()
}

// 分片文件的合并
func TestMergeChunkFile(t *testing.T) {
	myFile, err := os.OpenFile("testMerge.mp4", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	fileInfo, err := os.Stat("test.mp4")
	if err != nil {
		return
	}
	// 分片的个数
	chunkNum := math.Ceil(float64(fileInfo.Size()) / float64(chunkSize))
	for i := 0; i < int(chunkNum); i++ {
		file, err := os.OpenFile("./"+strconv.Itoa(i)+".chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			return
		}
		all, err := ioutil.ReadAll(file)
		if err != nil {
			return
		}
		myFile.Write(all)
		file.Close()
	}
	myFile.Close()
}

// 文件一致性校验
func TestCheck(t *testing.T) {
	file1, err := os.OpenFile("./test.mp4", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return
	}
	all1, err := ioutil.ReadAll(file1)
	if err != nil {
		return
	}
	file2, err := os.OpenFile("testMerge.mp4", os.O_RDONLY, os.ModePerm)
	if err != nil {
		return
	}
	all2, err := ioutil.ReadAll(file2)
	if err != nil {
		return
	}
	s1 := fmt.Sprintf("%x", md5.Sum(all1))
	s2 := fmt.Sprintf("%x", md5.Sum(all2))
	if s1 != s2 {
		return
	}
}

func TestInitCosChunkUpload(t *testing.T) {
	u, _ := url.Parse(define.CosURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CosSecretId),
			SecretKey: os.Getenv(define.CosSecretKey),
		},
	})
	key := "cloud-disk/example.jpeg"
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		panic(err)
	}
	UploadID := v.UploadID
	fmt.Println(UploadID)
}

func TestCosChunkUpload(t *testing.T) {
	u, _ := url.Parse(define.CosURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CosSecretId),
			SecretKey: os.Getenv(define.CosSecretKey),
		},
	})
	key := "cloud-disk/example.jpeg"
	UploadID := ""
	file, err := os.ReadFile("0.chunk")
	if err != nil {
		return
	}
	resp, err := client.Object.UploadPart(
		context.Background(), key, UploadID, 1, bytes.NewReader(file), nil,
	)
	if err != nil {
		panic(err)
	}
	PartETag := resp.Header.Get("ETag") //Md5值
	fmt.Println(PartETag)
}

func TestCosChunkUploadComplete(t *testing.T) {
	u, _ := url.Parse(define.CosURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(define.CosSecretId),
			SecretKey: os.Getenv(define.CosSecretKey),
		},
	})
	key := "cloud-disk/example.jpeg"
	PartETag, UploadID := "", ""
	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: PartETag},
	)
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), key, UploadID, opt,
	)
	if err != nil {
		return
	}
}
