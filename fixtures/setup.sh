#!/bin/bash
#
# Copyright MedChain, Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#Clean up
rm -rf crypto-config/
rm -rf artifacts/

#Create certificates
cryptogen generate --config=./crypto-config.yaml

mkdir -p artifacts

#Create genesis block
FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputBlock ./artifacts/orderer.genesis.block

#Initialize permission chain medchain
FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputCreateChannelTx ./artifacts/medchain.channel.tx -channelID medchain

#Join each org to permission chain
FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/providers.medchain.anchors.tx -channelID medchain -asOrg Providers
FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/storages.medchain.anchors.tx -channelID medchain -asOrg Storages
FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/patients.medchain.anchors.tx -channelID medchain -asOrg Patients
FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/researchers.medchain.anchors.tx -channelID medchain -asOrg Researchers

#Initialize emr channel
FABRIC_CFG_PATH=$PWD configtxgen -profile emr -outputCreateChannelTx ./artifacts/emr.channel.tx -channelID emr

#Add providers and storages to emr
FABRIC_CFG_PATH=$PWD configtxgen -profile emr -outputAnchorPeersUpdate ./artifacts/providers.emr.anchors.tx -channelID emr -asOrg Providers
FABRIC_CFG_PATH=$PWD configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/storages.emr.anchors.tx -channelID emr -asOrg Storages

#Initialize amr channel
FABRIC_CFG_PATH=$PWD configtxgen -profile amr -outputCreateChannelTx ./artifacts/amr.channel.tx -channelID amr

#Add providers and researchers to channel
FABRIC_CFG_PATH=$PWD configtxgen -profile amr -outputAnchorPeersUpdate ./artifacts/providers.amr.anchors.tx -channelID amr -asOrg Providers
FABRIC_CFG_PATH=$PWD configtxgen -profile amr -outputAnchorPeersUpdate ./artifacts/researchers.amr.anchors.tx -channelID amr -asOrg Researchers

#Initialize pii channel
FABRIC_CFG_PATH=$PWD configtxgen -profile pii -outputCreateChannelTx ./artifacts/pii.channel.tx -channelID pii

#Add providers and patients to channel
FABRIC_CFG_PATH=$PWD configtxgen -profile pii -outputAnchorPeersUpdate ./artifacts/providers.pii.anchors.tx -channelID pii -asOrg Providers
FABRIC_CFG_PATH=$PWD configtxgen -profile pii -outputAnchorPeersUpdate ./artifacts/patients.pii.anchors.tx -channelID pii -asOrg Patients

#Initialize directory channel
FABRIC_CFG_PATH=$PWD configtxgen -profile directory -outputCreateChannelTx ./artifacts/directory.channel.tx -channelID directory

#Add providers and patients to channel
FABRIC_CFG_PATH=$PWD configtxgen -profile directory -outputAnchorPeersUpdate ./artifacts/providers.directory.anchors.tx -channelID directory -asOrg Providers
FABRIC_CFG_PATH=$PWD configtxgen -profile directory -outputAnchorPeersUpdate ./artifacts/patients.directory.anchors.tx -channelID directory -asOrg Patients
