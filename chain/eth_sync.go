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
	"sort"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var (
	levelDbPath = "/.misc/eth"
	le          *level
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
	curr_height_prefix key_prefix = 0x01
	eth_history_prefix key_prefix = 0x02
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
}

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

	tx.Input = "0x"
	return nil
}

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

	return b.Bytes()
}

func init() {
	if config.Conf.Eth.Enable {
		var err error
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("Error init level db %s", err.Error())
			os.Exit(-1)
		}
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
			b:    new(leveldb.Batch),
		}
	}
}

//Sync sync chain data
func SyncEth() {
	go func() {
		for {
			if err := doSyncEth(le); err != nil {
				log.Infof("Sync ETH Height Error : %v \n", err.Error())
			} else {
				le.l.Write(le.b, nil)
			}
			<-time.After(time.Millisecond * 10000)
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
			le.waitGroup.Add(config.Conf.Eth.BatchSize)
			log.Infof("Syncing ETH , Height From %d To %d \n", le.currHeight+1, le.currHeight+int64(config.Conf.Eth.BatchSize)+1)
		} else {
			le.waitGroup.Add(gap)
			log.Infof("Syncing ETH , Height From %d To %d \n", le.currHeight+1, le.currHeight+int64(gap)+1)
		}
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
			if count%config.Conf.Eth.BatchSize == 0 {
				break
			}
		}
	}
	le.waitGroup.Wait()
	le.currHeight += int64(config.Conf.Eth.BatchSize)
	le.b.Put([]byte{byte(curr_height_prefix)}, []byte(strconv.Itoa(int(le.currHeight))))
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
		ret = append(ret, v)
	}
	defer iter.Release()
	sort.Sort(ret)
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
