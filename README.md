# [Rolls-Royce-Challenge](https://www.rolls-royce-blockchain-innovation-challenge.com/challenges)

Create a Consortium Blockchain for Supply Chain and Logistics Management.

## Introduction - An Introduction of Trellis Labs

Trellis Labs is a consortia to create broad collaboration on blockchain projects with Industry actors. Our team of experts works with you to develop blockchain solutions. From identifying the best use-case, assessing suitability, developing a proof of concepts, and building market-ready solutions.

## Challenge of choice - Explain the challenge of choice
> Traceability and Transparency

The potential blockchain solution is proposed, To introduce Traceability and Transparency, to improve Integrity and making it Seamless. For Prototype, we have focused our work on the first four steps of Maintenance Scenario, in the future, we will be focusing on the next 12 steps. This Planned protocol will be able to reduce the complete process time period by 33% with enabled blockchain security. The solution doesn't use any 3rd party involvement.
This Solution is developed by keeping in mind the generic implementation and replaces the physical documents with a consortium blockchain ledger.

## Solution - Articulate the idea and your strategy in solving the problem.


## Tech - What makes your blockchain technology special? Elaborate on tools and approach used.

* Hyperledger Fabric
  * Consortium Blockchain
  * Flexible Architecture

* Golang(ChainCode)
* Nodejs
* Flask

## Beyond challenge - How can your solution be implemented into an existing business? Elaborate on compatibility.

## Setting-up Developement Environment
### Prerequisites
1. cUrl(7.x)
```
sudo apt-get update
sudo apt-get install curl
curl --version
```
2. Docker(>17.x) and Docker-compose(>1.14.0)
```
sudo snap install docker
docker --version
docker-compose --version
```
3. Go(1.11.x)
- Download the [archive](https://golang.org/dl).
```
cd Archive_Download_location
sudo tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
source $HOME/.profile
go version
```
4. Nodejs(>10.x)
```
curl -sL https://deb.nodesource.com/setup_12.x | sudo -E bash -
sudo apt-get install nodejs
node -v
npm -v
```
5. Python
```
sudo apt-get install python
python --version
```
### Building the Network
1. Install Samples, Binaries and Docker Images
```
cd RollsRoyce/src/scripts
./bootstrap.sh
```

2. Install npm
```
cd ../client/javascript
npm install
```

3. Bringing Up the Network
```
cd ..
./startFabric.sh
```

4. Running the Frontend
```
cd py/
pip3 install -r requirements.txt
python3 app.py
```
