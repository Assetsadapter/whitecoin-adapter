# whitecoin-adapter

## 项目依赖库

- [openwallet](https://github.com/blocktree/openwallet.git)

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf目录，新建XWC.ini文件，编辑如下内容：

```ini
# is enable scanner block chain
isScan = false
# core api url
ServerAPI = "http:node.core"
# Wallet api url
WalletAPI = "http:wallet"
# ChainID
ChainID = "ea1ecf2d8a22d5894280aca2327423f42226e0ecf656f4869972c1c83b6f2a63"
# Fix XWC Required Fee
FixFees = 100000
```