package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"taheri24.ir/blockchain/internal/cmd/wallet"
)

type CreateWalletRequest struct {
	PublicKeyBase64     string `json:"publicBase64"`
	publicKey           []byte
	PersonalInformation struct {
		ID                  string
		FirstName, LastName string
	} `json:"personaInfo"`
	AddressInfo struct {
		PostalCode       string
		CountryCode, Etc string
	} `json:"addrInfo"`
}
type CreateWalletResponse struct {
	ID string `json:"ID"`
}

var walletCollection *wallet.Collection

func validateCreateWalletRequest(req *CreateWalletRequest) error {
	dat, err := base64.StdEncoding.DecodeString(req.PublicKeyBase64)
	if err != nil {
		return fmt.Errorf("decode publicKeyBase64 failed,fieldName:%q,innerError:%w", err)
	}
	req.publicKey = dat

}

var CreateWallet http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	// decode requestBody and validate it
	var requestBody CreateWalletRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		replyWithError(err, w, http.StatusBadRequest)

		return
	}
	if err := validateCreateWalletRequest(&requestBody); err != nil {
		replyWithError(err, w, http.StatusBadRequest)

		return
	}
	// process
	id := walletCollection.BeginCreate(requestBody.publicKey)
	output := &CreateWalletResponse{
		ID: id,
	}

	// encode output
	if err := json.NewEncoder(w).Encode(output); err != nil {
		panic(err)
	}
}
var GetWallet http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {

}
