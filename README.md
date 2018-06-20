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
     Make orderer genisis 
     1. FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputBlock ./artifacts/orderer.genesis.block

     Initialize permission chain medchain 
     1. FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputCreateChannelTx ./artifacts/medchain.channel.tx -channelID medchain

     Join each org to permission chain
     1. FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/providers.medchain.anchors.tx -channelID medchain -asOrg Providers
     1. configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/storages.medchain.anchors.tx -channelID medchain -asOrg Storages
     1. configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/patients.medchain.anchors.tx -channelID medchain -asOrg Patients
     1. configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/researchers.medchain.anchors.tx -channelID medchain -asOrg Researchers

     Initialize emr channel
     1. configtxgen -profile emr -outputCreateChannelTx ./artifacts/emr.channel.tx -channelID emr

     Add providers and storages to emr 
     1. configtxgen -profile emr -outputAnchorPeersUpdate ./artifacts/providers.emr.anchors.tx -channelID emr -asOrg Providers
     1. configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/storages.emr.anchors.tx -channelID emr -asOrg Storages

     Initialize amr channel 
     1. configtxgen -profile amr -outputCreateChannelTx ./artifacts/amr.channel.tx -channelID amr

     Add providers and researchers to channel
     1. configtxgen -profile amr -outputAnchorPeersUpdate ./artifacts/providers.amr.anchors.tx -channelID amr -asOrg Providers
     1. configtxgen -profile amr -outputAnchorPeersUpdate ./artifacts/researchers.amr.anchors.tx -channelID amr -asOrg Researchers

     Initialize pii channel
     1. configtxgen -profile pii -outputCreateChannelTx ./artifacts/pii.channel.tx -channelID pii

     Add providers and patients to channel 
     1. configtxgen -profile pii -outputAnchorPeersUpdate ./artifacts/providers.pii.anchors.tx -channelID pii -asOrg Providers
     1. configtxgen -profile pii -outputAnchorPeersUpdate ./artifacts/patients.pii.anchors.tx -channelID pii -asOrg Patients

     Initialize directory channel
     1. configtxgen -profile directory -outputCreateChannelTx ./artifacts/directory.channel.tx -channelID directory

     Add providers and patients to channel
     1. configtxgen -profile directory -outputAnchorPeersUpdate ./artifacts/providers.directory.anchors.tx -channelID directory -asOrg Providers
     1. configtxgen -profile directory -outputAnchorPeersUpdate ./artifacts/patients.directory.anchors.tx -channelID directory -asOrg Patients

 1. make

 Seems to compile, sets up network, starts web server, accepts POSTs, changes values in blockchain
