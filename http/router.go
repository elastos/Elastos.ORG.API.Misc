package http

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	// path|method|handler
	routers = map[string]map[string]http.HandlerFunc{

		//ela
		"/api/1/history/{addr}": {
			"GET": history,
		},
		"/api/1/did/{did}/{key}": {
			"GET": searchKey,
		},
		//ela dpos
		"/api/1/dpos/producer/{producer}": {
			"GET": producerStatistic,
		},
		"/api/1/dpos/address/{address}": {
			"GET": voterStatistic,
		},
		"/api/1/dpos/reward/height/{height}": {
			"GET": rewardByHeight,
		},
		"/api/1/dpos/rank/height/{height}": {
			"GET": producerRankByHeight,
		},
		"/api/1/dpos/vote/height/{height}": {
			"GET": totalVoteByHeight,
		},

		//heartbeat
		"/api/1/ping": {
			"GET": ping,
		},

		//ela frontend
		"/api/1/list": {
			"Get": list,
		},

		//btc
		"/api/1/btc/block/height": {
			"Get": getBtcBlockHeight,
		},
		"/api/1/btc/transaction/{txid}": {
			"Get": getBtcTransaction,
		},
		"/api/1/btc/balance/{addr}": {
			"Get": getBtcBalance,
		},
		"/api/1/btc/detail/block/{height}": {
			"Get": getBtcBlock,
		},

		//CoinMarketCap
		"/api/1/cmc": {
			"Get": getCmcPrice,
		},
	}
	router = mux.NewRouter()
)

func init() {
	for p, r := range routers {
		for m, h := range r {
			router.HandleFunc(p, h).Methods(m)
		}
	}
}
