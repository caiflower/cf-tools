# cf-tools

​    本项目是为了使用命令行一键生成web开发架子的。

**安装cf命令行**

```bash
git clone https://github.com/caiflower/cf-tools.git
cd cf-tools/build
make install
chmod +x "cf"
echo "sudo mv ./cf /usr/local/bin"
sudo mv ./cf /usr/local/bin
```

**cf**

```bash
# cf create --help
path:

Usage:
  cf create [flags]

Flags:
      --branch string       demo project git branch (default "v1.0.0")
      --git-origin string   project git origin address
  -h, --help                help for create
      --module string       module name
      --url string          demo project git address (default "https://github.com/caiflower/demo-api.git")
```

**创建一个test-demo应用**

```bash
# cf create --module github.com/caiflower/test-demo
git clone https://github.com/caiflower/demo-api.git -b v1.0.0
正克隆到 'demo-api'...
rename project
git mod tidy
git init
提示：使用 'master' 作为初始分支的名称。这个默认分支名称可能会更改。要在新仓库中
提示：配置使用初始分支名，并消除这条警告，请执行：
提示：
提示：	git config --global init.defaultBranch <名称>
提示：
提示：除了 'master' 之外，通常选定的名字有 'main'、'trunk' 和 'development'。
提示：可以通过以下命令重命名刚创建的分支：
提示：
提示：	git branch -m <name>
已初始化空的 Git 仓库于 /Users/lijinlong/test-demo/.git/
git add .
git commit -m 'init'
[master（根提交） dbe7a50] init
 18 files changed, 508 insertions(+)
 create mode 100644 .gitignore
 create mode 100644 LICENSE
 create mode 100644 constants/config.go
 create mode 100644 controller/v1/base/hello.go
 create mode 100644 dao/test.go
 create mode 100644 etc/config.yaml
 create mode 100644 etc/default.yaml
 create mode 100644 go.mod
 create mode 100644 go.sum
 create mode 100644 main.go
 create mode 100644 model/api/base.go
 create mode 100644 model/api/base/hello.go
 create mode 100644 model/bean/base.go
 create mode 100644 model/bean/test.go
 create mode 100644 service/caller/default.go
 create mode 100644 web/action.go
 create mode 100644 web/restful.go
 create mode 100644 web/server.go
```

![image-20250302205020429](/Users/lijinlong/workspace/cf-tools/images/image-20250302205020429.png)

**应用启动**

![image-20250302205105429](/Users/lijinlong/workspace/cf-tools/images/image-20250302205105429.png)

![image-20250302205302587](/Users/lijinlong/workspace/cf-tools/images/image-20250302205302587.png)

或者使用以下命令行发送http请求

```bash
curl --location '127.0.0.1:8080/v1/base/helloworld?Action=SayHelloWorld'
```



