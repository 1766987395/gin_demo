# 自用学习代码md不做使用解析
### 命令

*   安装使用 `go get -u github.com/gin-gonic/gin`
*   卸载使用 `go clean -i github.com/gin-gonic/gin`
*   更新依赖 `go mod tidy`
*   查看依赖 `go list -m all` 或直接打开 `go.mod` 文件

### 安装

*   安装命令`go get -u github.com/gin-gonic/gin`
*   安装超时问题

```shell
// 1. 环境变量问题
// 2. 网络问题 设置国内镜像
go env -w GO111MODULE=on 
go env -w GOPROXY=https://goproxy.cn,direct
go mod init YourProjectName
go get -u github.com/gin-gonic/gin
```

*   [参考链接：Go Gin安装解决国内go get 方式安装超时问题](https://zhuanlan.zhihu.com/p/488101096)

***

### helloworld 首次运行使用

*   创建文件`main.go`添加以下代码

```go
package main
import (
	"github.com/gin-gonic/gin" // 引入框架
	"net/http"
)
 
func main() {
	r := gin.Default() // 初始化
	
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "搭建完成")
	}) // get请求方式 指向内置方法
	
	r.Run(":8888")  // 运行 端口号8888 
}
```

*   运行代码 `go run main.go`

*   访问地址 `http://localhost:8888`

*   [参考链接：社区gin文档#快速入门#开始](https://learnku.com/docs/gin-gonic/1.7/quickstart/11354#a3e3b8)

*   [参考链接2：Golang 入门-Gin框架安装及使用](https://blog.csdn.net/qq_34284638/article/details/104944319)

***

### 热编译

*   作用意义：频繁修改代码后可不再编译直接查看修改结果，无需手动重新编译和重新启动应用程序，提升开发效率
*   执行安命令  `go get -u github.com/cosmtrek/air`，网络问题参考第一条
*   执行查看  `air`
*   访问地址  `http://localhost:8888`
*   可修改代码等等热编译完成查看是否生效

***

### 路由

*   形如`main.go`实例

```

import (
	"github.com/gin-gonic/gin" // 引入框架
)

...
// 路由设置
r.GET("/", func(c *gin.Context) {
	c.String(http.StatusOK, "Hello World")
})
```

*   外置方法调取

<!---->

    package main

    import (
    	"net/http"

    	"github.com/gin-gonic/gin"
    )

    func main() {

    	r := gin.Default()

    	r.GET("/", func(c *gin.Context) {

    		c.String(http.StatusOK, "Hello World")

    	})

    	r.GET("/test", test)

    	r.Run(":9000") // 端口号8888

    }

    func test(c *gin.Context) {

    	c.JSON(http.StatusOK, "路由调取外置方法测试！")

    }
    // 可通过，访问地址  `http://localhost:9000/test` 得到方法返回参数

#### 跨文件调取语法

    import ( "gin_demo/router" ) // 项目目录\package包名称

    router.SetupRouter() // package包名称.包下方法名

*   外置路由文件设置
    *   创建文件`router/router.go`输入以下内容
    ```go
    package router

    import (
    	
    	"gin_demo/controller" // 项目目录名+包名

    	"github.com/gin-gonic/gin"

    )

    func SetupRouter() *gin.Engine {

        router := gin.New()

        // 创建路由和路由分组，并添加路由处理程序
        router.GET("/api/test", controller.TestFunc)

        return router
    }
    ```
    *   其中代码引入`controller`包，定义路由`get`方式`/api/test`访问到控制器`TestFunc`方法
    *   定义main.go文件入口
    ```go
    package main

    import (

    	"time"
    	
    	"net/http"
    	
    	"gin_demo/router" // 引入路由包
    	
    )

    func main() {

    	r := router.SetupRouter() // 调用路由方法
    	
    	s := &http.Server{ // 自定义HTTP服务器
    		Addr:           ":9000",
    		Handler:        r,
    		ReadTimeout:    10 * time.Second,
    		WriteTimeout:   10 * time.Second,
    		MaxHeaderBytes: 1 << 20,
    	}
    	
    	s.ListenAndServe()
    	
    }
    ```
    *   可以通过`/api/test`访问到`TestFunc`方法

***

### 控制器

*   创建文件`controller/controller.go`输入以下内容

```go
package controller

import (

	"github.com/gin-gonic/gin"

)

func TestFunc(c *gin.Context) {

	c.JSON(200, gin.H{"msg": "hello world"})

}
```

*   调取方式参考路由

***

### 配置

*   安装包`go get -u gopkg.in/ini.v1`
*   创建文件`config\config.ini` 添加以下内容

```ini
#服务配置
[service]
HttpRort=:9000
# 两种模式 gin.ReleaseMode[release] 和 gin.DebugMode[debug] 使用 gin.SetMode(gin.ReleaseMode)设置 这里配置值就可以
AppMode=debug
ReadTimeOut=60
WriteTimeOut=60
#Mysql数据库配置

[mysql]
Db=mysql
DbHost=127.0.0.1
DBPort=3306
DbUser=root
DbPassWord=root
DbName=your_database
DbPrefix=blog_  #表后缀

#应用配置
[app]
PageSize = 10 #分页
JwtSecret=23347$040412 #jwt中间件生成token使用
```

*   创建文件`config\config.go` 添加以下内容

```go
package conf

import (
    "gopkg.in/ini.v1"
)

var (
    Cfg          *ini.File
    HttpRort    int
    AppMode      string
    ReadTimeOut  int
    WriteTimeOut int
    Db          string
    DbHost      string
    DBPort      string
    DbUser      string
    DbPassWord  string
    DbName      string
    DbPrefix    string
    PageSize    int
    JwtSecret    string
)

func init() {
    var err error
    Cfg, err = ini.Load("config/config.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'config/config.ini': %v", err)
    }
    loadService()
    loadMysql()
    loadApp()
}

func loadService() {
    HttpRort = Cfg.Section("service").Key("HttpRort").String()
    AppMode = Cfg.Section("service").Key("AppMode").String()
    ReadTimeOut, _ = Cfg.Section("service").Key("ReadTimeOut").Int()
    WriteTimeOut, _ = Cfg.Section("service").Key("WriteTimeOut").Int()
}

func loadMysql() {
    Db = Cfg.Section("mysql").Key("Db").String()
    DbHost = Cfg.Section("mysql").Key("DbHost").String()
    DBPort = Cfg.Section("mysql").Key("DBPort").String()
    DbUser = Cfg.Section("mysql").Key("DbUser").String()
    DbPassWord = Cfg.Section("mysql").Key("DbPassWord").String()
    DbName = Cfg.Section("mysql").Key("DbName").String()
    DbPrefix = Cfg.Section("mysql").Key("DbPrefix").String()
}

func loadApp() {
    PageSize, _ = Cfg.Section("app").Key("PageSize").Int()
    JwtSecret = Cfg.Section("app").Key("JwtSecret").String()
}

```

*   引入配置包后使用`config.HttpRort`调用

```
import (
    "gin_demo/config"
)

...

config.HttpRort

```

*   如上配置改造`main.go`

```go
package main

import (

	"time"

	"net/http"

	"gin/router"

	"gin/config"

	"github.com/gin-gonic/gin"

)

func main() {

	// debug
	gin.SetMode(config.AppMode) 

	// 路由
	r := router.SetupRouter()

	s := &http.Server{
		Addr:           config.HttpRort,
		Handler:        r,
		ReadTimeout:    time.Duration(config.ReadTimeOut) * time.Second,
		WriteTimeout:   time.Duration(config.WriteTimeOut) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	
	s.ListenAndServe()
}
```

***

### 增删改查

*   数据库链接
    *   原生连接数据库 `"database/sql"` 访问层
    *   获取`go get github.com/go-sql-driver/mysql`数据库驱动层
        * sqlite `go get -u gorm.io/driver/sqlite`
        * mysql `go get -u gorm.io/driver/mysql`
        * postgres `go get -u gorm.io/driver/postgres`
        * sqlserver `go get -u gorm.io/driver/sqlserver`
    ```
    flowchart LR
        访问层 --> 驱动层 --> 数据库
        
        database/sql --> github.com/go-sql-driver/mysql --> mysql
    ```

    *   创建`db`文件夹，创建`database.go`文件 修改程序为
    ```go
        package db
        
        import (
            "database/sql"
            "log"
            "gin/config"
            _ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动程序
        )
        
        var DB *sql.DB
        
        func InitDB() {
            // 连接数据库
            db, err := sql.Open(config.Db, config.DbUser+":"+config.DbPassWord+"@tcp("+config.DbHost+":"+config.DBPort+")/"+config.DbName)
            if err != nil {
                log.Fatal(err)
            }
            // 测试数据库连接
            err = db.Ping()
            if err != nil {
                log.Fatal(err)
            }
            DB = db
        }
        
        func CloseDB() {
            // 关闭数据库连接
            if DB != nil {
                DB.Close()
            }
        }
    ```
    *   修改入口程序`main.go`
    ```go
    package main

    import (

    	"time"

    	"net/http"

    	"gin/router"

    	"gin/config"

    	"gin/db"

    	"github.com/gin-gonic/gin"

    )

    func main() {

    	// debug
    	gin.SetMode(config.AppMode) 

    	// 初始化数据库连接
    	db.InitDB()

    	// 关闭数据库连接
    	defer db.CloseDB() 

    	// 路由
    	r := router.SetupRouter()
    	
    	s := &http.Server{
    		Addr:           config.HttpRort,
    		Handler:        r,
    		ReadTimeout:    time.Duration(config.ReadTimeOut) * time.Second,
    		WriteTimeout:   time.Duration(config.WriteTimeOut) * time.Second,
    		MaxHeaderBytes: 1 << 20,
    	}
    	
    	s.ListenAndServe()
    }
    ```

#### 接参数语法
* get
```go
// 形如 
router.GET("/api/test", controller.GetUser)
xxx.com/api/test?id=1

// 可是使用 获取到参数id
id := c.Query("id")
```
* post
```go
// 形如 
router.POST("/api/test", controller.GetUser)
xxx.com/api/test
// body 内参数 {id:1}

// 可是使用 获取到参数id
id := c.PostForm("id")
```
* 路由
```go
// 形如 
router.POST("/api/test/:id", controller.GetUser)
xxx.com/api/test/1
// 可是使用 获取到参数id
id := c.Param("id")
```
* json
```go
// 形如 
xxx.com/api/test
// json 格式参数
{ username:xxx, email:xxx }

// 创建一个结构体来表示 JSON 数据的格式
type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
}

func PostJSONData(c *gin.Context) {
    // 使用 ShouldBindJSON 将请求体数据绑定到结构体
    var user User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 可以进一步处理 user 结构体中的字段值
    c.JSON(http.StatusOK, user)
}
```

*   增
    * 参考**查**：修改`sql`语句即可
*   删
    * 参考**查**：修改`sql`语句即可
*   改
    * 参考**查**：修改`sql`语句即可
*   查
    *   在`TestController.go`中创建新方法`GetUsers`
    ```go
    func GetUser(c *gin.Context) {
        rows, err := db.DB.Query("select id,name,phone from users limit 10")

        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

    	defer rows.Close()

        // 创建一个切片用于存储查询结果
        var users []gin.H

        // 迭代查询结果，将数据添加到切片中
        for rows.Next() {
            var id int
            var name, phone string
            err := rows.Scan(&id, &name, &phone)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning database rows"})
            }
            users = append(users, gin.H{"id": id, "name": name, "phone": phone})
        }

        // 返回 JSON 格式的查询结果
        c.JSON(http.StatusOK, users)

        return
    }
    ```
    *   在路由中添加
    ```go
    router.GET("/api/get/users", controller.GetUser)
    ```

***

### orm
* ORM 是对象关系映射（Object-Relational Mapping）的缩写，是一种编程技术，用于在关系型数据库（如 MySQL、PostgreSQL、SQLite 等）和对象导向编程语言（如 Python、Java、Go 等）之间建立映射关系。ORM 具允许开发人员使用面向对象的方式操作数据库，而不需要编写原始的 SQL 查询。ORM 在开发中有很多优点，例如提高了开发效率、减少了编写 SQL 查询的工作量、降低了代码的复杂性，并使数据库访问更加面向对象。
* 这里需要使用到GORM（Go）。
* 安装orm库：`go get -u gorm.io/gorm`
* 安装驱动：
    * sqlite `go get -u gorm.io/driver/sqlite`
    * mysql `go get -u gorm.io/driver/mysql`
    * postgres `go get -u gorm.io/driver/postgres`
    * sqlserver `go get -u gorm.io/driver/sqlserver`
* 

* [参考链接：Golang GORM实战(一)：快速安装与入门 PS:同上方增删改查内容参考](https://juejin.cn/post/7084472581432016910)
***

### 中间件

***

### 解藕分离

***

### 目录结构优化

***

### 事务

***

### 读写分离

***

### 日志

***

### 定时任务

***

### 微服务

***

### 脚手架生成下代码保证编码速度

***

