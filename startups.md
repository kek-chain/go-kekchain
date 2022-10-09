
For a hidden node such as exchange hot wallet add the nodiscover command:
```
--nodiscover 
```

For a private RPC enable firewall, such as UFW and then deny or reject calls to the http port:

```
./geth --datadir walletDirName1 --port 12345 --ipcpath=/geth12345.ipc --http.api eth,net,txpool,web3 --http --http.vhosts '*' --http.port 8551 --http.corsdomain '*' --http.addr 127.0.0.1 --syncmode='full' --gcmode='archive' --txlookuplimit=0 --nodiscover --bootnodes="enr:-KO4QKKNUDs0hOAbGFmrqSwIkc08SvLJSyRV9Ig40edh-DxDNH0vc3DcK5w9DI2Ww-ornc2lMKjj_NessjDnnEj-xw6GAYNFQdmcg2V0aMfGhBbi2EuAgmlkgnY0gmlwhCby7PeJc2VjcDI1NmsxoQKMBMD4LxidDE9A5n6d_PwcZFtUThMTdQKa2cvxequHqYRzbmFwwIN0Y3CC-s6DdWRwgvrO" --kekistan console
```
As shown above the http port is 8551, block outside access to this port with firewall such as UFW.

_____________________________________

To run multiple nodes in parallel 
for the open node, don't open personal http, and do not unlock wallet, usually best with an alternate wallet which is EMPTY (no balance)
open node: serverA (open NAT)
node 1,2
we can also run multiple nodes, and open the nat, change ENTER_SERVER_IP_ADDRESS_HERE with the localhost server IP address.(see below)
```
./geth --datadir walletDirName2 --port 2345 --ipcpath=/geth2345.ipc --syncmode='full' --bootnodes="enr:-KO4QKKNUDs0hOAbGFmrqSwIkc08SvLJSyRV9Ig40edh-DxDNH0vc3DcK5w9DI2Ww-ornc2lMKjj_NessjDnnEj-xw6GAYNFQdmcg2V0aMfGhBbi2EuAgmlkgnY0gmlwhCby7PeJc2VjcDI1NmsxoQKMBMD4LxidDE9A5n6d_PwcZFtUThMTdQKa2cvxequHqYRzbmFwwIN0Y3CC-s6DdWRwgvrO" --kekistan console


./geth --datadir walletDirName3 --nat extip:ENTER_SERVER_IP_ADDRESS_HERE --port 5432 --ipcpath=/geth5432.ipc --syncmode='full' --bootnodes="enr:-KO4QKKNUDs0hOAbGFmrqSwIkc08SvLJSyRV9Ig40edh-DxDNH0vc3DcK5w9DI2Ww-ornc2lMKjj_NessjDnnEj-xw6GAYNFQdmcg2V0aMfGhBbi2EuAgmlkgnY0gmlwhCby7PeJc2VjcDI1NmsxoQKMBMD4LxidDE9A5n6d_PwcZFtUThMTdQKa2cvxequHqYRzbmFwwIN0Y3CC-s6DdWRwgvrO" --kekistan console
```

For a public RPC these are the commands:
```
./geth --datadir walletDirName1 --port 64206 --ipcpath=/geth420698.ipc --http.api eth,net,txpool,web3 --http --http.vhosts '*' --http.port 8551 --http.corsdomain '*' --http.addr 0.0.0.0 --syncmode='full' --gcmode='archive' --txlookuplimit=0 --bootnodes="enr:-KO4QKKNUDs0hOAbGFmrqSwIkc08SvLJSyRV9Ig40edh-DxDNH0vc3DcK5w9DI2Ww-ornc2lMKjj_NessjDnnEj-xw6GAYNFQdmcg2V0aMfGhBbi2EuAgmlkgnY0gmlwhCby7PeJc2VjcDI1NmsxoQKMBMD4LxidDE9A5n6d_PwcZFtUThMTdQKa2cvxequHqYRzbmFwwIN0Y3CC-s6DdWRwgvrO" --kekistan console
```

On node 1 (with nat open) 
Once inside geth console, take the local bootnode ENR from the admin.nodeInfo command
```
admin.nodeInfo
```

geth console output will contain a JSON object, look for this node's ENR within the JSON object output, looks like the below:'
```
enr: "enr:-KO4QKKNUDs0hOAbGFmrqSwIkc08SvLJSyRV9Ig40edh-DxDNH0vc3DcK5w9DI2Ww-ornc2lMKjj_NessjDnnEj-xw6GAYNFQdmcg2V0aMfGhBbi2EuAgmlkgnY0gmlwhCby7PeJc2VjcDI1NmsxoQKMBMD4LxidDE9A5n6d_PwcZFtUThMTdQKa2cvxequHqYRzbmFwwIN0Y3CC-s6DdWRwgvrO"
```

copy the entire ENR address from node1, and use this ENR as a bootnode for node3 to restart the 1st node with the --bootnodes flag command line argument
```
./geth --datadir walletDirName1 --port 64206 --ipcpath=/geth420698.ipc --http.api eth,net,txpool,web3 --http --http.vhosts '*' --http.port 8551 --http.corsdomain '*' --http.addr 0.0.0.0 --syncmode='full' --gcmode='archive' --txlookuplimit=0 ---bootnodes="enr:-KO4QKKNUDs0hOAbGFmrqSwIkc08SvLJSyRV9Ig40edh-DxDNH0vc3DcK5w9DI2Ww-ornc2lMKjj_NessjDnnEj-xw6GAYNFQdmcg2V0aMfGhBbi2EuAgmlkgnY0gmlwhCby7PeJc2VjcDI1NmsxoQKMBMD4LxidDE9A5n6d_PwcZFtUThMTdQKa2cvxequHqYRzbmFwwIN0Y3CC-s6DdWRwgvrO" --kekistan console
```

repeat the process doing the same step from node1, get the ENR and then copy it, restart node3 with node1 ENR as bootnode.

There you have it, your local geth (kekchain) genesis was initialized with flag --kekistan, and the local chain is now synchronized with the mainnet kekchain, and the seed nodes are constructed to bootnodes from each other (meaning they will reach those nodes for synchronization)
