# Computing Infrastructure Migration

This repository provides computing infrastructure migration features.
This is a sub-system on [Cloud-Barista platform](https://github.com/cloud-barista/docs)
and utilizes [CB-Tumblebug](https://github.com/cloud-barista/cb-tumblebug)
to depoly a multi-cloud infra as a target computing infrastructure.


## Overview

Computing Infrastructure Migration framework (codename: cm-beetle) is going to support:
- migration execution and control from source to target computing infrastructure, and
- recommendation of optimal configuration of target cloud infrastructure.


## Execution and development environment

- Operating system (OS): 
    - Ubuntu 20.04
- Languages: 
    - Go: 1.19
    - Python: 3.8.10
- Container runtime:
    - Docker: 20.10.12


## How to run CM-Beetle

### Source code based installation and execution

#### Configure build environment

1. Install dependencies

```bash
# Ensure that your system is up to date
sudo apt update -y

# Ensure that you have installed the dependencies, 
# such as `ca-certificates`, `curl`, and `gnupg` packages.
sudo apt install make gcc git
```
2. Install Go

To install Go v1.19+, see [Go all releases](https://golang.org/dl/) and [Download and install](https://go.dev/doc/install)

Example - Go 1.19 installtion 

```bash
# Get Go archive
wget https://go.dev/dl/go1.19.linux-amd64.tar.gz

# Remove any previous Go installation and
# Extract the archive into /usr/local/
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz

# Append /usr/local/go/bin to .bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> ~/.bashrc
echo 'export GOPATH=$HOME/go' >> ~/.bashrc

# Apply the .bashrc changes
source ~/.bashrc

# Verify the installation
echo $GOPATH
go version
```

#### Download source code

1. Clone CM-Beetle repository

```bash
git clone https://github.com/cloud-barista/cm-beetle.git ${HOME}/cm-beetle
```

#### Build CM-Beetle

```bash
cd ${HOME}/cm-beetle/pkg
make
```

(Optional) Update Swagger API document
```bash
cd ${HOME}/cm-beetle/pkg
make swag
```

#### Run CM-Beetle binary

```bash
cd ${HOME}/cm-beetle/pkg
make run
```

#### Health-check CM-Beetle

```bash
curl http://localhost:8056/beetle/health

# Output if it's running successfully
# {"message":"CM-Beetle API server is running"}
```


### Container based execution

Check a tag of CM-Beetle container image in [cloudbaristaorg/cm-beetle](https://hub.docker.com/r/cloudbaristaorg/cm-beetle/tags)

#### Run CM-Beetle container

```bash
docker run -p 8056:8056 \
--name cm-beetle \
cloudbaristaorg/cm-beetle:latest
```

#### Health-check CM-Beetle
```bash
curl http://localhost:8056/beetle/health

# Output if it's running successfully
# {"message":"CM-Beetle API server is running"}
```