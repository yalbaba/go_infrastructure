# infrastructure

#### 项目介绍
基础框架

#### 软件架构
1. 所有的业务服务依赖此仓库

2. 文件说明
```text
/component  常用组件, 按需添加
/config     配置实例,按需添加
/logger     日志工具,不要修改
/protos     grpc 协议文件,根据业务需求修改
/server     服务核心代码,不要修改
/models     数据库模型
/utils      工具代码库

```



#### 安装教程

1. xxxx
2. xxxx
3. xxxx

#### 使用说明

1. xxxx
2. xxxx
3. xxxx

#### 参与开发
1. 新功能开发,需要从`master` 分支拉取 `feature/xxx` 分支
2. `feature/xxx` 分支开发完成,提交`pr` ,合并到`develop`分支测试
3. 测试通过后 `feature/xxx` 分支合并到 `master` 分支
4. `develop` 不能直接合并到 `master` 分支