# tokenagent
Provide multiple chain access services

Currently only the following links are supported:
```bash
    ETH_MAINNET
    ETH_POLYGON
    ETH_POLYGON_MUMBAI
    ETH_RINKEBY
```

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
    https://documenter.getpostman.com/view/17454164/UVsTq2KR
```

### tokenagent.yaml
```bash
    `ethmainnet: https://mainnet.infura.io/v3/4308b9607b5541779f829c7b28f16866` Corresponding front-end docking dictionary `ETH_MAINNET`
    `ethpolygon: https://polygon-mainnet.infura.io/v3/4308b9607b5541779f829c7b28f16866` Corresponding front-end docking dictionary `ETH_POLYGON`
    `ethpolygonmumbai: https://polygon-mumbai.infura.io/v3/4308b9607b5541779f829c7b28f16866` Corresponding front-end docking dictionary `ETH_POLYGON_MUMBAI`
    `ethrinkeby: https://rinkeby.infura.io/v3/4308b9607b5541779f829c7b28f16866` Corresponding front-end docking dictionary `ETH_RINKEBY`
```
