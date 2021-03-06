#!/bin/bash

rm -rf ~/.bitsong*

# Initialize the genesis.json file that will help you to bootstrap the network
bitsongd init MyValidator --chain-id=bitsong-localnet

bitsongcli config chain-id bitsong-localnet
bitsongcli config output json
bitsongcli config indent true
bitsongcli config trust-node true
bitsongcli config keyring-backend test

# Change default bond token genesis.json
#sed -i 's/stake/ubtsg/g' ~/.bitsongd/config/genesis.json
sed -i 's/"leveldb"/"goleveldb"/g' ~/.bitsongd/config/config.toml
sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.bitsongd/config/config.toml
sed -i 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.bitsongd/config/config.toml

# Change gov parameters (2 min)
# sed -i 's/"max_deposit_period": "172800000000000"/"max_deposit_period": "240000000000"/g' ~/.bitsongd/config/genesis.json
# sed -i 's/"voting_period": "172800000000000"/"voting_period": "240000000000"/g' ~/.bitsongd/config/genesis.json

# Create a key to hold your validator account
bitsongcli keys add jack
bitsongcli keys add alice

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
bitsongd add-genesis-account jack 100000000000ubtsg --keyring-backend test
bitsongd add-genesis-account alice 100000000000ubtsg --keyring-backend test
bitsongd add-genesis-account bitsong1t0f4ktdux33zlka7930ux87p7zkjjv09kj8xkk 100000000000ubtsg

# Generate the transaction that creates your validator
bitsongd gentx --name jack --amount=10000000ubtsg --keyring-backend test

# Add the generated bonding transaction to the genesis file
bitsongd collect-gentxs
bitsongd validate-genesis

# Now its safe to start `bitsongd`
bitsongd start

