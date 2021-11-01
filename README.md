目录结构

```
pkg
src             源代码
    common      公共代码
    config      配置文件处理
    controller  控制器(api输入参数检查和跳转)
    db          数据库连接及抽象层
        redis
        mysql
        elasticsearch
    middleware  gin中间件
    mdole       实体数据库映射
    router      路由表
    service     api业务逻辑处理
    utils       非业务逻辑处理
```