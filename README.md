# goview
-----------

## 目录结构
-----------

        +-- /private 这个目录存放的是前端打包构建的文件，主要包含模版文件，它并不对外访问，发布的时候，这个目录无需复制到服务器
        | +-- /private/iris iris框架
        | +-- /private/layui layui前端框架
        +-- /public 所有的样式文件(css文件），前端脚本，网站小图标，上传的图片，打包生成的js文件都存放在这个目录
        | +-- /public/images 图片
        | +-- /public/layui layui前端框架
        | +-- /public/scripts js files, such as neffos.js is for websocket
        +-- /src go代码的存放目录
        | +-- /src/access 数据层go代码
        | +-- /src/cache 缓存层代码
        | +-- /src/controller 处理器代码
        | +-- /src/inits 初始化代码
        | +-- /src/middleware 中间件代码
        | +-- /src/model 数据模型
        | +-- /src/net 网络代码
        | +-- /src/routes 路由代码
        | +-- /src/supports 日志等支持代码
        +-- /test 测试代码
        +-- /views go语言的视图模板
        -- main.go 主程序

## 功能测试
-----------------
- 启动selenium服务器

        java -jar selenium-server-standalone-3.9.1.jar

- 启动goview

- 在控制终端运行测试

        cd functional_test
        go test -v
    
## 网站输入
----------------
- 0C遥测帧处理结果

## 用户管理
-----------------
- 学习自项目go-iris

        https://gitee.com/yhm_my/go-iris

- 采用JWT生成token
- 浏览器端记录到localStorage中，并附在请求的Header中[Authorization=bearer <token>]
- todo：暂时短路了用户登录这一步，所有客户端默认分配一个GUEST账号，以及对应的JWT-token
- todo：主要用于通过记录用户的RoleID，来分配相应的权限；如：admin、guest、operator等角色
- todo：重构删除用不到的copy代码