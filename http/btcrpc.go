package http

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"strconv"
)

var (
	client *rpcclient.Client
	btcNet *chaincfg.Params
)

type rpchelper struct {
	param map[string]string
}

func init() {
	if config.Conf.Btc.Enable {
		go func() {
			var err error
			client, err = rpcclient.New(&rpcclient.ConnConfig{
				HTTPPostMode: true,
				DisableTLS:   true,
				Host:         config.Conf.Btc.Host,
				User:         config.Conf.Btc.Rpcuser,
				Pass:         config.Conf.Btc.Rpcpasswd,
			}, nil)
			if err != nil {
				fmt.Println("Error Connect to Bitcoin node :", err.Error())
			}
			net := config.Conf.Btc.Net
			if net == "mainet" {
				btcNet = &chaincfg.MainNetParams
			} else if net == "regtest" {
				btcNet = &chaincfg.RegressionNetParams
			} else {
				btcNet = &chaincfg.TestNet3Params
			}
		}()
	}
}

func (h *rpchelper) getBalance() (float64, error, int) {
	addr := h.param["addr"]
	address, err := btcutil.DecodeAddress(addr, btcNet)
	utxos, err := client.ListUnspentMinMaxAddresses(config.Conf.Btc.MinConfirm, 99999999, []btcutil.Address{address})
	if err != nil {
		return 0, err, 400
	}
	var totalBalance float64
	for _, v := range utxos {
		totalBalance += v.Amount
	}

	txIds, err := client.GetRawMempool()
	if err != nil {
		return 0, err, 500
	}
	for _, txId := range txIds {
		txInfo, err := client.GetRawTransactionVerbose(txId)
		if err != nil {
			fmt.Printf("Error fetching transaction %s \n", txId.String())
			continue
		}
		for _, in := range txInfo.Vin {
			hash := &chainhash.Hash{}
			b, _ := hex.DecodeString(in.Txid)
			hash.SetBytes(tools.ReverseBytes(b))
			r, err := client.GetRawTransactionVerbose(hash)
			if err != nil {
				fmt.Printf("Error fetching transaction %s \n", in.Txid)
				continue
			}
			opAddrs := r.Vout[in.Vout].ScriptPubKey.Addresses
			if tools.Contains(opAddrs, addr) {
				totalBalance += r.Vout[in.Vout].Value
			}
		}
	}
	return totalBalance, nil, 200
}

func (h *rpchelper) getTransaction() (string, error, int) {
	txid := h.param["txid"]
	btxid, err := hex.DecodeString(txid)
	if err != nil || len(btxid) != 32 {
		return "", errors.New("Invalid txid"), 400
	}

	hash := &chainhash.Hash{}
	hash.SetBytes(tools.ReverseBytes(btxid))
	tx, err := client.GetRawTransactionVerbose(hash)

	if err != nil {
		return "", err, 500
	}
	buf, _ := json.Marshal(tx)
	return string(buf), nil, 200
}

func (h *rpchelper) getBestheight() (int64, error, int) {
	blockHeight, err := client.GetBlockCount()
	if err != nil {
		return -1, err, 500
	}
	return blockHeight, nil, 200
}

func (h *rpchelper) getBlockDetail() (string, error, int) {
	height := h.param["height"]
	i, err := strconv.Atoi(height)

	if err != nil || i < 0 {
		return "", err, 400
	}

	bestHeight, err, _ := h.getBestheight()
	if err != nil {
		return "", err, 500
	}

	if i > int(bestHeight) {
		return "", errors.New("No such height"), 400
	}

	blockhash, err := client.GetBlockHash(int64(i))
	if err != nil {
		return "", err, 500
	}

	block, err := client.GetBlockVerbose(blockhash)
	if err != nil {
		return "", err, 500
	}

	result, _ := json.Marshal(block)

	return string(result), nil, 200
}
