# MedChain Network

Work in progress (WIP)

Steps performed to date:
 1. copied https://github.com/chainHero/heroes-service to medchain/medchain-networks
 1. changed *almost* all references to hero over to MedChain, medchain, medchain-network, emr, etc.
 1. followed instructions at https://github.com/chainHero/heroes-service-network to rebuild fixtures/
   1. used:
     1. cd fixtures/
     1. sudo rm -fr crypto-config/
     1. cryptogen generate --config=./crypto-config.yaml
     1. FABRIC_CFG_PATH=$PWD configtxgen -profile MedChain -outputBlock ./artifacts/orderer.genesis.block
     1. FABRIC_CFG_PATH=$PWD configtxgen -profile MedChain -outputCreateChannelTx ./artifacts/emr.channel.tx -channelID emr
     1. FABRIC_CFG_PATH=$PWD configtxgen -profile MedChain -outputAnchorPeersUpdate ./artifacts/org1.medchain.anchors.tx -channelID emr -asOrg Org1MedChain
 1. make

 Seems to compile, sets up network, starts web server, accepts POSTs, changes values in blockchain
