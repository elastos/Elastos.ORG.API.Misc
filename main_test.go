package main

import (
	"bytes"
	"fmt"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcutil"
	"testing"
)

var client, _ = rpcclient.New(&rpcclient.ConnConfig{
	HTTPPostMode: true,
	DisableTLS:   true,
	Host:         "127.0.0.1:18332",
	User:         "clark",
	Pass:         "DrGhlxKuqP02m47TnDRAeZqir6Gt5V0secCHajUKW-0=",
	}, nil)

func Test_main(t *testing.T) {

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

	//addr, err := client.GetNewAddress("Test")

	//println(addr.EncodeAddress())
	addr , _ := btcutil.DecodeAddress("2NBN7wXptuq4bKNiKeabGUdhckLYURJB4jm",&chaincfg.TestNet3Params)
	//addr1 , _ := btcutil.DecodeAddress("mqeRnQdfTsbELix9vXUZhTReMifqUgfJaD",&chaincfg.TestNet3Params)

	addrs := []btcutil.Address{addr}

	utxo ,err := client.ListUnspentMinMaxAddresses(0,99999,addrs)
	inputs := btcjson.TransactionInput{utxo[0].TxID,utxo[0].Vout}
	amt , _:= btcutil.NewAmount(.0001)
	//amt1 , _:= btcutil.NewAmount(.0002)
	tx, err := client.CreateRawTransaction([]btcjson.TransactionInput{inputs},map[btcutil.Address]btcutil.Amount{
		//addr:amt,
		addr:amt,
	},nil)
	tx , b , err := client.SignRawTransaction(tx)
	if err != nil {
		println(err.Error())
	}
	buf := new(bytes.Buffer)
	txid , err := client.SendRawTransaction(tx,true)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("%v %v %v %v \n",tx,buf,b,txid)

	//b, err := hex.DecodeString("662b775b5b4eb8b581e59ac9d905921b8452bec765e0a47a7de6f75fd5695594")
	//h := new(chainhash.Hash)
	//
	//h.SetBytes(tools.ReverseBytes(b))
	rawTx, err := client.GetRawTransactionVerbose(txid)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("%v\n", rawTx)

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


	// list accounts
	utxos, err := client.ListUnspentMinMaxAddresses(1,99999999,addrs)
	for _ , v := range utxos {
		fmt.Printf("%v\n",v)
	}

	//amt , err := client.GetRawTransactionVerbose()
	//if err != nil {
	//	println(err.Error())
	//}
	//println(amt)
	txids , err := client.GetRawMempool()
	if err != nil {
		println(err.Error())
	}
	for _ , txid := range txids {
		println(txid.String())
	}
}
