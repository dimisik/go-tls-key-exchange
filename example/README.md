# Example TLS 1.3 Post-Quantum Key Exchange

This example project combines this fork of Go with [goliboqs](https://github.com/thales-e-security/goliboqs) to
demonstrate a TLS handshake with a post-quantum candidate algorithm. goliboqs is a Go wrapper around the
https://github.com/open-quantum-safe/liboqs library.

## Build

### Install liboqs

Follow [the instructions](https://github.com/open-quantum-safe/liboqs#building-and-running-liboqs-master-branch) to 
build and install liboqs. This example code assumes you have installed liboqs into `/usr/local/liboqs`. 
This is achieved by calling `configure` like so:

```go
./configure --prefix=/usr/local/liboqs [additional args...]
```

Don't forget to run `make install` (most like with 'sudo') once you've built and tested liboqs.

### Install Go fork

See the parent 
[README.md](https://github.com/thales-e-security/go-tls-key-exchange/blob/go1.12.5_private_key_exchanges/README.md)
file to understand how to install the Go fork.

The remaining instructions assume you've installed and compiled the fork in `~/git/go-tls-key-exchange`.

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
hello

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
world


```

# Acknowledgements

This example code is based on https://github.com/denji/golang-tls.
