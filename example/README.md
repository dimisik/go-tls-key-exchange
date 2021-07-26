# Post-Quantum Key Exchange Example in Golang TLS 1.3 

This script automatically installs a modified GO fork that enables the use of Post-Quantum key exchanges using the [OQS liboqs](https://github.com/open-quantum-safe/liboqs) library.
This is enabled by a custom Golang wrapper around liboqs, namely [dimisik/goliboqs](https://github.com/dimisik/goliboqs) that supports the latest liboqs version.
The original goliboqs wrapper can be found in https://github.com/thales-e-security/goliboqs (current liboqs lib support is discontinued). 
IN what follows, we describe the PQ-KEM-enabled Go fork installation process step-by-step.

## Prerequisites

Install Golang, either using the Go binary from Google or by running:
```
cd /tmp
wget https://golang.org/dl/go1.16.6.linux-amd64.tar.gz
sudo  tar -C /usr/local -xzf go1.16.6.linux-amd64.tar.gz
```
Add /usr/local/go/bin to the PATH environment variable by adding the following line to your $HOME/.profile or /etc/profile (for a system-wide installation):
```
export PATH=$PATH:/usr/local/go/bin
```
Save and close the file when using vim. Run the source command into the current bash/shell to load environment variables on Ubuntu. E.g.:
```
source ~/.profile
```
Verify by running:
```
go version
```

## Building

### 1. Build/Install liboqs

- Install dependencies:
    ```
	sudo apt install astyle cmake gcc ninja-build libssl-dev python3-pytest python3-pytest-xdist unzip xsltproc doxygen graphviz
    ```
- Get the source:
    ```
	git clone -b main https://github.com/open-quantum-safe/liboqs.git
	cd liboqs
    ```
- Build:
    ```
    mkdir build && cd build 
    ```
    a. There is the option to choose which PQ schemes to build by using (here only kyber/dilithium are built): 
	```
        cmake  -DBUILD_SHARED_LIBS=ON  -DOQS_ENABLE_KEM_BIKE=OFF -DOQS_ENABLE_KEM_CLASSIC_MCELIECE=OFF -DOQS_ENABLE_KEM_FRODOKEM=OFF -DOQS_ENABLE_KEM_HQC=OFF -DOQS_ENABLE_KEM_NTRU=OFF -DOQS_ENABLE_KEM_NTRUPRIME=OFF -DOQS_ENABLE_KEM_SABER=OFF  -DOQS_ENABLE_KEM_SIDH=OFF -DOQS_ENABLE_KEM_SIKE=OFF -DOQS_ENABLE_SIG_PICNIC=OFF -DOQS_ENABLE_SIG_RAINBOW=OFF -DOQS_ENABLE_SIG_SPHINCS=OFF   -GNinja  .. 
    ```
    b. Alternatively to build everything:  
    ```  
	cmake -DBUILD_SHARED_LIBS=ON -GNinja ..
	```
>  Various `cmake` build options to customize the resultant artifacts are available and are [documented in the project Wiki](https://github.com/open-quantum-safe/liboqs/wiki/Customizing-liboqs).   

and finaly run:
```
    ninja
    sudo ninja install
```

The above will result to the creation of liboqs.so file in the /usr/local/lib directorycd .

### 2. Install modyfied Go fork from source

- Get the source:
    ```
    git clone https://github.com/dimisik/go-tls-key-exchange
    cd go-tls-key-exchange
    go env -w GO111MODULE=auto
    ./all.bash
    cd ../..
    ```
  The PQ-kEM-enabled GO fork installed commands are in `[path-to]/go-tls-key-exchange/bin`

- Get the goliboqs wrapper into go-tls-key-exchange directory:
    ```
    cd go-tls-key-exchange
    git clone https://github.com/dimisik/goliboqs
    ```


### Build example code

Building the example is now as simple as updating your `GOROOT` and building the server and client binaries:

```
cd ~/git/go-tls-key-exchange/example
export GOROOT=~/git/go-tls-key-exchange/
go build -o run-server
go build -o run-client ./client/...
```

## Run

In one terminal, execute `./run-server`. In another terminal, execute `./run-client`. The output should look
like this:

**Server**

```
$ ./run-server 
2019/05/16 10:26:43 SecretFromClientShare
2019/05/16 10:26:43 	[in]  share=01eff9003691a93814f89a86ce6680df13700e692f712b20e9143d262bc329ef0ec7243e17cc65df18baf5a0f196ab...
2019/05/16 10:26:43 	[out] secret=49f1dbc83278821d27596c1c31688da6cf89b33040c337f4caea0ab7f65ec3aa
2019/05/16 10:26:43 	[out] share=076788e16abec807641f57a07d78238d2afdf483840479cff0c47ddc4ae2acca86db38062bab977bfe5fc9b9a7fec3...
PQ-KEM-enabled TLS 1.3 encrypted tunnel established.

2019/05/16 10:26:43 EOF

```

(You will need to Ctrl-c to close the server)

**Client**

```
$ ./run-client 
2019/05/16 10:26:43 ClientShare
2019/05/16 10:26:43 	[out] share=01eff9003691a93814f89a86ce6680df13700e692f712b20e9143d262bc329ef0ec7243e17cc65df18baf5a0f196ab...
2019/05/16 10:26:43 SecretFromServerShare
2019/05/16 10:26:43 	[in]  share=076788e16abec807641f57a07d78238d2afdf483840479cff0c47ddc4ae2acca86db38062bab977bfe5fc9b9a7fec3...
2019/05/16 10:26:43 	[out] secret=49f1dbc83278821d27596c1c31688da6cf89b33040c337f4caea0ab7f65ec3aa
PQ-KEM-enabled TLS 1.3 encrypted tunnel established.


```

# Acknowledgements

This example code is based on https://github.com/thales-e-security/go-tls-key-exchange.

This example code is based on https://github.com/denji/golang-tls.
