## Install go-kekchain-testnet
When done, check [startups.md](https://github.com/lol-chain/go-kekchain-testnet/blob/main/startups.md) 
Installation same as go-ethereum, forked from geth.

Clone source, and Install geth:
1) clone the repository
``` 
git clone https://github.com/lol-chain/go-kekchain-testnet
```

2) set GOPATH to the local path of the repository we just cloned, add /bin 
windows requires quotes wrapped around the path, linux does not allow them...
linux:
```
export GOPATH=~/path/to/go-kekchain-testnet/bin
```
windows: 
```
export GOPATH="/c/path/to/go-kekchain-testnet/bin"
```

steps 3, and 4 starting from the root path of go-kekchain-testnet 

3) use script [mkdirbin.sh](https://github.com/go-electronero/go-kekchain/blob/published/mkdirbin.sh)
OR
manually make a /bin directory, and install with go 
```
mkdir bin && go install ./...
```
The bin/bin directory will contain the executables of geth

4) use the script [cp-genesis.sh](https://github.com/go-electronero/go-kekchain/blob/published/cp-genesis.sh) 
OR 
manually copy testnet-genesis.json from root to the bin/bin directory, use geth tochange directory to bin/bin & create a keystore with geth
```
cp testnet-genesis.json bin/bin/testnet-genesis.json && cd bin/bin && ./geth account new --datadir walletDirName
```

5)  initialize the testnet-genesis.json
```
./geth init --datadir walletDirName testnet-genesis.json
```

Run geth to sync the testnet chain (I hard coded the bootnodes for auto sync.
```
./geth --datadir walletDirName
```

If sync is too slow, or to speed up the sync...
From within the geth console add trusted nodes 
```
admin.addPeer("enode://569320f40d1761120a780303b8be806643508eb6d75ccb63796710e42f2e888f1ef28069df719bd820d46f6f44a22c338a10e16c03df10509517edf4f0326d97@161.97.121.104:36514")
admin.addTrustedPeer("enode://569320f40d1761120a780303b8be806643508eb6d75ccb63796710e42f2e888f1ef28069df719bd820d46f6f44a22c338a10e16c03df10509517edf4f0326d97@161.97.121.104:36514")
admin.addPeer("enode://569320f40d1761120a780303b8be806643508eb6d75ccb63796710e42f2e888f1ef28069df719bd820d46f6f44a22c338a10e16c03df10509517edf4f0326d97@127.0.0.1:64206")
admin.addTrustedPeer("enode://569320f40d1761120a780303b8be806643508eb6d75ccb63796710e42f2e888f1ef28069df719bd820d46f6f44a22c338a10e16c03df10509517edf4f0326d97@127.0.0.1:64206")
admin.addPeer("enode://fb0af07315239a316f6012ad0c851617fb2253555c3b0679feb0e8ae15d2f8f2dc4b0e963fbac146a9bdc778b7147854da70df0c3eb7b40ba8b094d153730fde@161.97.121.84:64206")
admin.addTrustedPeer("enode://fb0af07315239a316f6012ad0c851617fb2253555c3b0679feb0e8ae15d2f8f2dc4b0e963fbac146a9bdc778b7147854da70df0c3eb7b40ba8b094d153730fde@161.97.121.84:64206")
admin.addPeer("enode://fb0af07315239a316f6012ad0c851617fb2253555c3b0679feb0e8ae15d2f8f2dc4b0e963fbac146a9bdc778b7147854da70df0c3eb7b40ba8b094d153730fde@161.97.121.84:60374")
admin.addTrustedPeer("enode://fb0af07315239a316f6012ad0c851617fb2253555c3b0679feb0e8ae15d2f8f2dc4b0e963fbac146a9bdc778b7147854da70df0c3eb7b40ba8b094d153730fde@161.97.121.84:60374")
admin.addPeer("enode://569320f40d1761120a780303b8be806643508eb6d75ccb63796710e42f2e888f1ef28069df719bd820d46f6f44a22c338a10e16c03df10509517edf4f0326d97@161.97.121.104:64206")
admin.addTrustedPeer("enode://569320f40d1761120a780303b8be806643508eb6d75ccb63796710e42f2e888f1ef28069df719bd820d46f6f44a22c338a10e16c03df10509517edf4f0326d97@161.97.121.104:64206")
admin.addPeer("enode://f942f6d4706bdcecb445d7d36ecc31ba3c377219e579b8f6a49675a0d5b58782cf16b82d756c5b0904280f11a459387f7b956747cd94a9ea96c8647a5f7b77ad@161.97.121.84:64206")
admin.addTrustedPeer("enode://f942f6d4706bdcecb445d7d36ecc31ba3c377219e579b8f6a49675a0d5b58782cf16b82d756c5b0904280f11a459387f7b956747cd94a9ea96c8647a5f7b77ad@161.97.121.84:64206")
admin.addPeer("enode://93dff98c642ccbf26f785b67695d86bf247381c124c120fea79576a35807d338e38df2ae0e9c40a4d97efd409e337dc6cc2ed94387d65685ffae1855e85c4929@161.97.121.104:52494")
admin.addTrustedPeer("enode://93dff98c642ccbf26f785b67695d86bf247381c124c120fea79576a35807d338e38df2ae0e9c40a4d97efd409e337dc6cc2ed94387d65685ffae1855e85c4929@161.97.121.104:52494")
```
