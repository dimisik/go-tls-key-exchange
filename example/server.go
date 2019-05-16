// Copyright 2019 Thales UK Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"crypto/tls"
	"example/common"
	"log"
	"net"
	"os"
)

// Based on code from https://github.com/denji/golang-tls
func main() {
	config := &tls.Config{
		MinVersion: tls.VersionTLS13,
		Certificates: []tls.Certificate{
			{
				Certificate: [][]byte{testRSACertificate, testRSACertificateIssuer},
				PrivateKey:  testRSAPrivateKey,
			},
		},
		CurvePreferences:    []tls.CurveID{common.CurveID},
		PrivateKeyExchanges: map[tls.CurveID]tls.PrivateKeyExchange{common.CurveID: &common.KyberKeyExchange{}},
	}

	ln, err := tls.Listen("tcp", ":8443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

func init() {
	// We need TLS 1.3 support
	os.Setenv("GODEBUG", os.Getenv("GODEBUG")+",tls13=1")
}
