# Ethereum Examples
This repo holds various Ethereum examples and PoCs.

## Vyper
Steps to generate go functions for interaction and deployment of Smart Contracts:

```
vyper -f abi example.vy > example.abi
vyper -f json example.vy > example.abi
abigen --abi example.abi --bin example.bin --type Example --pkg main --out example.go
```

## Deployment
### Geth
For deployment with `geth` refer to [this wiki article](https://github.com/ethereum/go-ethereum/wiki/Native-DApps:-Go-bindings-to-Ethereum-contracts).

### Ganash
For deployment on Ganash simulator consult [this issue](https://github.com/trufflesuite/ganache-cli/issues/555).
