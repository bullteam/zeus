# zeus 宙斯权限后台
<img src="./docs/images/logo.png" height=145></img>

[![golang](https://img.shields.io/badge/golang-1.12.1-green.svg?style=plastic)](https://www.golang.org/)
[![casbin](https://img.shields.io/badge/casbin-1.8.1-brightgreen.svg?style=plastic)](https://github.com/casbin/casbin)

#### 项目介绍
> `Zeus 宙斯`权限后台，为企业提供统一后台权限管理私有化Sass云服务。    
> - 项目使用`golang beego`框架开发，用`jwt + casbin`做权限管理,提供OAuth2.0 的Restful Api 接口。
> - 为企业后台系统提供统一登陆鉴权、菜单管理、权限管理、组织架构管理、员工管理、配置中心、日志管理等。
> - 支持企业微信、钉钉登陆和同步企业组织架构。
> - 统一管理员工入离职，强化权限审批流程化。
> - 打通开源软件、付费Sass软件，企业内部开发系统等，包括不限于jenkis、jira、gitlab、confluence、禅道、企业邮箱、OA、CRM、财务软件、企业Sass云服务等内外部系统，解决企业多个软件和平台账号不同步的痛点。     
> - `打造统一开放平台生态标准，为企业引进外部系统不再困难。`

更多请进入官网介绍[公牛开源战队](http://www.bullteam.cn) 以及详细的[开发文档指南](http://doc.bullteam.cn)
## Features （目前实现功能）
- 登录/登出
- 权限管理
    - 用户管理(人员管理)
    - 角色管理(功能权限管理)
    - 部门管理
    - 项目管理
    - 菜单管理
    - 数据权限管理
- 个人帐户
    - 第三方登陆（钉钉）
    - 安全设置（[Google 2FA 二次验证](http://www.ruanyifeng.com/blog/2017/11/2fa-tutorial.html)）

## Roadmap （计划实现）
- 组织架构管理(同步钉钉)
- 安全风控
- 操作日志监控
    - 登陆日志
    - 异常登陆
    - 操作日志
- 页面管理
    - 页面配置管理
- 配置中心
- 应用中心 （开放平台）
- 个人帐户
    - 手机验证
    - 邮箱验证
- 增加支持企业微信、微信、Github、Gmail、QQ等登陆
- 登陆授权（OAuth 2.0、Ldap、SAML2.0、Cas、阿里云RAM、AWS IAM、腾讯云CAM、华为云IAM等）
- 打通Worklite、Teambition、Github、墨刀、Tapd 等Sass 服务
- 打通jenkis、jira、gitlab、confluence、禅道等开源软件
  

# Docker 部署
可参考 [Docker Documentation][2] 或者直接看[官方文档][1]

本项目参考，可以一键部署该项目 [docker-composer 部署脚本](http://github.com/bullteam/zeus-deploy)

# 架构
<img src="./docs/images/arch.png" height=920></img>

# 数据库E-R图
<img src="./docs/images/dber.png" height=376></img>

### 快速开始
需要golang 1.11+ 编译环境,设置git clone 权限
````
git clone git@github.com:bullteam/zeus.git
export GOPROXY=https://goproxy.io
export GO111MODULE=on
go build -o zeus
./zeus start -c ./conf

````
# 数据移值

```bash
# 执行 sql 语句
mysql> source ./install/install.sql;
```

## Git 工作流

[Git 协作工作流](docs/Cooperation.md)

# openssl jwt 密钥生成
[openssl jwt 密钥](docs/GenrsaKey.md)

# 演示 Demo
* [admin.bullteam.cn](http://admin.bullteam.cn)  账号 admin  密码  123456   （为了防止恶意使用、系统将不定时重置，请各位客官尽情享用）
  
# 接入权限系统 client demo
* [python-client](https://github.com/bullteam/zeusclient-python)
* [php-client](https://github.com/bullteam/zeusclient-php)
* [java-client](https://github.com/bullteam/zeusclient-java)
* [go-client](https://github.com/bullteam/zeusclient-go)
# WebUI
* [官方](https://github.com/bullteam/zeus-ui)
# API文档
API 开发文档如下：
* [POSTMAN](https://documenter.getpostman.com/view/159835/S1LyTSN3 )

[1]: https://docs.docker.com/ "Docker Documentation"
[2]: https://github.com/yeasy/docker_practice "docker_practice"

## 开发者

* [wutongci](http://github.com/wutongci)
* [funlake](https://github.com/funlake)
* [Hyman](https://github.com/zhengcog)
* [severHo](https://github.com/qq330967496)

更多请进入我们的官网了解我们  [公牛开源战队](http://www.bullteam.cn)

欢迎各路开发者加入或者疑问加入讨论群，请加我微信,说明加入群原因 `zeus 开源交流`

<img src="./docs/images/wx.jpg" height=430></img>

## 相关截图

<img src="./docs/images/screenshot1.png"></img>
<img src="./docs/images/screenshot2.png"></img>
<img src="./docs/images/screenshot3.png"></img>