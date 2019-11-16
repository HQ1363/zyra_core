### zyra_core
> 旨在打造一款Go版前后端分离的发布平台，解决中小团队/组织的发布上线问题。

### 技术栈
> 后端
- `Go`
- `Beego`
> 前端
- `Vue + ElementUI`
- `axios`
- `vuex`

## 项目结构
- `gopath所指向目录下新建src/bin/pkg三个目录`
- `进入src目录，clone对应的go项目代码至本地`

### 包依赖管理
- `go mod init`
- `go mod tidy || go get -v ./...`
- `go get xxx`

### 查看API文档
- `bee run -gendoc=true -downdoc=true`
- 浏览器访问`http://对应域名/swagger/`

### Todo-List
- `原Py模块迁移至Go版本`
- `Docker化部署`
- `消息触达模块(dingding/email)`
- `权限管理模块`
- `发布、备份、回滚`
- `日志审计模块`
- `agent实现，废弃原有salt/ansible一套`
- `对接jenkins`
- `对接gitlab`
- `对接K8S(区分可容器化和不可容器化应用)`
- `基础镜像的维护`