## Install go-kekchain
When done, check [startups.md](https://github.com/kek-chain/go-kekchain/blob/main/startups.md) 
Installation same as go-ethereum, forked from geth.

Clone source, and Install geth:
1) clone the repository
``` 
git clone https://github.com/kek-chain/go-kekchain
```

2) set GOPATH to the local path of the repository we just cloned, add /bin 
windows requires quotes wrapped around the path, linux does not allow them...
linux:
```
export GOPATH=~/path/to/go-kekchain/bin
```
windows: 
```
export GOPATH="/c/path/to/go-kekchain/bin"
```

steps 3, and 4 starting from the root path of go-kekchain 

3) use script [mkdirbin.sh](https://github.com/go-electronero/go-kekchain/blob/published/mkdirbin.sh)
OR
manually make a /bin directory, and install with go 
```
mkdir bin && go install ./...
```
The bin/bin directory will contain the executables of geth

4) use the flag --kekistan to initialize the genesis of mainnet kekchain, or for testnet use the flag --kektest
mainnet: 
```
--kekistan
```
testnet: 
```
--kektest
```

Run geth to sync the chain (Hard coded bootnodes for auto sync.)
```
./geth --datadir walletDirName --kekistan
```

Continue to startups.md
