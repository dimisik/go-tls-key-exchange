// Copyright 2019 Thales UK Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/tls"
	"example/common"
	"log"
	"os"
)

// Based on code from https://github.com/denji/golang-tls
func main() {

	conf := &tls.Config{
		MinVersion:          tls.VersionTLS13,
		CurvePreferences:    []tls.CurveID{common.CurveID},
		PrivateKeyExchanges: map[tls.CurveID]tls.PrivateKeyExchange{common.CurveID: &common.KyberKeyExchange{}},
		InsecureSkipVerify:  true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:8443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("PQ-KEM-enabled TLS 1.3 encrypted tunnel established.\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}

func init() {
	// We need TLS 1.3 support
	os.Setenv("GODEBUG", os.Getenv("GODEBUG")+",tls13=1")
}
