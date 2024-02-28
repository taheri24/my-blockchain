package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"taheri24.ir/blockchain/utils"
)

type takeOverRequestBody struct {
	Miners            []string         `json:"miners"`
	Tokens            []uint64         `json:"tokens"`
	CurrentSignatures map[string]int64 `json:"currentSignatures"`
	NextSignatures    map[string]int64 `json:"nextSignatures"`
}

type takeOverResponse struct {
	SuccessSignature string `json:"successSignature"`
}

var TokenTakeOver http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	setErr := utils.ErrorHandlerByRequest(r, w)
	var reqBody takeOverRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		setErr(err)
		return
	}
	tokens, miners := utils.ArrayToBoolMap(reqBody.Tokens), utils.ArrayToBoolMap(reqBody.Miners)
	valididateTokenKey := func(tokenKey string) bool {
		segments := strings.Split(tokenKey, ":")
		if len(segments) != 2 {
			setErr(fmt.Errorf(`invalidKey:%q ,error:%w`, tokenKey, errInvalidTokenByFormat))
			return false
		}
		tokenNumber, err := strconv.ParseUint(segments[0], 10, 64)
		if err != nil {
			setErr(fmt.Errorf(`invalidKey:%q, tokenNumber:%s ,ParseUintError:%w,error:%w`, tokenKey, segments[0], err, errInvalidTokenByFormat))
			return false

		}
		if !tokens[tokenNumber] {
			setErr(fmt.Errorf(`invalidKey:%q, tokenNumber:%s , error:%w`, tokenKey, tokenKey, errInvalidTokenBySignatureMismatch))
			return false

		}
		minerId := segments[1]
		if !miners[minerId] {
			setErr(fmt.Errorf(`invalidKey:%q, minerID:%s , error:%w`, tokenKey, minerId, errInvalidTokenBySignatureMismatch))
			return false

		}
		return true
	}
	allKeys := map[string]uint8{}
	for key := range reqBody.CurrentSignatures {
		if !valididateTokenKey(key) {
			return
		}
		allKeys[key] = 1
	}
	for key := range reqBody.NextSignatures {
		if !valididateTokenKey(key) {
			return
		}

		allKeys[key] = allKeys[key] | 2
	}

	for key, val := range allKeys {
		if val&0xff != 3 {
			setErr(fmt.Errorf("invalid key:%q, innerError:%w", key, errUnmatchedToken))
			return
		}
	}

}
