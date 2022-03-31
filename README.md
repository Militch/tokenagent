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

## 开发指南

### 代码提交规范

1. 该仓库托管于内网 Gitlab 服务，并且采取主分支保护限制，
故项目成员具有除保护分支以外的推送权限。

2. 分支代码合并到主分支需要提交 `Margin Request` 合并请求。

### 服务持续集成

该仓库已集成 Jenkins 自动化构建服务。当主分支触发例如：推送（Push）事件时，
将会使用 Gitlab Webhook 推送事件至内网 Jenkins 服务，并执行相关构建流程。

自动化构建流程使用 Jenkins Pipeline 编排，
具体构建流程脚本定义在 [Jenkinsfile](./Jenkinsfile) 文件中。

> 关于构建流程脚本文档参与 Jenkins 官方文档
> [https://www.jenkins.io/doc/book/pipeline/](https://www.jenkins.io/doc/book/pipeline/)

### Building From Source

To build from source, Go 1.17  must be
installed on the system. Clone the repo and run
make:
```bash
git clone https://github.com/feiyiban/tokenagent.git
cd tokenagent && make

windows generate:
    tokenagent.exe
linux generate:
    tokenagent
```
start service:
The tokenagent app and tokenagent.yaml configuration file are in the same directory
```bash
    windows:
    tokenagent.exe daemon
    linux:
    tokenagent daemon
```

### Front end docking interface document address
### NFT base on openzeppelin
```bash
  `http JSON-RPC API refer to https://documenter.getpostman.com/view/17454164/UVsTq2KR`
	
	`eth websocket standard API refer to http://cw.hubwiz.com/card/c/infura-api/1/1/2/`
```

### tokenagent.yaml
```bash
  ethmainnet: 
    httpurl: https://mainnet.infura.io/v3/761e7e2b770e4245a67660336ce29bfd
    wsurl: wss://mainnet.infura.io/ws/v3/761e7e2b770e4245a67660336ce29bfd
  ethpolygon: 
    httpurl: https://polygon-mainnet.infura.io/v3/4308b9607b5541779f829c7b28f16866
    wsurl: wss://rpc-mainnet.matic.quiknode.pro
  ethpolygonmumbai: 
    httpurl: https://nd-454-232-011.p2pify.com/6d51099c92d16662778a189b4d0aa0e3
    wsurl: wss://ws-nd-454-232-011.p2pify.com/6d51099c92d16662778a189b4d0aa0e3
  ethrinkeby: 
    httpurl: https://rinkeby.infura.io/v3/a64cc9552107441e8a9c6de95692045f
    wsurl: wss://rinkeby.infura.io/ws/v3/a64cc9552107441e8a9c6de95692045f

    httpurl --> JSON-RPC over HTTPs
    `ethmainnet: https://mainnet.infura.io/v3/761e7e2b770e4245a67660336ce29bfd Corresponding front-end docking dictionary ETH_MAINNET`
    `ethpolygon: https://polygon-mainnet.infura.io/v3/4308b9607b5541779f829c7b28f16866 Corresponding front-end docking dictionary ETH_POLYGON`
    `ethpolygonmumbai: https://nd-454-232-011.p2pify.com/6d51099c92d16662778a189b4d0aa0e3 Corresponding front-end docking dictionary ETH_POLYGON_MUMBAI`
    `ethrinkeby: https://rinkeby.infura.io/v3/a64cc9552107441e8a9c6de95692045f Corresponding front-end docking dictionary ETH_RINKEBY`
	
    wsurl -->JSON-RPC over websockets 
    `ethmainnet: wss://mainnet.infura.io/ws/v3/761e7e2b770e4245a67660336ce29bfd Corresponding front-end ws://localhost:9001/ethmainnet`
    `ethpolygon: wss://rpc-mainnet.matic.quiknode.pro Corresponding front-end ws://localhost:9001/ethpolygon`
    `ethpolygonmumbai: wss://ws-nd-454-232-011.p2pify.com/6d51099c92d16662778a189b4d0aa0e3 ws://localhost:9001/ethpolygonmumbai`
    `ethrinkeby: wss://rinkeby.infura.io/ws/v3/a64cc9552107441e8a9c6de95692045f ws://localhost:9001/ethrinkeby`
		
```
