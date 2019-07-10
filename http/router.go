package http

import (
	"github.com/elastos/Elastos.ORG.API.Misc/config"
	"github.com/gorilla/mux"
	"log"
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
		"/api/1/pubkey/{addr}": {
			"GET": getPublicKey,
		},

		//ela dpos
		"/api/1/dpos/producer/{producer}": {
			"GET": producerStatistic,
		},
		"/api/1/dpos/address/{address}": {
			"GET": voterStatistic,
		},
		"/api/1/dpos/rank/height/{height}": {
			"GET": producerRankByHeight,
		},
		"/api/1/dpos/vote/height/{height}": {
			"GET": totalVoteByHeight,
		},
		"/api/1/dpos/confirmed/height/{height}": {
			"GET": confirmedDetailByHeight,
		},

		// post
		"/api/1/dpos/transaction/producer": {
			"POST": getProducerByTxs,
		},


		//heartbeat
		"/api/1/ping": {
			"GET": ping,
		},

		//sync checking
		"/api/1/history/checking/sync": {
			"GET": syncChecking,
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
	router.Use(cors)
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
		if config.Conf.EnableCors {
			w.Header().Add("Access-Control-Allow-Headers", "Authorization, Origin, X-Requested-With, Content-Type, Accept")
			w.Header().Set("Access-Control-Allow-Methods","GET, POST, OPTIONS, HEAD")
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
	})
}
