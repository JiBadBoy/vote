
```
├─conf              项目配置目录
├─models            数据结构目录（表结构）
│  └─models.go      数据库初始化文件（我用的sqlserv,你要换成mysql）
├─routers           路由目录（resultful——api）
│  └─api
│      └─v1         v1版本的路由（后期升级，v2、v3....）
└─utils             工具包
    ├─app           封装系统级的函数
    │  └─request.go 请求验证等
    │  └─response.go    统一返回函数           
    ├─e             定义错误消息常量
    │  └─code.go    消息code
    │  └─msg.go     消息内容
    ├─file          一些文件操作常用函数
    ├─logging       打点日志
    └─setting       初始化配置
    pagination.go   分页
    main.go         主程序
```

