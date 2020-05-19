# 

This is a template for practicing the normalization of the structure of the construction project
#
这个是gin+gorm前后端分离的后端项目demo，使用gorm访问mysql实现数据的增删改查
#
### 项目结构说明
<pre><code>
├── cmd  程序入口
├── initial 初始化配置的执行函数模块
├── config 配置文件
    ├──example 配置文件编写的例子
├── internal 程序的内部实现
    ├──app 对应的应用软件模块
        ├──blog 博客模块
        ├──shop 商店模块
        ├──user 用户管理模块
    ├──pkg 程序实现的相关包
        ├──models 数据库表模型
        ├──myutil 工具包
    ├──repositories 使用gorm实现增删改查功能的函数仓库
├── response 统一回复格式模块 
├── routers 路由组
├── test 测试接口函数
</code></pre>

#
### 程序启动方式
#
配置文件说明：./configs/example 目录下的yaml文件，按照注释提示进行修改配置文件项目，然后把修改后的yaml文件copy到configs目录下
#
进入cmd目录下，执行

    go run main.go
