### Module & namespace
1. Module  决定如何寻找到这个模块作为依赖
2. Namespace 决定生成类的位置，类似java的package,防止命名冲突
3. 认证 - kinit xxx

### repository
- GitLab仓库地址中的域名是：code.byted.org
- Gerrit仓库地址中的域名是：git.byted.org


### key

alt +（shift） + 上下左右

.
├── build.sh // 用来编译的脚本，一般情况下不需要更改
├── conf // 用来存放配置文件的目录，一般情况下不需要更改
│   └── kitex.yml // 配置文件，可以通过修改里面的值来更改一些配置
├── handler.go // 服务端的业务逻辑都放在这里，这也是我们需要更改和编写的文件
├── kitex_gen // kitex 生成代码存放的目录，不可更改，更改后重新生成会丢失改动
│   ├── base
│   │   ├── base.go   // 根据 IDL 生成的编解码文件
│   │   ├── k-base.go // kitex 专用的一些拓展内容
│   │   └── k-consts.go
│   └── kitex
│       └── example
│           └── item
│               ├── itemservice // kitex 封装代码主要在这里
│               │   ├── client.go
│               │   ├── invoker.go
│               │   ├── itemservice.go
│               │   └── server.go
│               ├── k-consts.go
│               ├── k-kitex_example_item.go // kitex 专用的一些拓展内容
│               └── kitex_example_item.go   // 根据 IDL 生成的编解码文件
├── kx.yml // 如果使用 kx 的话会生成的文件
├── main.go // go 中的 main 函数，一般在这里做一些资源初始化的工作，可以更改
└── script // runtime 标准环境所需的一些启动脚本，一般情况下不需要更改
├── bootstrap.sh
└── settings.py