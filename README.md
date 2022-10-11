# cloud-disk 网盘

> 轻量级云盘系统，基于 go-zero、xorm、go-redis实现

使用到的命令
```shell
go get -u github.com/zeromicro/go-zero 

go get  xorm.io/xorm 
go get github.com/go-sql-driver/mysql  

go install github.com/zeromicro/go-zero/tools/goctl@latest
goctl api new core
 
go get github.com/jordan-wright/email

cd core
goctl api go -api core.api -dir . -style go_zero
go run core.go -f etc/core-api.yaml
```