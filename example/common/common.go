// Copyright 2019 Thales UK Limited. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package common

const CurveID = 0xFE00

type SimpleKeyExchange struct{}

func (SimpleKeyExchange) ClientShare() ([]byte, error) {
	return []byte("not used"), nil
}

func (SimpleKeyExchange) SecretFromClientShare(clientShare []byte) (secret, serverShare []byte, err error) {
	return []byte("secret"), []byte("not used"), nil
}

func (SimpleKeyExchange) SecretFromServerShare(serverShare []byte) ([]byte, error) {
	return []byte("secret"), nil
}
