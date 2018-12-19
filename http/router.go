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