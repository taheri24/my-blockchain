package mux

import (
	"fmt"
	"net/http"
	"strings"

	"taheri24.ir/blockchain/internal/handlers"
)

func GroupFn(parent *http.ServeMux) func(prefix string, sub *http.ServeMux) {
	return func(prefix string, sub *http.ServeMux) {
		parent.HandleFunc(fmt.Sprintf("%s/*"), func(w http.ResponseWriter, r *http.Request) {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, prefix)
			sub.ServeHTTP(w, r)
		})
	}
}

func Api() *http.ServeMux {
	w, a, scMux, bMux := Wallet(), Auth(), SmartContracts(), Blocks()
	m := http.NewServeMux()
	group := GroupFn(m)
	group("/auth", a)
	group("/smart-contract", scMux)
	group("/blocks", bMux)
	group("/wallets", w)

	return m
}

func SmartContracts() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/{id}", handlers.GetSmartContractByID)
	m.HandleFunc("/", handlers.CreateSmartContract)
	m.HandleFunc("/{id}/{action}", handlers.DoActionOnSmartContract)
	return m
}

func Blocks() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("POST /", handlers.PostToken)
	return m
}

func Wallet() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("POST /", handlers.CreateWallet)
	m.HandleFunc("GET /{id}", handlers.GetWallet)

	return m

}
