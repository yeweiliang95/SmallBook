# go-user-server
go实现用户登录服务

# 项目目录
```text
-go-user-server
    -cmd  // 项目执行文件夹
        main.go
    -config // 配置文件夹
    -api    // 存放api文件
    -pkg    // 外部可以使用代码
    -internal // 私有应用和库代码
        -routers // 路由文件夹
        -models // 模型定义
        -middleware // 中间件
    -docs   // 文档存放位置
    -vendor // 项目依赖
```

# 初始化项目
```shell
go mod init go-user-server
```

# 安装gin框架
```shell
go get -u github.com/gin-gonic/gin
```

# 安装依赖





