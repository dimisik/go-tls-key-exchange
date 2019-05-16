// Copyright 2019 Thales UK Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

import (
	"errors"
	"fmt"
	"log"

	"github.com/thales-e-security/goliboqs"
)

const CurveID = 0xFE00
const libPath = "/usr/local/liboqs/lib/liboqs.so"

type KyberKeyExchange struct {
	privateKey []byte
}

func loadKEM() (*goliboqs.Lib, goliboqs.Kem, error) {
	lib, err := goliboqs.LoadLib(libPath)
	if err != nil {
		return nil, nil, err
	}

	kem, err := lib.GetKem(goliboqs.KemKyber512)
	if err != nil {
		_ = lib.Close()
		return nil, nil, err
	}

	return lib, kem, nil
}

func (k *KyberKeyExchange) ClientShare() (share []byte, err error) {
	lib, kem, err := loadKEM()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = kem.Close()
		_ = lib.Close()
	}()

	pub, priv, err := kem.KeyPair()
	if err != nil {
		return nil, err
	}

	k.privateKey = priv

	log.Println("ClientShare")
	log.Printf("\t[out] share=%s\n", truncateShare(pub))

	return pub, nil
}

func (*KyberKeyExchange) SecretFromClientShare(clientShare []byte) (secret, serverShare []byte, err error) {
	lib, kem, err := loadKEM()
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		_ = kem.Close()
		_ = lib.Close()
	}()

	secret, serverShare, err = kem.Encaps(clientShare)

	log.Println("SecretFromClientShare")
	log.Printf("\t[in]  share=%s", truncateShare(clientShare))
	log.Printf("\t[out] secret=%x", secret)
	log.Printf("\t[out] share=%s", truncateShare(serverShare))

	return
}

func (k *KyberKeyExchange) SecretFromServerShare(serverShare []byte) ([]byte, error) {
	if k.privateKey == nil {
		return nil, errors.New("call to SecretFromServerShare without call to ClientShare")
	}

	lib, kem, err := loadKEM()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = kem.Close()
		_ = lib.Close()
	}()

	res, err := kem.Decaps(serverShare, k.privateKey)
	log.Println("SecretFromServerShare")
	log.Printf("\t[in]  share=%s", truncateShare(serverShare))
	log.Printf("\t[out] secret=%x", res)

	return res, err
}

func truncateShare(in []byte) string {
	const maxLen = 50

	if len(in) <= maxLen {
		return fmt.Sprintf("%x", in)
	}

	return fmt.Sprintf("%x...", in[:maxLen-3])
}
