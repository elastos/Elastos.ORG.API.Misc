package chain

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	. "github.com/elastos/Elastos.ORG.API.Misc/tools"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"io"
	"math/rand"
	"os"
	"os/user"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	levelDbPath = "/.misc/eth"
	le          *level
	Tokens      = make([]interface{}, 0)
)

type level struct {
	l          *leveldb.DB
	b          *leveldb.Batch
	currHeight int64
	waitGroup  sync.WaitGroup
	m          sync.Mutex
	path       string
}

type key_prefix byte

var (
	curr_height_prefix       key_prefix = 0x01
	eth_history_prefix       key_prefix = 0x02
	eth_token_history_prefix key_prefix = 0x03
)

type Eth_transaction struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	From              string `json:"from"`
	Gas               string `json:"gas"`
	To                string `json:"to"`
	GasPrice          string `json:"gasPrice"`
	Hash              string `json:"hash"`
	Input             string `json:"input"`
	Nonce             string `json:"nonce"`
	TransactionIndex  string `json:"transactionIndex"`
	Value             string `json:"value"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	IsError           string `json:"isError"`
	ContractAddress   string `json:"contractAddress"`
	Timestamp         string `json:"timeStamp"`
	Confirmations     string `json:"confirmations"`
	TransferType      string `json:"transferType,omitempty"`
}

type Eth_token_transaction struct {
	Address          string        `json:"address"`
	BlockNumber      string        `json:"blockNumber"`
	Data             string        `json:"data"`
	LogIndex         string        `json:"logIndex"`
	Topics           []interface{} `json:"topics"`
	TransactionHash  string        `json:"transactionHash"`
	TransactionIndex string        `json:"transactionIndex"`
	GasUsed          string        `json:"gasUsed"`
	GasPrice         string        `json:"gasPrice"`
	TimeStamp        string        `json:"timeStamp"`
}

const (
	sidechain_crossChain_deposit  = "crossChainEthDeposit"
	sidechain_crossChain_withdraw = "crossChainEthWithdraw"
	sidechain_eth_transfer        = "ethTransfer"
)

type TransactionHistorySorter []Eth_transaction

func (a TransactionHistorySorter) Len() int           { return len(a) }
func (a TransactionHistorySorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TransactionHistorySorter) Less(i, j int) bool { return a[i].Timestamp < a[j].Timestamp }

func (tx *Eth_transaction) Deserialize(data []byte) error {
	var r bytes.Buffer
	_, err := r.Write(data)
	if err != nil {
		return err
	}

	// blockHash
	len, err := readByte(&r)
	if err != nil {
		return err
	}
	blockHash, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	tx.BlockHash = string(blockHash)

	// blockNumber
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	blockNumber, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.BlockNumber = decodeHexToDecimal(string(blockNumber))

	// From

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	From, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	tx.From = string(From)

	// Gas

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Gas, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.Gas = decodeHexToDecimal(string(Gas))

	// To

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	To, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	tx.To = string(To)

	// GasPrice

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	GasPrice, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.GasPrice = decodeHexToDecimal(string(GasPrice))

	// Hash

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Hash, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	tx.Hash = string(Hash)

	//Nonce

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Nonce, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.Nonce = decodeHexToDecimal(string(Nonce))

	// TransactionIndex

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	TransactionIndex, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.TransactionIndex = decodeHexToDecimal(string(TransactionIndex))

	// Value

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Value, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.Value = decodeHexToDecimal(string(Value))

	// GasUsed

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	GasUsed, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.GasUsed = decodeHexToDecimal(string(GasUsed))

	// CumulativeGasUsed

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	CumulativeGasUsed, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.CumulativeGasUsed = decodeHexToDecimal(string(CumulativeGasUsed))

	// IsError
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	IsError, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	tx.IsError = string(IsError)

	// ContractAddress

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	ContractAddress, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	tx.ContractAddress = string(ContractAddress)

	//Timestamp

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Timestamp, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.Timestamp = decodeHexToDecimal(string(Timestamp))

	//TxType

	if config.Conf.Eth.SideChain {
		len, err = readByte(&r)
		if err != nil {
			return err
		}
		TxType, err := readBytesToStr(&r, len, false)
		if err != nil {
			return err
		}
		tx.TransferType = TxType
	}

	tx.Input = "0x"
	return nil
}

func (tx *Eth_token_transaction) Deserialize(data []byte) error {
	var r bytes.Buffer
	_, err := r.Write(data)
	if err != nil {
		return err
	}

	// blockNumber
	len, err := readByte(&r)
	if err != nil {
		return err
	}
	blockNumber, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.BlockNumber = blockNumber

	// GasPrice

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	GasPrice, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.GasPrice = GasPrice

	// Hash

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Hash, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	tx.TransactionHash = Hash

	// TransactionIndex

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	TransactionIndex, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.TransactionIndex = TransactionIndex

	// GasUsed

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	GasUsed, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.GasUsed = GasUsed

	// ContractAddress

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	ContractAddress, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	tx.Address = ContractAddress

	//Timestamp

	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Timestamp, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.TimeStamp = Timestamp

	//Topics
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	tLen, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	l, err := strconv.Atoi(tLen)
	if err != nil {
		return err
	}
	var topics []interface{}
	for i := 0; i < l; i++ {
		len, err = readByte(&r)
		if err != nil {
			return err
		}
		topic, err := readBytesToHexStr(&r, len)
		if err != nil {
			return err
		}
		topics = append(topics, topic)
	}
	tx.Topics = topics

	//Data
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	d, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	tx.Data = d

	// LogIndex
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	LogIndex, err := readBytesToStr(&r, len, true)
	if err != nil {
		return err
	}
	tx.LogIndex = LogIndex

	return nil
}

func (txt *Eth_token_transaction) Serialize() []byte {
	var b bytes.Buffer

	b.WriteByte(byte(len(txt.BlockNumber) - 2))
	b.Write([]byte(txt.BlockNumber[2:]))

	b.WriteByte(byte(len(txt.GasPrice) - 2))
	b.Write([]byte(txt.GasPrice[2:]))

	hash := decodeHexToByte(txt.TransactionHash)
	b.WriteByte(byte(len(hash)))
	b.Write(hash)

	b.WriteByte(byte(len(txt.TransactionIndex) - 2))
	b.Write([]byte(txt.TransactionIndex[2:]))

	b.WriteByte(byte(len(txt.GasUsed) - 2))
	b.Write([]byte(txt.GasUsed[2:]))

	contractAddress := decodeHexToByte(txt.Address)
	b.WriteByte(byte(len(contractAddress)))
	b.Write(contractAddress)

	b.WriteByte(byte(len(txt.TimeStamp) - 2))
	b.Write([]byte(txt.TimeStamp[2:]))

	b.WriteByte(byte(len(strconv.Itoa(len(txt.Topics)))))
	b.Write([]byte(strconv.Itoa(len(txt.Topics))))
	for _, topic := range txt.Topics {
		t := decodeHexToByte(topic.(string))
		b.WriteByte(byte(len(t)))
		b.Write(t)
	}

	data := decodeHexToByte(txt.Data)
	b.WriteByte(byte(len(data)))
	b.Write(data)

	b.WriteByte(byte(len(txt.LogIndex) - 2))
	b.Write([]byte(txt.LogIndex[2:]))

	return b.Bytes()
}

type TransactionTokenHistorySorter []Eth_token_transaction

func (a TransactionTokenHistorySorter) Len() int           { return len(a) }
func (a TransactionTokenHistorySorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TransactionTokenHistorySorter) Less(i, j int) bool { return a[i].TimeStamp < a[j].TimeStamp }

func (tx *Eth_transaction) Serialize() []byte {
	var b bytes.Buffer
	blockHash := decodeHexToByte(tx.BlockHash)
	b.WriteByte(byte(len(blockHash)))
	b.Write(blockHash)

	b.WriteByte(byte(len(tx.BlockNumber) - 2))
	b.Write([]byte(tx.BlockNumber[2:]))

	from := decodeHexToByte(tx.From)
	b.WriteByte(byte(len(from)))
	b.Write(from)

	b.WriteByte(byte(len(tx.Gas) - 2))
	b.Write([]byte(tx.Gas[2:]))

	to := decodeHexToByte(tx.To)
	b.WriteByte(byte(len(to)))
	b.Write(to)

	b.WriteByte(byte(len(tx.GasPrice) - 2))
	b.Write([]byte(tx.GasPrice[2:]))

	hash := decodeHexToByte(tx.Hash)
	b.WriteByte(byte(len(hash)))
	b.Write(hash)

	b.WriteByte(byte(len(tx.Nonce) - 2))
	b.Write([]byte(tx.Nonce[2:]))

	b.WriteByte(byte(len(tx.TransactionIndex) - 2))
	b.Write([]byte(tx.TransactionIndex[2:]))

	b.WriteByte(byte(len(tx.Value) - 2))
	b.Write([]byte(tx.Value[2:]))

	b.WriteByte(byte(len(tx.GasUsed) - 2))
	b.Write([]byte(tx.GasUsed[2:]))

	b.WriteByte(byte(len(tx.CumulativeGasUsed) - 2))
	b.Write([]byte(tx.CumulativeGasUsed[2:]))

	b.WriteByte(byte(len(tx.IsError)))
	b.Write([]byte(tx.IsError))

	contractAddress := decodeHexToByte(tx.ContractAddress)
	b.WriteByte(byte(len(contractAddress)))
	b.Write(contractAddress)

	b.WriteByte(byte(len(tx.Timestamp) - 2))
	b.Write([]byte(tx.Timestamp[2:]))

	if config.Conf.Eth.SideChain {
		b.WriteByte(byte(len(tx.TransferType)))
		b.Write([]byte(tx.TransferType))
	}

	return b.Bytes()
}

func init() {
	if config.Conf.Eth.Enable {
		var err error
		user.Current()
		user, err := user.Current()
		if err != nil {
			fmt.Printf("Error init level db %s", err.Error())
			os.Exit(-1)
		}
		homeDir := user.HomeDir
		dir := homeDir + levelDbPath
		_, err = os.Stat(dir)
		if err != nil {
			if !os.IsExist(err) {
				err = os.MkdirAll(dir, 0755)
				if err != nil {
					fmt.Printf("Error init level db %s", err.Error())
					os.Exit(-1)
				}
			}
		}
		db, err := leveldb.OpenFile(dir, &opt.Options{
			Filter: filter.NewBloomFilter(10),
		})
		if err != nil {
			fmt.Printf("Error init level db %s", err.Error())
			os.Exit(-1)
			return
		}

		le = &level{
			l:    db,
			path: levelDbPath,
		}

	}
}

//Sync sync chain data
func SyncEth() {
	go func() {
		for {
			le.b = new(leveldb.Batch)
			if err := doSyncEth(le); err != nil {
				log.Infof("Sync ETH Height Error : %v", err.Error())
			} else {
				err := le.l.Write(le.b, nil)
				if err != nil {
					log.Errorf(" Error Syncing From Height : %d", le.currHeight+1)
				}
			}
			<-time.After(time.Millisecond * 1000)
		}
	}()
}

func doSyncEth(le *level) error {
	if le.currHeight == 0 {
		data, err := le.l.Get([]byte{byte(curr_height_prefix)}, nil)
		if data == nil || err != nil {
			le.currHeight = -1
		} else {
			tmpInt, err := strconv.Atoi(string(data))
			if err != nil {
				return err
			}
			le.currHeight = int64(tmpInt)
		}
	}

	if le.currHeight < config.Conf.Eth.StartHeight && config.Conf.Eth.StartHeight != 0 {
		le.currHeight = config.Conf.Eth.StartHeight - 1
	}

	var resp map[string]interface{}
	var err error
	resp, err = Post(config.Conf.Eth.Endpoint, `{"jsonrpc":"2.0","method":"eth_blockNumber","params": [],"id":1}`)
	if err != nil {
		return err
	}

	hexHeight, ok := resp["result"]
	var unexpected error = nil
	var waitSize int
	if ok {
		height, err := hexutil.DecodeUint64(hexHeight.(string))
		if err != nil {
			return err
		}
		if le.currHeight >= int64(height) {
			return nil
		}
		gap := int(height) - int(le.currHeight)
		if gap >= config.Conf.Eth.BatchSize {
			waitSize = config.Conf.Eth.BatchSize
		} else {
			waitSize = gap
		}
		log.Infof("Syncing ETH , Height From %d To %d \n", le.currHeight+1, le.currHeight+int64(waitSize)+1)
		le.waitGroup.Add(waitSize)
		count := 0
		for curr := le.currHeight; curr <= int64(height); {
			go func() {
				time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
				atomic.AddInt64(&curr, 1)
				err = handleHeightEth(int(curr))
				if err != nil {
					unexpected = err
				}
				le.waitGroup.Done()
			}()
			count++
			if count%waitSize == 0 {
				break
			}
		}
	}
	le.waitGroup.Wait()
	if unexpected == nil {
		le.currHeight += int64(waitSize)
		le.b.Put([]byte{byte(curr_height_prefix)}, []byte(strconv.Itoa(int(le.currHeight))))
	}
	return unexpected
}

func handleHeightEth(curr int) error {
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
	var keys [][]byte
	var values [][]byte
	for index, txv := range txArr {
		transaction := txv.(map[string]interface{})
		v := Eth_transaction{}
		Map2Struct(transaction, &v)
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
			isError = "1"
		}
		cumulativeGasUsed := receipt["cumulativeGasUsed"]
		contractAddress := receipt["contractAddress"]
		if contractAddress == nil {
			contractAddress = "0x0000000000000000000000000000000000000000"
		}
		v.Timestamp = timestamp.(string)
		v.ContractAddress = contractAddress.(string)
		v.CumulativeGasUsed = cumulativeGasUsed.(string)
		v.GasUsed = gasUsed.(string)
		v.IsError = isError
		v.Input = "0x"

		if config.Conf.Eth.SideChain {
			v.TransferType = sidechain_eth_transfer
		}

		if v.To == "0x0000000000000000000000000000000000000000" {
			logs, ok := receipt["logs"].([]interface{})
			if ok && len(logs) == 1 {
				l := logs[0]
				rl := l.(map[string]interface{})
				topics, ok := rl["topics"].([]interface{})
				if ok && len(topics) >= 5 && topics[0].(string) == "0x09f15c376272c265d7fcb47bf57d8f84a928195e6ea156d12f5a3cd05b8fed5a" {
					v.To = GetEthAddress(topics[3].(string))
					v.Value = GetEthValue(topics[4].(string))
					v.TransferType = sidechain_crossChain_deposit
				}
			}
		}

		if strings.ToUpper(v.To) == strings.ToUpper("0xC445f9487bF570fF508eA9Ac320b59730e81e503") {
			v.TransferType = sidechain_crossChain_withdraw
		}

		// From
		var keyFrom bytes.Buffer
		keyFrom.Write([]byte{byte(eth_history_prefix)})
		keyFrom.Write(decodeHexToByte(v.From))
		keyFrom.WriteRune(rune(curr))
		keyFrom.WriteRune(rune(index))
		val := v.Serialize()
		keys = append(keys, keyFrom.Bytes())
		values = append(values, val)

		var keyTo bytes.Buffer
		// TO
		if v.From != v.To {
			keyTo.Write([]byte{byte(eth_history_prefix)})
			keyTo.Write(decodeHexToByte(v.To))
			keyTo.WriteRune(rune(curr))
			keyTo.WriteRune(rune(index))
			keys = append(keys, keyTo.Bytes())
			values = append(values, val)
		}

		if isError == "0" {
			logs, ok := receipt["logs"].([]interface{})
			if ok {
				for _, l := range logs {
					rl := l.(map[string]interface{})
					topics, ok := rl["topics"].([]interface{})
					if ok && len(topics) >= 3 && topics[0].(string) == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
						ett := Eth_token_transaction{}
						Map2Struct(rl, &ett)
						ett.GasUsed = v.GasUsed
						ett.GasPrice = v.GasPrice
						ett.TimeStamp = v.Timestamp
						fmt.Printf(" txid %v \n", ett.TransactionHash)
						// From
						var keyFromToken bytes.Buffer
						keyFromToken.Write([]byte{byte(eth_token_history_prefix)})
						keyFromToken.Write(decodeHexToByte(GetEthAddress(topics[1].(string))))
						keyFromToken.WriteRune(rune(curr))
						keyFromToken.WriteRune(rune(index))
						val := ett.Serialize()
						keys = append(keys, keyFromToken.Bytes())
						fmt.Printf(" address from %v \n", GetEthAddress(topics[1].(string)))
						values = append(values, val)

						var keyToToken bytes.Buffer
						// TO
						if topics[1].(string) != topics[2].(string) {
							keyToToken.Write([]byte{byte(eth_token_history_prefix)})
							keyToToken.Write(decodeHexToByte(GetEthAddress(topics[2].(string))))
							keyToToken.WriteRune(rune(curr))
							keyToToken.WriteRune(rune(index))
							keys = append(keys, keyToToken.Bytes())
							fmt.Printf(" address to %v \n", GetEthAddress(topics[2].(string)))
							values = append(values, val)
						}
					}
				}
			}
		}
	}
	le.m.Lock()
	for i, k := range keys {
		le.b.Put(k, values[i])
	}
	le.m.Unlock()
	return nil
}

func GetEthHistory(addr string) ([]Eth_transaction, error) {
	var buf bytes.Buffer
	buf.Write([]byte{byte(eth_history_prefix)})
	buf.Write(decodeHexToByte(addr))
	iter := le.l.NewIterator(util.BytesPrefix(buf.Bytes()), nil)
	ret := make(TransactionHistorySorter, 0)
	for iter.Next() {
		var v Eth_transaction
		value := iter.Value()
		err := v.Deserialize(value)
		if err != nil {
			return nil, err
		}
		bn, err := strconv.Atoi(v.BlockNumber)
		if err != nil {
			return nil, err
		}
		v.Confirmations = strconv.Itoa(int(le.currHeight) - bn)
		ret = append(ret, v)
	}
	defer iter.Release()
	sort.Sort(ret)
	return ret, nil
}

func GetEthTokenLogs(from string, to string) ([]Eth_token_transaction, error) {
	spend, err := doGetEthTokenLogs(from)
	if err != nil {
		return nil, err
	}
	var income TransactionTokenHistorySorter
	if from != to {
		income, err = doGetEthTokenLogs(to)
		if err != nil {
			return nil, err
		}
	}
	ret := make(TransactionTokenHistorySorter, len(spend)+len(income))
	copy(ret, spend)
	copy(ret[len(spend):], income)
	sort.Sort(ret)
	return ret, nil
}

func doGetEthTokenLogs(addr string) (TransactionTokenHistorySorter, error) {
	var buf bytes.Buffer
	buf.Write([]byte{byte(eth_token_history_prefix)})
	buf.Write(decodeHexToByte(addr))
	iter := le.l.NewIterator(util.BytesPrefix(buf.Bytes()), nil)
	ret := make(TransactionTokenHistorySorter, 0)
	for iter.Next() {
		var v Eth_token_transaction
		value := iter.Value()
		err := v.Deserialize(value)
		if err != nil {
			return nil, err
		}
		ret = append(ret, v)
	}
	defer iter.Release()
	return ret, nil
}

func decodeHexToDecimal(str string) string {
	if len(str) == 0 {
		return ""
	}
	desc, _ := strconv.ParseUint(str[2:], 16, 64)
	return strconv.Itoa(int(desc))
}

func decodeHexToByte(str string) []byte {
	if len(str) == 0 {
		return nil
	}
	b, err := hex.DecodeString(str[2:])
	if err != nil {
		log.Errorf("Error decodeHexToByte %s", str)
		return nil
	}
	return b
}

func readByte(r io.Reader) (int, error) {
	buf := make([]byte, 1)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return -1, err
	}
	var len byte
	binary.Read(bytes.NewBuffer(buf), binary.BigEndian, &len)
	if err != nil {
		return -1, err
	}
	return int(len), nil
}

func readBytesToHexStr(r io.Reader, len int) (string, error) {
	if len == 0 {
		return "0x", nil
	}
	buf := make([]byte, len)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return "", err
	}
	return "0x" + hex.EncodeToString(buf), nil
}

func readBytesToStr(r io.Reader, len int, prefix bool) (string, error) {
	if len == 0 {
		return "0x", nil
	}
	buf := make([]byte, len)
	_, err := io.ReadFull(r, buf)
	if err != nil {
		return "", err
	}
	if prefix {
		return "0x" + string(buf), nil
	} else {
		return string(buf), nil
	}
}
