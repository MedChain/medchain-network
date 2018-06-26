cryptogen generate --config=./crypto-config.yaml
FABRIC_CFG_PATH=$PWD configtxgen -profile MedChain -outputBlock ./artifacts/orderer.genesis.block
FABRIC_CFG_PATH=$PWD configtxgen -profile MedChain -outputCreateChannelTx ./artifacts/emr.channel.tx -channelID emr
FABRIC_CFG_PATH=$PWD configtxgen -profile MedChain -outputAnchorPeersUpdate ./artifacts/providers.medchain.anchors.tx -channelID emr -asOrg Providers


