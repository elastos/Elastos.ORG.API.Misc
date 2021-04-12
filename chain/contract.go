package chain

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
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
	Address         string `json:"address"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        string `json:"decimals"`
	Height          string `json:"height"`
	Description     string `json:"description"`
	DefaultGasLimit string `json:"defaultGasLimit"`
	DefaultGasPrice string `json:"defaultGasPrice"`
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

	// Decimals
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Decimals, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	token.Decimals = Decimals

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

	// Description
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	Desc, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	token.Description = Desc

	// DefaultGasLimit
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	DefaultGasLimit, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	token.DefaultGasLimit = DefaultGasLimit

	// DefaultGasPrice
	len, err = readByte(&r)
	if err != nil {
		return err
	}
	DefaultGasPrice, err := readBytesToStr(&r, len, false)
	if err != nil {
		return err
	}
	token.DefaultGasPrice = DefaultGasPrice
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

	b.WriteByte(byte(len(token.Decimals)))
	b.Write([]byte(token.Decimals))

	b.WriteByte(byte(len(token.Height)))
	b.Write([]byte(token.Height))

	b.WriteByte(byte(len(token.Description)))
	b.Write([]byte(token.Description))

	b.WriteByte(byte(len(token.DefaultGasLimit)))
	b.Write([]byte(token.DefaultGasLimit))

	b.WriteByte(byte(len(token.DefaultGasPrice)))
	b.Write([]byte(token.DefaultGasPrice))
	return b.Bytes()
}

func Call(contract string, height int) (Erc20Token, error) {
	fmt.Println("contract address ", contract)
	client, err := ethclient.Dial(config.Conf.Eth.Endpoint)
	if err != nil {
		fmt.Println("Unable to connect to eth")
		os.Exit(-1)
	}

	address := common.HexToAddress(contract)
	erc20, err := NewErc20(address, client)
	if err != nil {
		fmt.Printf("Unable to create eth exchange contract instance, wrong contract address ? , %s ", err.Error())
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

	decimals, err := erc20.Decimals(nil)
	if err != nil {
		return Erc20Token{}, errors.New("decimal fetching failed " + err.Error())
	}

	desc, err := erc20.Description(nil)
	if err != nil {
		fmt.Println("No description")
		desc = ""
	}

	var defaultGasLimit string
	gasLimit, err := erc20.DefaultGasLimit(nil)
	if err != nil {
		fmt.Println("No defaultGasLimit ", err.Error())
		defaultGasLimit = "0"
	} else {
		defaultGasLimit = gasLimit.String()
	}

	var defaultGasPrice string
	gasPrice, err := erc20.DefaultGasPrice(nil)
	if err != nil {
		fmt.Println("No defaultGasLimit", err.Error())
		defaultGasPrice = "0"
	} else {
		defaultGasPrice = gasPrice.String()
	}

	return Erc20Token{
		Name:            name,
		Decimals:        strconv.Itoa(int(decimals)),
		Symbol:          symbol,
		Description:     desc,
		DefaultGasPrice: defaultGasPrice,
		DefaultGasLimit: defaultGasLimit,
		Address:         contract,
		Height:          strconv.Itoa(int(height)),
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
