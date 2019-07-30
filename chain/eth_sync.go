package chain

import (
	"database/sql"
	"errors"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/db"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	. "github.com/elastos/Elastos.ORG.API.Misc/tools"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type eth_transaction struct {
	BlockHash         string
	BlockNumber       string
	From              string
	Gas               string
	To                string
	GasPrice          string
	Hash              string
	Input             string
	Nonce             string
	TransactionIndex  string
	Value             string
	GasUsed           string
	cumulativeGasUsed string
}

var (
	dbaEth     = db.NewInstance()
	currHeight int64
	waitgroup  sync.WaitGroup
)

//Sync sync chain data
func SyncEth() {
	go func() {
		for {
			tx, err := dbaEth.Begin()
			if err = doSyncEth(tx); err != nil {
				log.Infof("Sync ETH Height Error : %v \n", err.Error())
				tx.Rollback()
			} else {
				tx.Commit()
				currHeight += int64(config.Conf.Eth.BatchSize)
			}
			<-time.After(time.Millisecond * 1000)
		}
	}()
}

func doSyncEth(tx *sql.Tx) error {
	if currHeight == 0 {
		r, err := tx.Query("select blockNumber from chain_eth_block_transaction_history order by id desc limit 1")
		if err != nil {
			return err
		}
		if r.Next() {
			r.Scan(&currHeight)
		}
		r.Close()
		if currHeight == 0 {
			currHeight = -1
		}
	}

	if currHeight < config.Conf.Eth.StartHeight && config.Conf.Eth.StartHeight != 0 {
		currHeight = config.Conf.Eth.StartHeight - 1
	}

	var resp map[string]interface{}
	var err error
	resp, err = Post(config.Conf.Eth.Endpoint, `{"jsonrpc":"2.0","method":"eth_blockNumber","params": [],"id":1}`)
	if err != nil {
		return err
	}

	hexHeight, ok := resp["result"]
	if ok {
		height, err := hexutil.DecodeUint64(hexHeight.(string))
		if err != nil {
			return err
		}
		if currHeight == int64(height) {
			return nil
		}
		waitgroup.Add(config.Conf.Eth.BatchSize)
		var unexpected error
		count := 0
		log.Infof("Syncing ETH Height From %d To %d \n", currHeight+1, currHeight+int64(config.Conf.Eth.BatchSize)+1)
		for curr := currHeight + 1; curr <= int64(height); {
			go func() {
				atomic.AddInt64(&curr, 1)
				err = handleHeightEth(int(curr), tx)
				if err != nil {
					unexpected = err
					log.Error("Error handle height " + err.Error() + "\n")
				}
				waitgroup.Done()
			}()
			count++
			if unexpected != nil {
				return unexpected
			}
			if count%config.Conf.Eth.BatchSize == 0 {
				break
			}
		}
	}
	waitgroup.Wait()
	return nil
}

func handleHeightEth(curr int, tx *sql.Tx) error {
	var resp map[string]interface{}
	var err error
	resp, err = Post(config.Conf.Eth.Endpoint, `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params": ["`+hexutil.EncodeUint64(uint64(curr))+`",false],"id":1}`)
	if err != nil {
		return err
	}
	r, ok := (resp["result"].(map[string]interface{}))
	if !ok {
		return errors.New("illegal ETH Height")
	}
	txArr := r["transactions"].([]interface{})
	if len(txArr) == 0 {
		return nil
	}
	timestamp := r["timestamp"]
	for _, txv := range txArr {
		transaction := txv.(string)
		resp, err = Post(config.Conf.Eth.Endpoint, `{"jsonrpc":"2.0","method":"eth_getTransactionByHash","params": ["`+transaction+`"],"id":1}`)
		if err != nil {
			return err
		}
		t := eth_transaction{}
		if resp["result"] == nil {
			log.Errorf("%v ", resp)
			return errors.New("Invalid ETH Node , please change your ethereum node")
		}
		Map2Struct(resp["result"].(map[string]interface{}), &t)
		resp, err = Post(config.Conf.Eth.Endpoint, `{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params": ["`+transaction+`"],"id":1}`)
		if err != nil {
			return err
		}
		if resp["result"] == nil {
			log.Errorf("%v ", resp)
			return errors.New("Invalid ETH Node , please change your ethereum node")
		}
		receipt := resp["result"].(map[string]interface{})
		gasUsed := receipt["gasUsed"]
		status := receipt["status"]
		var isError = "0"
		if status != "0x1" {
			isError = "-1"
		}
		cumulativeGasUsed := receipt["cumulativeGasUsed"]
		contractAddress := receipt["contractAddress"]
		if contractAddress == nil {
			contractAddress = "0x0000000000000000000000000000000000000000"
		}
		stmt, err := tx.Prepare("insert into chain_eth_block_transaction_history (blockNumber,`timeStamp`,hash,nonce,blockHash,transactionIndex,`from`,`to`,`value`,gas,gasPrice,isError,`input`,`contractAddress`,cumulativeGasUsed,gasUsed) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(decode(t.BlockNumber), decode(timestamp.(string)), t.Hash, decode(t.Nonce), t.BlockHash, decode(t.TransactionIndex), t.From, t.To, decode(t.Value), decode(t.Gas), decode(t.GasPrice), isError, t.Input, contractAddress, decode(cumulativeGasUsed.(string)), decode(gasUsed.(string)))
		if err != nil {
			return err
		}
	}
	return nil
}

func decode(str string) string {
	desc, _ := strconv.ParseUint(str[2:], 16, 64)
	return strconv.Itoa(int(desc))
}
