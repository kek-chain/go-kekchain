To run multiple nodes in parallel 
for the open node, don't open personal http, and do not unlock wallet, usually best with an alternate wallet which is EMPTY (no balance)
open node: serverA (open NAT)
node 1,2
we can also run multiple nodes, and open the nat, change ENTER_SERVER_IP_ADDRESS_HERE with the localhost server IP address.(see below)
```
./geth --datadir walletDirName2 --port 64204 --ipcpath=/geth420298.ipc --syncmode='full' console
./geth --datadir walletDirName3 --nat extip:ENTER_SERVER_IP_ADDRESS_HERE --port 64206 --ipcpath=/geth420699.ipc --syncmode='full' console
```

On alternate server, (serverB) aka kekchain-testnet public RPC
node 3
```
./geth --datadir walletDirName1 --port 64206 --ipcpath=/geth420698.ipc --http.api eth,net,txpool,web3 --http --http.vhosts '*' --http.port 8551 --http.corsdomain '*' --http.addr 0.0.0.0 --syncmode='full' --gcmode='archive' --txlookuplimit=0 console
```

On node 1 (with nat open) 
Once inside geth console, take the local bootnode ENR from the admin.nodeInfo command
```
admin.nodeInfo
```

geth console output will contain a JSON object, look for this node's ENR within the JSON object output, looks like the below:'
```
enr: "enr:-KO4QOjwDhELQPePs2sbvMCKtCxzmR67NKmsWH8-ieqlLHU8c3u4ns3XMhd7gjH7di8Zi6k-LK3O6FbWmo7wZ5nVjd-GAYK51K9cg2V0aMfGhPGOZYqAgmlkgnY0gmlwhKFheVSJc2VjcDI1NmsxoQP5QvbUcGvc7LRF19NuzDG6PDdyGeV5uPaklnWg1bWHgoRzbmFwwIN0Y3CC-s6DdWRwgvrO"
```

copy the entire ENR address from node1, and use this ENR as a bootnode for node3 to restart the 1st node with the --bootnodes flag command line argument
```
./geth --datadir walletDirName1 --port 64206 --ipcpath=/geth420698.ipc --http.api eth,net,txpool,web3 --http --http.vhosts '*' --http.port 8551 --http.corsdomain '*' --http.addr 0.0.0.0 --syncmode='full' --gcmode='archive' --txlookuplimit=0 --bootnodes "enr:-KO4QOjwDhELQPePs2sbvMCKtCxzmR67NKmsWH8-ieqlLHU8c3u4ns3XMhd7gjH7di8Zi6k-LK3O6FbWmo7wZ5nVjd-GAYK51K9cg2V0aMfGhPGOZYqAgmlkgnY0gmlwhKFheVSJc2VjcDI1NmsxoQP5QvbUcGvc7LRF19NuzDG6PDdyGeV5uPaklnWg1bWHgoRzbmFwwIN0Y3CC-s6DdWRwgvrO" console
```

repeat the process doing the same step from node1, get the ENR and then copy it, restart node3 with node1 ENR as bootnode.

There you have it, your local geth (kekchain-testnet) genesis was initialized, and the local chain is now synchronized with the testnet-kekchain, and the seed nodes are constructed to bootnodes from each other (meaning they will reach those nodes for synchronization)
