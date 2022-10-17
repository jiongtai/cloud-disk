# cloud-disk 网盘

> 轻量级云盘系统，基于 go-zero、xorm、go-redis实现

### 涉及模块：
> 用户模块：
>   - 密码登陆 done
>   - 邮箱注册 done
>   - 用户详情 done
> 
> 存储池模块：
>   - 中心存储池资源管理 done
>   - 文件上传 done
>   - 个人存储池资源管理 done
>   - 文件关联存储 done
>   - 文件列表 done
>   - 文件名修改 done
>   - 文件夹创建 done
>   - 文件删除 done
>   - 文件移动
> 
> 文件共享模块：
>   - 文件分享

### 使用到的库以及安装命令：
```shell
// go-zero 库
go get -u github.com/zeromicro/go-zero 

// xorm 库
go get xorm.io/xorm 
go get github.com/go-sql-driver/mysql  

// goctl 工具包
go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl api new core
 
// email 库 
go get github.com/jordan-wright/email

// redis 库
go get github.com/go-redis/redis/v8

// uuid 库
go get github.com/satori/go.uuid

// 腾讯云 COS SDK 包
go get -u github.com/tencentyun/cos-go-sdk-v5

cd core
// 使用工具生成相关路由的 handler 跟 logic 代码
goctl api go -api core.api -dir . -style go_zero
// 运行项目
go run core.go -f etc/core-api.yaml
```