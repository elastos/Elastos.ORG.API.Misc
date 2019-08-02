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
		r, err := tx.Query("select blockNumber from chain_eth_block_transaction_history order by blockNumber desc limit 1")
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
	var unexpected error = nil
	if ok {
		height, err := hexutil.DecodeUint64(hexHeight.(string))
		if err != nil {
			return err
		}
		if currHeight == int64(height) {
			return nil
		}
		waitgroup.Add(config.Conf.Eth.BatchSize)
		count := 0
		log.Infof("Syncing ETH Height From %d To %d \n", currHeight+1, currHeight+int64(config.Conf.Eth.BatchSize)+1)
		for curr := currHeight; curr <= int64(height); {
			go func() {
				atomic.AddInt64(&curr, 1)
				err = handleHeightEth(int(curr), tx)
				if err != nil {
					unexpected = err
				}
				waitgroup.Done()
			}()
			count++
			if count%config.Conf.Eth.BatchSize == 0 {
				break
			}
		}
	}
	waitgroup.Wait()
	return unexpected
}

func handleHeightEth(curr int, tx *sql.Tx) error {
	var resp map[string]interface{}
	var err error
	for i := 0; i < config.Conf.Eth.Retry; i++ {
		resp, err = Post(config.Conf.Eth.Endpoint, `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params": ["`+hexutil.EncodeUint64(uint64(curr))+`",true],"id":1}`)
		if err == nil {
			break
		}
	}
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
		transaction := txv.(map[string]interface{})
		t := eth_transaction{}
		Map2Struct(transaction, &t)
		for i := 0; i < config.Conf.Eth.Retry; i++ {
			resp, err = Post(config.Conf.Eth.Endpoint, `{"jsonrpc":"2.0","method":"eth_getTransactionReceipt","params": ["`+transaction["hash"].(string)+`"],"id":1}`)
			if err == nil {
				break
			}
		}
		if err != nil {
			return err
		}
		if resp["result"] == nil {
			log.Errorf("%v ", resp)
			return errors.New("Invalid ETH Node , please change your ethereum node")
		}
		receipt := resp["result"].(map[string]interface{})
		gasUsed := receipt["gasUsed"]
		if gasUsed == nil {
			gasUsed = ""
		}
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
	if len(str) == 0 {
		return ""
	}
	desc, _ := strconv.ParseUint(str[2:], 16, 64)
	return strconv.Itoa(int(desc))
}
