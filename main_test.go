package main

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	"log"
	"testing"
)

func Test_main(t *testing.T) {
	// create new client instance
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         "127.0.0.1:18332",
		User:         "clark",
		Pass:         "DrGhlxKuqP02m47TnDRAeZqir6Gt5V0secCHajUKW-0=",
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}
	height, err := client.GetBlockCount()
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("Height is %d\n", height)

	//mempool
	//txids , err := client.GetRawMempool();
	//for _ , txid := range txids {
	//	m , err := client.GetMempoolEntry(txid.String())
	//	if err != nil {
	//		println(err.Error())
	//	}
	//	fmt.Printf("%v\n",m)
	//}

	//addr , _ := btcutil.DecodeAddress("2NBN7wXptuq4bKNiKeabGUdhckLYURJB4jm",&chaincfg.TestNet3Params)
	//addrs := []btcutil.Address{addr}
	//utxo ,err := client.ListUnspentMinMaxAddresses(0,99999,addrs)
	//inputs := btcjson.TransactionInput{utxo[0].TxID,utxo[0].Vout}
	//amt , _:= btcutil.NewAmount(.00099)
	//tx, err := client.CreateRawTransaction([]btcjson.TransactionInput{inputs},map[btcutil.Address]btcutil.Amount{
	//	addr:amt,
	//},nil)
	//tx , b , err := client.SignRawTransaction(tx)
	//buf := new(bytes.Buffer)
	//txid , err := client.SendRawTransaction(tx,true)
	//fmt.Printf("%v %v %v %v \n",tx,buf,b,txid)

	b, err := hex.DecodeString("c094d4fb7b928d614682e2240a3bfd1b0bf9c63d853f5d2a4a5b1605657cdb0c")
	h := new(chainhash.Hash)

	h.SetBytes(tools.ReverseBytes(b))
	rawTx, err := client.GetRawTransaction(h)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("%v\n", rawTx)

	// list accounts
	accounts, err := client.ListAccounts()
	if err != nil {
		log.Fatalf("error listing accounts: %v", err)
	}
	// iterate over accounts (map[string]btcutil.Amount) and write to stdout
	for label, amount := range accounts {
		log.Printf("%s: %s", label, amount)
	}

	//// prepare a sendMany transaction
	//receiver1, err := btcutil.DecodeAddress("1someAddressThatIsActuallyReal", &chaincfg.MainNetParams)
	//if err != nil {
	//	log.Fatalf("address receiver1 seems to be invalid: %v", err)
	//}
	//receiver2, err := btcutil.DecodeAddress("1anotherAddressThatsPrettyReal", &chaincfg.MainNetParams)
	//if err != nil {
	//	log.Fatalf("address receiver2 seems to be invalid: %v", err)
	//}
	//receivers := map[btcutil.Address]btcutil.Amount{
	//	receiver1: 42,  // 42 satoshi
	//	receiver2: 100, // 100 satoshi
	//}
	//
	//// create and send the sendMany tx
	//txSha, err := client.SendMany("some-account-label-from-which-to-send", receivers)
	//if err != nil {
	//	log.Fatalf("error sendMany: %v", err)
	//}
	//log.Printf("sendMany completed! tx sha is: %s", txSha.String())
}
