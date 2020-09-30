package chain

import (
	"bytes"
	"errors"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/syndtr/goleveldb/leveldb/util"
	"os"
	"sort"
	"strconv"
)

type Erc20TokenSorter []Erc20Token

func (a Erc20TokenSorter) Len() int           { return len(a) }
func (a Erc20TokenSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Erc20TokenSorter) Less(i, j int) bool { return a[i].Height > a[j].Height }

type Erc20Token struct {
	Address string `json:"address"`
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Decimal string `json:"decimal"`
	Height  string `json:"height"`
}

func (token *Erc20Token) Deserialize(data []byte) error {
	var r bytes.Buffer
	_, err := r.Write(data)
	if err != nil {
		return err
	}

	//Address
	len, err := readByte(&r)
	if err != nil {
		return err
	}
	ContractAddress, err := readBytesToHexStr(&r, len)
	if err != nil {
		return err
	}
	token.Address = string(ContractAddress)

	// Symbol
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Symbol, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	token.Symbol = Symbol

	// Name
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Name, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	token.Name = Name

	// Decimal
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Decimal, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	token.Decimal = Decimal

	// Height
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Height, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	token.Height = Height

	return nil
}

func (token *Erc20Token) Serialize() []byte {
	var b bytes.Buffer
	contractAddress := decodeHexToByte(token.Address)
	b.WriteByte(byte(len(contractAddress)))
	b.Write(contractAddress)

	symbol := []byte(token.Symbol)
	b.WriteByte(byte(len(symbol)))
	b.Write(symbol)

	name := []byte(token.Name)
	b.WriteByte(byte(len(name)))
	b.Write(name)

	b.WriteByte(byte(len(token.Decimal)))
	b.Write([]byte(token.Decimal))

	b.WriteByte(byte(len(token.Height)))
	b.Write([]byte(token.Height))

	return b.Bytes()
}

func Call(contract string, height int) (Erc20Token, error) {
	log.Info("contract address ", contract)
	client, err := ethclient.Dial(config.Conf.Eth.Endpoint)
	if err != nil {
		log.Error("Unable to connect to eth")
		os.Exit(-1)
	}

	address := common.HexToAddress(contract)
	erc20, err := NewErc20(address, client)
	if err != nil {
		log.Errorf("Unable to create eth exchange contract instance, wrong contract address ? , %s ", err.Error())
		os.Exit(-1)
	}
	name, err := erc20.Name(nil)
	if err != nil {
		return Erc20Token{}, errors.New("Name fetching failed " + err.Error())
	}

	symbol, err := erc20.Symbol(nil)
	if err != nil {
		return Erc20Token{}, errors.New("symbol fetching failed " + err.Error())
	}

	decimal, err := erc20.Decimals(nil)
	if err != nil {
		return Erc20Token{}, errors.New("decimal fetching failed " + err.Error())
	}

	return Erc20Token{
		Name:    name,
		Decimal: strconv.Itoa(int(decimal)),
		Symbol:  symbol,
		Address: contract,
		Height:  strconv.Itoa(int(height)),
	}, nil
}

func isTokenExist(addr string) bool {
	var buf bytes.Buffer
	buf.Write([]byte{byte(eth_token_list_prefix)})
	buf.Write(decodeHexToByte(addr))
	iter := le.l.NewIterator(util.BytesPrefix(buf.Bytes()), nil)
	for iter.Next() {
		return true
	}
	defer iter.Release()
	return false
}

func GetTokenList() ([]Erc20Token, error) {
	var buf bytes.Buffer
	buf.Write([]byte{byte(eth_token_list_prefix)})
	iter := le.l.NewIterator(util.BytesPrefix(buf.Bytes()), nil)
	ret := make(Erc20TokenSorter, 0)
	for iter.Next() {
		var v Erc20Token
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

