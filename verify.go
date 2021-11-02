package main

import (
	"encoding/hex"
	"fmt"
	"os"

	gkeys "github.com/number571/go-cryptopro/gost_r_34_10_2012"
	ghash "github.com/number571/go-cryptopro/gost_r_34_11_2012"
)

func verify(in string, out *OutPut) (bool, error) {
	var (
		data = make([]byte, 2048)
		hash = make([]byte, ghash.Size256)
		sign = make([]byte, gkeys.SignatureSize256)
	)

	pub, err := gkeys.LoadPubKey(decodeHex(out.Pubkey))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}
	sign = decodeHex(out.Signature)
	data = []byte(in)
	hash = ghash.Sum(ghash.H256, data)
	return verifyHash(pub, hash, sign), nil
}

func decodeHex(hexdata string) []byte {
	data, err := hex.DecodeString(hexdata)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(10)
	}
	return data
}

func verifyHash(pub gkeys.PubKey, hash, sign []byte) bool {
	return pub.VerifySignature(hash, sign)
}
