# Elrond Stress
Elrond stress testing tools

## Installation/compilation: ##
```
mkdir -p $GOPATH/src/github.com/herumi
cd $GOPATH/src/github.com/herumi
git clone https://github.com/herumi/mcl
git clone https://github.com/herumi/bls

mkdir -p $GOPATH/src/github.com/SebastianJ
cd $GOPATH/src/github.com/SebastianJ
git clone https://github.com/SebastianJ/elrond-sdk
git clone https://github.com/SebastianJ/elrond-stress

cd $GOPATH/src/github.com/SebastianJ/elrond-stress
make clean && make all
```

#### Execution: ####
Put a funded wallet/.pem key (ideally with >2.5m) in `keys/`
If you want to propagate txs with payloads, place your payload in `data/data.txt`
Run the tool:
```
./dist/stress txs --amount 1 --concurrency 1000
```
