package chain

import (
	"database/sql"
	"errors"
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/elastos/Elastos.ORG.API.Misc/log"
	"strconv"
)

var height int = 350000

//Patch patch is used to do smooth upgrade, for the sake of sync data from start
func Patch(){
	tx, err := dba.Begin()
	if err = doSyncPatch(tx); err != nil {
		log.Infof("Sync Height Error : %v \n", err.Error())
		tx.Rollback()
	} else {
		log.Infof("Patch Success")
		tx.Commit()
	}
}

func doSyncPatch(tx *sql.Tx) error {

	resp, err := get("http://" + config.Conf.Ela.Restful + BlockHeight)

	if err != nil {
		return err
	}

	storeHeight := height

	chainHeight, ok := resp["Result"]
	if ok {
		if storeHeight == int(chainHeight.(float64)) {
			return nil
		}
		for curr := storeHeight + 1; curr <= int(chainHeight.(float64)); curr++ {
			println(curr)
			err = handleHeightPatch(curr, tx)
			if err != nil {
				return err
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
	for _, v := range txArr {
		stmt, err := tx.Prepare("update chain_block_transaction_history set txType = ? where txid = ?")
		if err != nil {
			return err
		}
		vm := v.(map[string]interface{})
		txid := vm["txid"].(string)
		t := vm["type"].(float64)
		if int(t) == RegisterProducer || int(t) == CancelProducer || int(t) == UpdateProducer || int(t) == ReturnDepositCoin || int(t) == ActivateProducer {
			_, err := stmt.Exec(txTypeMap[int(t)],txid)
			if err != nil {
				return err
			}
		}
		stmt.Close()
	}

	return nil
}

func init()  {
	log.InitLog(0, 50)
	Patch()
}