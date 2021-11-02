package main

import (
	"encoding/hex"
	"fmt"
	"os"

	gkeys "github.com/number571/go-cryptopro/gost_r_34_10_2012"
	ghash "github.com/number571/go-cryptopro/gost_r_34_11_2012"
)

func sign(in *InPut) (*OutPut, error) {
	var (
		hash      = make([]byte, ghash.Size256)
		signature = make([]byte, gkeys.SignatureSize256)
	)
	data := []byte(in.Data)
	cfg := gkeys.NewConfig(gkeys.K256, "username", "password")
	err := gkeys.GenPrivKey(cfg)
	if err != nil {
		fmt.Println("Warning: key already exist?")
	}
	priv, err := gkeys.NewPrivKey(cfg)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	pub := priv.PubKey()
	hash = ghash.Sum(ghash.H256, data)
	signature = signHash(priv, hash)
	fmt.Println(hex.EncodeToString(signature))

	return &OutPut{
		Data:      in.Data,
		Signature: hex.EncodeToString(signature),
		Pubkey:    hex.EncodeToString(pub.Bytes()),
	}, nil
}

func signHash(priv gkeys.PrivKey, hash []byte) []byte {
	sign, err := priv.Sign(hash)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(10)
	}
	return sign
}
