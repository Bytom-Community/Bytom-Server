Bytom Wallet
============
[![Build Status](https://travis-ci.org/Bytom/bytom.svg)](https://travis-ci.org/Bytom/bytom)
[![AGPL v3](https://img.shields.io/badge/license-AGPL%20v3-brightgreen.svg)](./LICENSE)

## 架构
![aaa](https://github.com/Bytom-Community/Bytom-Server/blob/master/architecture.jpg)

## 安装编译bytomd
``` bash
$ make bytomd    # build bytomd
```

## 初始化

```bash
$ cd ./cmd/bytomd
$ ./bytomd init --chain_id testnet

```

## 启动sync节点
```
$ ./bytomd node --web.closed --sync_to_db
```

## 启动api节点
```
$ ./bytomd node --web.closed --sync_to_db

```
