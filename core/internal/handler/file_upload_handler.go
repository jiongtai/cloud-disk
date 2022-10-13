package handler

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/models"
	"crypto/md5"
	"fmt"
	"net/http"
	"path"

	"cloud-disk/core/internal/logic"
	"cloud-disk/core/internal/svc"
	"cloud-disk/core/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			return
		}
		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))

		// 判断文件是否已存在
		rp := new(models.RepositoryPool)
		get, err := svcCtx.Engine.Where("hash=?", hash).Get(rp)
		if err != nil {
			return
		}
		if get {
			httpx.OkJson(w, &types.FileUploadReply{Identity: rp.Identity})
			return
		}

		// 文件不存在，上传腾讯云
		cosPath, err := helper.CosUpload(r)
		if err != nil {
			return
		}

		// 往logic传递request
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = cosPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
