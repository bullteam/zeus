# api-auth

#### 项目介绍
后台接口，使用golang beego框架开发，用jwt+casbin做权限管理。

#### 功能模块
##### 登录
```
curl -d "username=xxx&password=yyy" /login
```
返回
```
{
  code : 0,
  msg :　"success",
  data : {
    access_token : "xxxxxxxxxxxxxxx"
  }
}
```

##### 请求带令牌
请求带上Authorization头
```
curl -H "Authorization: Bearer [登录获取的令牌]" /changepwd
```

> ps : beego项目的具体controller如果需要jwt验证，只需`继承`TokenCheckController即可,如
```
type ChangepasswdController struct {
	TokenCheckController
}
```
