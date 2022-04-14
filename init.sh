KEY="yourkey"
CHAINID="echelon_3000-3"
MONIKER="Yournodename"
KEYRING="file"
KEYPASSWD=""
KEYALGO="eth_secp256k1"
LOGLEVEL="warn"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# Reinstall daemon
rm -rf ~/.echelond*
make install
# Copy the built binary to the correct location
cp ~/go/bin/echelond /usr/local/bin/echelond

# Set client config
echelond config keyring-backend $KEYRING
echelond config chain-id $CHAINID

# if $KEY exists it should be deleted
yes $KEYPASSWD | echelond keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

# Set moniker and chain-id for Echelon (Moniker can be anything, chain-id must be an integer)
echelond init $MONIKER --chain-id $CHAINID

# Change parameter token denominations to aechelon
cat $HOME/.echelond/config/genesis.json | jq '.app_state["staking"]["params"]["bond_denom"]="aechelon"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json
cat $HOME/.echelond/config/genesis.json | jq '.app_state["crisis"]["constant_fee"]["denom"]="aechelon"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json
cat $HOME/.echelond/config/genesis.json | jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="aechelon"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json
cat $HOME/.echelond/config/genesis.json | jq '.app_state["evm"]["params"]["evm_denom"]="aechelon"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json
cat $HOME/.echelond/config/genesis.json | jq '.app_state["inflation"]["params"]["mint_denom"]="aechelon"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json

# Decrease the block time target spacing 1000ms = 1s
cat $HOME/.echelond/config/genesis.json | jq '.consensus_params["block"]["time_iota_ms"]="1000"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json

# Set gas limit in genesis
cat $HOME/.echelond/config/genesis.json | jq '.consensus_params["block"]["max_gas"]="10000000"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json

# Set claims start time
# node_address=$(echelond keys list | grep  "address: " | cut -c12-)
# current_date=$(date -u +"%Y-%m-%dT%TZ")
# cat $HOME/.echelond/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["airdrop_start_time"]=$current_date' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json

# Set claims records for validator account
amount_to_claim=1000000
# cat $HOME/.echelond/config/genesis.json | jq -r --arg node_address "$node_address" --arg amount_to_claim "$amount_to_claim" '.app_state["claims"]["claims_records"]=[{"initial_claimable_amount":$amount_to_claim, "actions_completed":[false, false, false, false],"address":$node_address}]' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json

# Set claims decay
# cat $HOME/.echelond/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["duration_of_decay"]="1000000s"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json
# cat $HOME/.echelond/config/genesis.json | jq -r --arg current_date "$current_date" '.app_state["claims"]["params"]["duration_until_decay"]="100000s"' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json

# Claim module account:
# 0xA61808Fe40fEb8B3433778BBC2ecECCAA47c8c47 || echelon15cvq3ljql6utxseh0zau9m8ve2j8erz8jy7kzu
cat $HOME/.echelond/config/genesis.json | jq -r --arg amount_to_claim "$amount_to_claim" '.app_state["bank"]["balances"] += [{"address":"echelon15cvq3ljql6utxseh0zau9m8ve2j8erz8jy7kzu","coins":[{"denom":"aechelon", "amount":$amount_to_claim}]}]' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json

# disable produce empty block
if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.echelond/config/config.toml
  else
    sed -i 's/create_empty_blocks = true/create_empty_blocks = false/g' $HOME/.echelond/config/config.toml
fi

if [[ $1 == "pending" ]]; then
  if [[ "$OSTYPE" == "darwin"* ]]; then
      sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.echelond/config/config.toml
      sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.echelond/config/config.toml
      sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.echelond/config/config.toml
      sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.echelond/config/config.toml
      sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.echelond/config/config.toml
      sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.echelond/config/config.toml
      sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.echelond/config/config.toml
      sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.echelond/config/config.toml
      sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.echelond/config/config.toml
  else
      sed -i 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "30s"/g' $HOME/.echelond/config/config.toml
      sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' $HOME/.echelond/config/config.toml
      sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' $HOME/.echelond/config/config.toml
      sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' $HOME/.echelond/config/config.toml
      sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' $HOME/.echelond/config/config.toml
      sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' $HOME/.echelond/config/config.toml
      sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' $HOME/.echelond/config/config.toml
      sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' $HOME/.echelond/config/config.toml
      sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' $HOME/.echelond/config/config.toml
  fi
fi

# Allocate genesis accounts (cosmos formatted addresses) 150 million to evm
echelond add-genesis-account $KEY 150000000000000000000000000aechelon --keyring-backend $KEYRING

# Update total supply with claim values
validators_supply=$(cat $HOME/.echelond/config/genesis.json | jq -r '.app_state["bank"]["supply"][0]["amount"]')
# Bc is required to add this big numbers
# total_supply=$(bc <<< "$amount_to_claim+$validators_supply")
total_supply=150000000000000000001000000 # 1000000
cat $HOME/.echelond/config/genesis.json | jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"][0]["amount"]=$total_supply' > $HOME/.echelond/config/tmp_genesis.json && mv $HOME/.echelond/config/tmp_genesis.json $HOME/.echelond/config/genesis.json

# Sign genesis transaction for validator
echelond gentx $KEY 1000000000000000000000000aechelon --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
echelond collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
echelond validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
echelond start --pruning=nothing $TRACE --log_level $LOGLEVEL --minimum-gas-prices=0.0001aechelon --json-rpc.api eth,txpool,personal,net,debug,web3
