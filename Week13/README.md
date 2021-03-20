## 毕业总结

​	13周的学习很快就结束了，在这个过程中收获很多，学习就是为了应用。根据毛老师的讲解，我把学到知识应用到了自己的工作中，解决了项目中的很多痛点。在此，感谢毛老师的精彩讲解。

### 一、公司项目改造

#### 1.工程化改造

一开始公司项目文件目录如下：

```
.
├── web
│   ├── controller // 控制层
│   ├── service    // 业务逻辑层
|   ├── dao        // 数据持久层
│   ├── middleware // 中间件
│   ├── model      // 数据模型定义
│   └── router     // 路由
├── conf           // 配置相关
│   └── app.conf 
├── consts         // 常量定义
├── errors         // 错误定义
├── go.mod
├── test           // 测试
└── main.go        // 入口
```

##### 存在的问题：

1.main.go 文件里各种初始化代码(eg:数据库、redis、配置、日志、rpc等)，可读性特别差；

2.错误定义模块命令不规范、同意错误码对应多个定义；

3.数据持久层代码每个人定义都不一样，没有抽象接口，无法mock代码写单元测试；

##### 改造后：

```
.
├── cmd
│   ├── ycl           // YCL项目
│       ├── main.go   // 入口文件
│       ├── wire_gen.go // wire 生成代码
│       └── wire.go   // 依赖注入
│   
├── config            // 配置
├── go.mod
├── internal
|   ├── errors        // 错误定义
│   ├── pkg           // 内部公共包 
│   │   ├── util      // 工具         
│   │   └── test      // 测试
│   └── server
│       ├── ctrl      // 控制层
│       ├── repo      // 数据持久层
│       ├── service   // 业务逻辑层    
│       └── entity    // 数据实体cd ..
└── third_party       // 第三方文件
```

#### 二、错误处理

golang 的错误处理一直是被外界诟病的一个问题，代码中充斥着各种重复的错误处理代码 `if else` ，这使得很多用惯了 `try catch` 的人觉得这种方式即为头疼。因为这是官方对错误处理的设计理念，使得目前并没有太好的方法解决这个问题。

```go
if err != nil {
	log.Printf("Err:%v\n", err)
	return err
}
```

通过训练营，我了解到了一个包: `github.com/pkg/errors`，通过这个包我们能很好解决 go 错误处理问题。

```go
func (p *UserRepository) GetInfo() (*UserEntity, error) {
	// exec query SQL
	err := sql.ErrNoRows
	if err != nil {
    // 包装一层实际的 err，然后层层的向外抛出
		return nil, errors.Wrap(err, "query user failed")
	}
	return &UserEntity{}, nil
}
```

最终只需要在控制层处理这个错误即可：

```go
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	info, err := NewUserService().GetInfo()
	if err != nil {
		log.Printf("%+v\n", err) // 打印具体的错误堆栈，包含了堆栈信息
		if errors.Is(err, ErrNotFound) {
			fmt.Fprintf(os.Stdout, "%s", "查无记录")
			return
		}
		fmt.Fprintf(os.Stdout, "%s", "查询失败")
		return
	}

	fmt.Fprintf(os.Stdout, "%+v\n", info)
}
```

通过这种方式，我们可以选择包装错误、写入错误信息，最后只需要在入口处理错误，打印日志即可。代码看起来非常干净、处理方式极为优雅。

#### 三、goroutine 生命周期管理

以前，代码里大家要写一个异步处理函数时都会这样做：

```go
go func(){
	// 业务逻辑
}()
```

这种方式被称为：野生的 goroutine，我们完全没有管理这个 goroutine 的声明周期，如果代码里面造成 panic 还使得整个程序崩溃。

现在把 `errgroup` 包里的代码拷贝处理修改，管理 golang 的生命周期：

```go
var g errgroup.Group
g.Go()
g.Go()
g.Wait()
```

`errgroup` 可以使用 `context` 的方式管理 goroutine 声明周期，同时适用 defer revocer 捕获 panic ，防止意外情况发生，大大提升了代码的可靠性。

### 二、自我总结

​	说实话，一开始我真的是抱着试试的想法报班的，因为确实不知道质量如何，但是考虑到毛老师的口碑以及极客时间以往课程质量，我还是下定了决心报班。

​	通过13周的学习，真的是让我收获颇丰。对于在小公司的我们，研发流程不规范、制度不完善、业务场景单一、人际关系简单，而这节课说的是golang培训，实际上却是涵盖了各个方面的知识，毛老师真的是倾囊相授。从golang知识到架构演进，从底层runtime到职场经验，大家都觉得超值。

​	由于工作较忙，目前视频只是看了一遍，而且没有做太多的总结笔记。因为毕业后有几个月的视频有效期，我准备再次重新复习一遍，然后带着问题学习，学习完后立马写总结笔记，同时把能应用到公司的技术都实际操作一番。最会再查阅相关相关资料，补充笔记内容。我相信通过这样的方法，能使得我最大程度的理解课程内容。

​	最后衷心的感谢毛老师的讲解，感谢极客时间、各位助教老师、班主任的辛苦付出。最后祝你们工作顺利，身体健康。