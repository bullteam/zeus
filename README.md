# zeus 宙斯权限后台
<img src="./docs/images/logo.png" height=145></img>

[![golang](https://img.shields.io/badge/golang-1.12.1-green.svg?style=plastic)](https://www.golang.org/)
[![casbin](https://img.shields.io/badge/casbin-1.8.1-brightgreen.svg?style=plastic)](https://github.com/casbin/casbin)

#### 项目介绍
Zeus 宙斯权限后台，为企业提供统一后台权限管理服务。项目使用golang beego框架开发，用jwt+casbin做权限管理,提供OAuth2.0 的Restful api 接口，为企业后台系统提供
统一菜单管理、权限管理、员工管理、配置中心，同步企业微信、钉钉，同步企业组织架构，打通jenkis、jira、gitlab、企业邮箱、OA、财务软件等内外部系统，解决企业多个
软件和平台账号不同步的痛点。

## Features
- 登录/登出
- 权限管理
    - 用户管理
    - 角色管理
    - 部门管理
    - 项目管理
    - 菜单管理
## Roadmap
- 支持企业微信/钉钉登陆
- 同步企业组织架构和用户
- 风控
- 操作日志监控
- 配置中心
- 应用中心
  
# Docker 部署
可参考 [Docker Documentation][2] 或者直接看[官方文档][1]

本项目参考，可以一键部署该项目 [docker-composer 部署脚本](http://github.com/bullteam/delopy)

# 架构
<img src="./docs/images/arch.png" height=920></img>

### 快速开始
````
export GOPROXY=https://goproxy.io
export GO111MODULE=on
go mod tidy
cd cmd/api-server
go build -o zeus
./zeus

````
# 数据移值

```bash
# 执行 sql 语句
mysql> source ./install/all.sql;
mysql> source ./install/casbin.sql;

# 分别导入到auth、casbin库
```

## Git 工作流

[Git 协作工作流](docs/Cooperation.md)

# openssl jwt 密钥生成
[openssl jwt 密钥](docs/GenrsaKey.md)
# Demo
* [demo](http://admin.bullteam.cn)
# WebUI
* [官方](https://github.com/bullteam/zeus-ui)
# API文档
* [POSTMAN](https://documenter.getpostman.com/view/159835/Rzfjk7Jh)

[1]: https://docs.docker.com/ "Docker Documentation"
[2]: https://github.com/yeasy/docker_practice "docker_practice"

## 协作者

* [wutongci](http://github.com/wutongci)
* [funlake](https://github.com/funlake)
