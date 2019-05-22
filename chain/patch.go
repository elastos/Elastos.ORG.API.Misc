package chain

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"github.com/elastos/Elastos.ORG.API.Misc/tools"
	"strconv"
	"time"
)

var height int = 0
var done bool = false

//Patch patch is used to do smooth upgrade, for the sake of sync data from start
func Patch() {
	for {
		tx, err := dba.Begin()
		if err = doSyncPatch(tx); err != nil {
			log.Infof("Sync Height Error : %v \n", err.Error())
			tx.Rollback()
		} else {
			log.Infof("Patch Success")
			tx.Commit()
		}
		if done {
			return
		}
		<-time.After(1 * time.Second)
	}
}

func doSyncPatch(tx *sql.Tx) error {

	resp, err := get("http://" + config.Conf.Ela.Restful + BlockHeight)

	if err != nil {
		return err
	}
	storeHeight := height
	chainHeight, ok := resp["Result"]
	if int(chainHeight.(float64)) <= height {
		done = true
		return nil
	}
	if ok {
		if storeHeight == int(chainHeight.(float64)) {
			return nil
		}
		l, err := dba.Query(`SELECT * FROM information_schema.COLUMNS WHERE TABLE_NAME = 'chain_block_transaction_history' AND COLUMN_NAME = 'publicKey'`)
		if err != nil {
			return err
		}
		if l.Len() == 0 {
			_, err := tx.Exec(`alter table chain_block_transaction_history add publicKey varchar(66) default null comment 'only spend transaction got public key'`)
			if err != nil {
				return err
			}
		}
		for curr := storeHeight + 1; curr <= int(chainHeight.(float64)); curr++ {
			height = curr
			log.Info("patch height :", curr)
			err = handleHeightPatch(curr, tx)
			if err != nil {
				return err
			}
			if curr%5000 == 0 {
				return nil
			}
		}
	}

	return nil
}

func handleHeightPatch(curr int, tx *sql.Tx) error {
	resp, err := get("http://" + config.Conf.Ela.Restful + BlockDetail + strconv.FormatInt(int64(curr), 10))
	if err != nil {
		return err
	}
	r, ok := (resp["Result"].(map[string]interface{}))
	if !ok {
		return errors.New("illegal Height")
	}
	txArr := r["tx"].([]interface{})
	if len(txArr) == 0 {
		return nil
	}
	stmt, err := tx.Prepare("update chain_block_transaction_history set publicKey = ? where txid = ? and `type` = 'spend' and address = ?")
	if err != nil {
		return err
	}

	for _, v := range txArr {
		vm := v.(map[string]interface{})
		txid := vm["txid"].(string)
		t := vm["type"].(float64)
		if int(t) != CoinBase {
			p := vm["programs"].([]interface{})
			for _, v := range p {
				pm := v.(map[string]interface{})
				code := pm["code"].(string)
				publicKeyByte, _ := hex.DecodeString(code)
				publicKeyStr := hex.EncodeToString(publicKeyByte[1:34])
				address, err := tools.GetAddress(publicKeyStr)
				_, err = stmt.Exec(publicKeyStr, txid, address)
				if err != nil {
					return err
				}
			}
		}
	}

	stmt.Close()
	return nil
}

func init() {
	log.InitLog(0, 50)
	go Patch()
}
