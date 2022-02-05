![wlop](./static/imgs/wlop-5se.jpg)

> 图片来自网络，侵删

### Go依赖管理
> 依赖管理的三个阶段： GOPATH、GOVENDOR、go mod

#### 1.GOPATH
- 默认在`~/go`目录
- 多项目使用的依赖版本不一致时，可以内建vendor文件夹，存放对应版本的依赖

#### 2.GOVENDOR
- 每个项目都有自己的vendor目录，存放第三方库
- 大量第三方依赖管理工具：glide、dep、go dep...

#### 3.go mod
> go命令统一管理，不用关心目录结构

- `go mod init <name> 初始化项目go.mod文件`
- `go get [@v...]  安装依赖`
- `go mod tidy 清除未引用的依赖`
- `项目迁移到go mod模式，需要将项目中的go文件进行build，go build ./...`
