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

登陆私有镜像仓库后，通过以下命令即可下载并运行程序

```bash
docker run --name tokenagent \
-p 9001:9001 
registry.cn-hangzhou.aliyuncs.com/dsyun/tokenagent
```

> 关于 `docker run` 指令参数说明参阅[官方参考手册](https://docs.docker.com/engine/reference/commandline/run/)

以上示例部分参数说明：

* -p 9001:9001 该参数将在主机暴露 9001 端口并映射至容器内部 9001 端口




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
