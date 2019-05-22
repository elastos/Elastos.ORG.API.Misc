package tools

import (
	"encoding/hex"
	"github.com/elastos/Elastos.ELA.Utility/crypto"
)

func GetAddress(publicKeyHex string) (string, error) {
	publicKey, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", err
	}
	pub, err := crypto.DecodePoint(publicKey)
	if err != nil {
		return "", err
	}
	code, err := crypto.CreateStandardRedeemScript(pub)
	if err != nil {
		return "", err
	}
	hash, err := crypto.ToProgramHash(code)
	if err != nil {
		return "", err
	}
	addr, err := hash.ToAddress()
	if err != nil {
		return "", err
	}
	return addr, nil
}
