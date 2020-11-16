Hyperledger Fabric 2 w/ Kubernetes
============================================

## Sources
- [Hyperledger Getting Started](https://hyperledger-fabric.readthedocs.io/en/release-2.2/getting_started.html).


## Instructions
1. Setup - install linux mint, zshrc, visual studio code, Golang, Nodejs, minikube, kubectl, vuecli, docker, docker-cli, goose
2. Install dependencies via `make install`
3. Set the BIN path so that you can utilize the latest binaries (i.e. export PATH=<path to download location>/bin:$PATH)
    - For me this is in my .zshrc `export PATH=~/Projects/hyperledger_2_fabric_kubernetes/fabric-samples/bin:$PATH`
4. verify install version. for me this says (should be the same Hyperledger version that's in your Makefile)
```bash
☁  hyperledger_2_fabric_kubernetes  configtxgen --version
configtxgen:
 Version: 2.2.1
 Commit SHA: 344fda602
 Go version: go1.14.4
 OS/Arch: linux/amd64
☁  hyperledger_2_fabric_kubernetes  
```



## My own notes
- [Good Readme.md Example](https://raw.githubusercontent.com/vuejs/vue/dev/README.md)

