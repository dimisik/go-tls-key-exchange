# go1.12.5_private_key_exchanges
[![Build Status](https://travis-ci.com/thales-e-security/go.svg?branch=go1.12.5_private_key_exchanges)](https://travis-ci.com/thales-e-security/go)

This fork of Go 1.12.5 supports custom key encapsulation mechanisms (KEMs) in `crypto/tls` for TLS 1.3 (only). Users of 
this fork may include researchers experimenting with quantum-resistant algorithms, such as those being 
[assessed by NIST](https://csrc.nist.gov/Projects/Post-Quantum-Cryptography).

To see the changes introduced in this fork, please review https://github.com/golang/go/compare/go1.12.5...thales-e-security:go1.12.5_private_key_exchanges.

## Building

Please follow the standard instructions for building Go from source: https://golang.org/doc/install/source.

## Adding new KEMs

This fork introduces a new interface: `tls.PrivateKeyExchange`:

```go
// A PrivateKeyExchange implements a TLS 1.3 key exchange mechanism.
type PrivateKeyExchange interface {

	// ClientShare initiates the key exchange and returns the client key
	// share.
	ClientShare() ([]byte, error)

	// SecretFromClientShare is called by the server to process the share from
	// the client. It generates (or deduces) the TLS secret and returns this, along with
	// its own share, which is sent to the client.
	SecretFromClientShare(clientShare []byte) (secret, serverShare []byte, err error)

	// SecretFromServerShare uses the server key share to deduce the TLS secret,
	// which is returned.
	SecretFromServerShare(serverShare []byte) ([]byte, error)
}
```

This interface should be familiar to anyone working with KEMs from the NIST competition.

The `tls.Config` struct has been extended with the following field:

```go
type Config struct {
	//...
	
    // PrivateKeyExchanges are TLS 1.3 key exchange implementations
    // for private named groups. The CurveIDs must be from the ecdhe_private_use range
    // (see RFC 8446 section 4.2.7). To enable these private groups, include their
    // CurveID in the CurvePreferences field.
    PrivateKeyExchanges map[CurveID]PrivateKeyExchange
}
```  

The `CurveID` chosen must be from the [ecdhe_private_use range](https://tools.ietf.org/html/rfc8446#section-4.2.7),
i.e. 0xFE00..0xFEFF. The curve IDs must also be added to `Config.CurvePreferences` otherwise they will be ignored.

See the `example` directory for a sample TLS server and client that use a dummy KEM to negotiate their connection.

### Enabling TLS 1.3 Support

In Go v1.12.x, TLS 1.3 support is optional and disabled by default. Quoting from their documentation:
>
> TLS 1.3 is available only on an opt-in basis in Go 1.12. To enable it, set the GODEBUG environment variable 
> (comma-separated key=value options) such that it includes "tls13=1". To enable it from within the process, 
> set the environment variable before any use of TLS:
>
> ```go
> func init() {
>     os.Setenv("GODEBUG", os.Getenv("GODEBUG")+",tls13=1")
> } 
> ```

## Support in Go proper

A proposal has been opened to add similar functionality into Go. Please see https://github.com/golang/go/issues/31520.
