#!/bin/bash

KEY="mykey"
CHAINID="echelon_3000-3"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t echelon-datadir.XXXXX)

echo "create and add new keys"
./echelond keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Echelon with moniker=$MONIKER and chain-id=$CHAINID"
./echelond init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./echelond add-genesis-account \
"$(./echelond keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aechelon,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./echelond gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./echelond collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./echelond validate-genesis --home $DATA_DIR

echo "starting echelon node $i in background ..."
./echelond start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started echelon node"
tail -f /dev/null