# TokenAgent

TokenAgent 是一个区块链 NFT 合约构造以及部署服务，
并且提供多链支持。

当前受支持的区块链：

| 区块链标识 | 描述 |
| :--- | :--- |
| ETH_MAINNET        | 以太坊 MAINNET - 主网       |
| ETH_RINKEBY        | 以太坊 RINKEBY - 测试网      |
| ETH_POLYGON        | 以太坊协议 POLYGON 侧链 - 主网  |
| ETH_POLYGON_MUMBAI | 以太坊协议 POLYGON 侧链 - 测试网 |


## 安装部署指引

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
-v /example.yml:/etc/tokenagent/config.yml \
registry.cn-hangzhou.aliyuncs.com/dsyun/tokenagent
```

### 使用 Docker Compose

关于 Docker Compose 服务编排工具的安装步骤参阅[官方安装手册](https://docs.docker.com/compose/install/)

本仓库内提供了示例的 [docker-compose.yml](./docker-compose.yml) 文件。

新建服务目录并将示例 `docker-compose.yml` 文件复制服务目录下。
执行以下命令快速启动服务：

```bash
sudo docker-compose up
```

使用 `-d` 参数可以将服务置于后台运行

```bash
sudo docker-compose up -d
```

使用 `docker-compose` 可以快速更新服务

```bash
sudo docker-compose pull
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

### 分布式文件系统安装部署

基于业务特性，开放并提供了基础的文件上传功能。
服务本身不提供文件存储能力，为兼顾一些特殊的应用场景，
服务内部代理并提供了文件上传功能，故此需要安装部署分布式文件系统。

> 此功能不推荐在正式环境中使用。应由业务系统根据需求自行提供文件上传服务。

#### IPFS

以下为示例安装步骤，详细安装过程请参阅[IPFS官方文档](https://docs.ipfs.io/install/command-line/)

1. 下载 IPFS 安装程序

```bash
curl https://dist.ipfs.io/go-ipfs/v0.12.0/go-ipfs_v0.12.0_linux-amd64.tar.gz | sudo tar -xz -C /opt
```

2. 安装至系统目录

```bash
sudo ln -s /opt/go-ipfs/ipfs /usr/local/bin
```

3. 检查安装版本

```bash
ipfs --version
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

> 关于构建流程脚本文档参阅 Jenkins 官方文档: 
> [https://www.jenkins.io/doc/book/pipeline/](https://www.jenkins.io/doc/book/pipeline/)

