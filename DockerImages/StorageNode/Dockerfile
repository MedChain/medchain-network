FROM hyperledger/fabric-orderer:x86_64-1.0.5
MAINTAINER russ.decker@medchain.us



COPY ./go-ipfs/* ipfs/
COPY ./go-ipfs/ipfs /usr/local/bin
COPY ./go-ipfs/ipfs /usr/bin

RUN mkdir -p ipfs/data && cd ipfs/data && ipfs init 

# Ports for Swarm TCP, Swarm uTP, API, Gateway, Swarm Websockets 
EXPOSE 4001 
EXPOSE 4002/udp 
EXPOSE 5001 
EXPOSE 8080 
EXPOSE 8081


