package handlers

import (
	"net/http"

	"taheri24.ir/blockchain/utils"
)

var GetMe http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	utils.ServeWithJsonFile(w, "./mocks/me.json")
}

var LogIn http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	utils.ServeWithJsonFile(w, "./mocks/me.json")
}
