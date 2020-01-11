## 1.项目结构概览

```json
├── cmd
│   └── server：主服务（程序入口）
├── configs：配置文件目录(包含运行配置参数)
├── docs：文档目录
│   └── swagger：swagger静态文件目录
├── internal：内部应用
│   └── app：主应用目录
│       ├── bll：业务逻辑层接口
│       │   └── impl：业务逻辑层的接口实现
│       ├── config：配置参数（与配置文件一一映射）
│       ├── context：统一上下文
│       ├── errors：统一的错误处理
│       ├── plus：gin的扩展函数库
│       ├── middleware：gin中间件
│       ├── model：存储层接口
│       │   └── impl：存储层接口实现
│       │       └── gorm：基于gorm的存储层实现
│       ├── routers：路由层
│       │   └── api：/api路由模块
│       │       └── ctl：/api路由模块对应的控制器层
│       ├── schema：对象模型
│       └── test：针对接口的单元测试
├── pkg：公共模块
│   ├── auth：认证模块
│   │   └── jwtauth：JWT认证模块实现
│   ├── logger：日志模块
│   └── util：工具库
└── scripts：执行脚本
```





## 2.Gin框架解析

#### 2.1解析结构体

```go
type User struct{
  //required表示该参数是必要参数，如果不传或为空会报错
  //对表单中的Name输入数据进行绑定
  Name string `json:"name" form:"name" binding:"required"`
}
```

#### 2.2 `获取"/user/:name"的name参数`

```go
r.GET("/user/:name", func(c *gin.Context) {
                    name := c.Param("name")
                    c.JSON(http.StatusOK,gin.H{"name":name})
})
```

#### 2.3 `/params?name=lcd&age=12`

```go
  router.GET("/params", func(context *gin.Context) {
            name := context.DefaultQuery("name","defaultName")
            age := context.Query("age")
            context.JSON(http.StatusOK,gin.H{
                "name":name,
                "age":age,
            })
  })
```

#### 2.4 `Multipart/Urlencoded Form`

```go
 r.POST("/form_post", func(c *gin.Context) {
            name := c.PostForm("name")
            nick := c.DefaultPostForm("age", "90")
            c.JSON(200, gin.H{
                "name": name,
                "age":    nick,
            })
 })
```

#### 1.5 `application/Json`

```go
r.POST("/json", func(context *gin.Context) {
        var user User
        context.BindJSON(&user)
        context.JSON(200,gin.H{"data":user})
 })
```

==`建议表单格式或是json格式都是用context.Bind`==