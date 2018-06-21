# MedChain Network

Work in progress (WIP)

## Setup ##
### Docker ###
 1.  

### GO ###
 1.  

### Binaries ###
 1. `cd $GOPATH`
 1. `mkdir -p src/github.com/hyperledger/`
 1. `cd src/github.com/hyperledger/`
 1. `curl -sSL https://goo.gl/6wtTN5 | bash -s 1.1.0`
 1. edit your .bashrc file to include the path to your new src/github.com/hyperledger/fabric-samples/bin/ dir in $PATH

### clone medchain-network ###
 1. cd into your own projects folder
 1. `git clone git@github.com:MedChain/medchain-network.git`
 1. `cd medchain-network`

### setup the fixture/network ###
 1. `cd fixtures/`
 1. `./setup.sh`
 1. `cd ..`

### test ###
 1. `make`
 1. go to http://localhost:3000/ in a browser
