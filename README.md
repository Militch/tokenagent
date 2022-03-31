# tokenagent
Provide multiple chain access services

Currently only the following links are supported:
```bash
    ETH_MAINNET
    ETH_POLYGON
    ETH_POLYGON_MUMBAI
    ETH_RINKEBY
```

## 安装指引

### 使用 Docker

发行版镜像仓库基于阿里云容器镜像服务，并且该发行版镜像属于私有镜像，
所以在使用 `docker` 部署安装时需要使用账号密码凭证登陆到阿里云镜像仓库。

```bash
docker login --username=<user_name> registry.cn-hangzhou.aliyuncs.com
```

登陆私有镜像仓库后，通过以下命令即可下载并运行程序：

```bash
docker run --rm --name tokenagent \
-p 9001:9001 \
registry.cn-hangzhou.aliyuncs.com/dsyun/tokenagent
```

> 关于 `docker run` 指令参数说明参阅[官方参考手册](https://docs.docker.com/engine/reference/commandline/run/)

容器内提供了默认的运行时配置文件 (/etc/tokenagent/config.yml)，
如需自定义配置文件使用 `-v` 参数覆盖替换容器内配置文件：

```bash
docker run --rm --name tokenagent \
-p 9001:9001 \
-v example.yml:/etc/tokenagent/config.yml \
registry.cn-hangzhou.aliyuncs.com/dsyun/tokenagent
```

### 使用 Docker Compose

> 关于 Docker Compose 服务编排工具的安装步骤参阅[官方安装手册](https://docs.docker.com/compose/install/)

本仓库内提供了示例的 [docker-compose.yml](./docker-compose.yml) 文件，
将文件复制到部署环境内执行以下命令：

```bash
docker-compose up
```

使用 `-d` 参数可以将服务置于后台运行

```bash
docker-compose up -d
```

> 更多 Docker Compose 使用方法参阅[官方参考手册](https://docs.docker.com/compose/reference/)

### 使用源码编译构建

#### 系统环境准备

使用源码构建时需要安装一些必要的环境依赖，以下列举常见系统的安装步骤：

Debian:

```bash
sudo apt install build-essential
```

Golang 环境安装：

```bash
curl -L https://golang.org/dl/go1.16.4.linux-amd64.tar.gz | sudo tar -xz -C /usr/local
```

设置系统环境变量

```bash
echo 'export PATH=/usr/local/go/bin:$PATH' >> ~/.bashrc && source ~/.bashrc
```

#### 编译安装

将源码下载到本地环境中, 并切换目录：

```bash
git clone https://git.dsyun.io/tokenagent/tokenagent.git
cd tokenagent
```

执行以下命令编译源码：

```bash
make
```

编译执行成功后将会在项目根路径下的 ./build 目录中生成可执行文件。
至此可以将可执行文件移动到系统环境变量指定的目录中，用以提供全局访问。示例：

```bash
cp build/tokenagent /usr/local/bin
```

## 使用指南

执行以下命令快速启动服务

```bash
tokenagent daemon
```

服务运行时默认使用当前执行路径下的 tokenagent.yaml 文件作为启动配置文件。
可以通过 -C 参数指定启动配置文件：

```bash
tokenagent daemon -C example.yml
```

## 配置文件参考

参阅 [tokenagent.yaml](./tokenagent.yaml)

## 开发指南

### 代码提交规范

1. 该仓库托管于内网 Gitlab 服务，并且采取主分支保护限制，
故项目成员具有除保护分支以外的推送权限。

2. 分支代码合并到主分支需要提交 Merge Requests（合并请求）。

### 服务持续集成

该仓库已集成 Jenkins 自动化构建服务。当主分支触发例如：推送（Push）事件时，
将会使用 Gitlab Webhook 推送事件至内网 Jenkins 服务，并执行相关构建流程。

自动化构建流程使用 Jenkins Pipeline 编排，
具体构建流程脚本定义在 [Jenkinsfile](./Jenkinsfile) 文件中。

> 关于构建流程脚本文档参与 Jenkins 官方文档: 
> [https://www.jenkins.io/doc/book/pipeline/](https://www.jenkins.io/doc/book/pipeline/)

