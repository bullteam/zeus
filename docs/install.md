# 安装文档

##  第一步：编译golang,注意:需要golang 1.11+ 编译环境,设置git clone 权限
````
git clone git@github.com:bullteam/zeus.git
cd zeus/
export GOPROXY=https://goproxy.io
export GO111MODULE=on
go build -o zeus
chmod 777 ./zeus
./zeus start -c ./conf
````
##  第二步：编译web前端UI
````
cd zeus/web
npm install --registry=https://registry.npm.taobao.org
npm run build:prod
````
##  第三步：安装mysql & redis 数据导入
安装mysql 5.7+  & redis 4.0 +
```bash
# 执行 sql 语句
mysql> source ./install/zeus.sql;
```
##  第四步：安装nginx 增加nginx 配置 
```bash
# admin.bullteam.cn 前端UI nginx 配置
server {
    charset utf-8;
    client_max_body_size 128M;
    listen 80;
    server_name admin.bullteam.cn;
    root /data/src/web/admin.bullteam.cn/dist;
    index index.html;
    location / {
        try_files $uri $uri/ /index.html;
    }
    access_log  /data/log/nginx/admin.bullteam.cn-access.log;
    error_log  /data/log/nginx/admin.bullteam.cn-error.log;
}

upstream api-auth {
    server zeus_api:8082;
}

# api.admin.bullteam.cn 后端接口配置
server {
    listen 80;
    server_name api.admin.bullteam.cn;
    location /(css|js|fonts|img)/ {
        access_log off;
        expires 1d;
        try_files $uri @backend;
    }

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;
        if ($request_method = OPTIONS ) {
                add_header 'Access-Control-Allow-Origin' *;
                add_header 'Access-Control-Allow-Headers' '*';
                add_header 'Access-Control-Allow-Methods' 'GET,POST,OPTIONS,PUT,DEL';
                return 200;
        }
        add_header 'Access-Control-Allow-Origin' *;
        add_header 'Access-Control-Allow-Headers' '*';
        add_header 'Access-Control-Allow-Methods' 'GET,POST,OPTIONS,PUT,DEL';
        proxy_pass http://api-auth;
    }

    access_log  /data/log/nginx/api.admin.bullteam.cn-access.log;
    error_log  /data/log/nginx/api.admin.bullteam.cn-error.log;
}


```

## 第五步：修改hosts
``` bash
# 修改 hosts
127.0.0.1 admin.bullteam.cn;
127.0.0.1 api.admin.bullteam.cn;
```
> 完成五步就可以运行程序，但要用生产，请按照自己需要修改配置的域名，搜索代码中的域名替换即可

> 前端UI 替换域名位置为 zeus/web/config/prod.env.js 修改 ZEUS_ADMIN_URL 为你自己的域名
