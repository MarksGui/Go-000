学习笔记

### 1.作业
按照自己的构想，写一个项目满足基本的目录结构和工程，代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

### 2.项目目录
├── README.md
├── api             # rpc文件
│   └── hello
│          ├── hello.pb.go
│          └── hello.proto
├── cmd
│   └── hello      hello 项目
│       ├── main.go
│       ├── wire.go
│       └── wire_gen.go
├── go.mod
├── go.sum
├── configs
│   ├── config.go       # 初始化配置
│    └── config.yaml     # 配置 yaml 文件
├── internal
│   ├── biz         # 业务组装层
│   │   └── hello.go
│   ├── data        # 数据层
│   │   └── hello.go
│   ├── pkg         # 公共库
│   │   └── errcode    # 自定义错误
│   │       ├── common.go  # 通用错误      
│   │       └── errors.go  # 错误定义
│   └── service     # API实现层
│       └── hello.go
└── pkg            # 测试服务API
    └── model # 数据库
          └── init    # 初始化