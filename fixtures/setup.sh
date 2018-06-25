#!/bin/bash
#
# Copyright MedChain, Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

export FABRIC_CFG_PATH=$PWD

#Clean up
rm -rf crypto-config/
rm -rf artifacts/

#Create certificates
cryptogen generate --config=./crypto-config.yaml

mkdir -p artifacts

#Create genesis block
configtxgen -profile OrdererGenesis -outputBlock ./artifacts/orderer.genesis.block

#medchain: Initialize channel, Add orgs
configtxgen -profile medchain -outputCreateChannelTx ./artifacts/medchain.channel.tx -channelID medchain
configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/providers.medchain.anchors.tx -channelID medchain -asOrg Providers
configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/storages.medchain.anchors.tx -channelID medchain -asOrg Storages
configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/patients.medchain.anchors.tx -channelID medchain -asOrg Patients
configtxgen -profile medchain -outputAnchorPeersUpdate ./artifacts/researchers.medchain.anchors.tx -channelID medchain -asOrg Researchers

#EMR: Initialize channel, Add orgs
configtxgen -profile emr -outputCreateChannelTx ./artifacts/emr.channel.tx -channelID emr
configtxgen -profile emr -outputAnchorPeersUpdate ./artifacts/providers.emr.anchors.tx -channelID emr -asOrg Providers
configtxgen -profile emr -outputAnchorPeersUpdate ./artifacts/storages.emr.anchors.tx -channelID emr -asOrg Storages

#AMR: Initialize channel, Add orgs
configtxgen -profile amr -outputCreateChannelTx ./artifacts/amr.channel.tx -channelID amr
configtxgen -profile amr -outputAnchorPeersUpdate ./artifacts/providers.amr.anchors.tx -channelID amr -asOrg Providers
configtxgen -profile amr -outputAnchorPeersUpdate ./artifacts/researchers.amr.anchors.tx -channelID amr -asOrg Researchers

#PII: Initialize channel, Add orgs
configtxgen -profile pii -outputCreateChannelTx ./artifacts/pii.channel.tx -channelID pii
configtxgen -profile pii -outputAnchorPeersUpdate ./artifacts/providers.pii.anchors.tx -channelID pii -asOrg Providers
configtxgen -profile pii -outputAnchorPeersUpdate ./artifacts/patients.pii.anchors.tx -channelID pii -asOrg Patients

#DIRECTORY: Initialize channel, Add orgs
configtxgen -profile directory -outputCreateChannelTx ./artifacts/directory.channel.tx -channelID directory
configtxgen -profile directory -outputAnchorPeersUpdate ./artifacts/providers.directory.anchors.tx -channelID directory -asOrg Providers
configtxgen -profile directory -outputAnchorPeersUpdate ./artifacts/patients.directory.anchors.tx -channelID directory -asOrg Patients
