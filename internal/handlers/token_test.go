package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"taheri24.ir/blockchain/utils"
)

func TestTakeOver(t *testing.T) {
	if err := evalError(TokenTakeOver, takeOverRequestBody{
		Miners: []string{"1"},
		Tokens: []uint64{1},
		CurrentSignatures: map[string]int64{
			"1:1:noise": 2,
		},
		NextSignatures: map[string]int64{
			"1:1": 400,
		},
	}); !errors.Is(err, errInvalidTokenByFormat) {
		t.Log(err)
		t.Fail()
	}
	if err := evalError(TokenTakeOver, takeOverRequestBody{
		Miners: []string{"1"},
		Tokens: []uint64{1, 10000},
		CurrentSignatures: map[string]int64{
			"1:1": 2,
		},
		NextSignatures: map[string]int64{
			"1:1":     400,
			"10000:1": 40,
		},
	}); !errors.Is(err, errUnmatchedToken) {
		t.Log(err)
		t.Fail()
	}
	if err := evalError(TokenTakeOver, takeOverRequestBody{
		Miners: []string{"1"},
		Tokens: []uint64{1, 10000},
		CurrentSignatures: map[string]int64{
			"1:1":     2,
			"10000:1": 240,
		},
		NextSignatures: map[string]int64{
			"1:1": 400,
		},
	}); !errors.Is(err, errUnmatchedToken) {
		t.Log(errUnmatchedToken)
		t.Fail()
	}

}
func eval[Tout any](fn http.HandlerFunc, req any) Tout {
	buf, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	var result Tout
	r, w := httptest.NewRequest("", "", bytes.NewReader(buf)), httptest.NewRecorder()
	fn(w, r)
	return result
}

func evalError(fn http.HandlerFunc, req any, possibleErrors ...error) error {
	buf, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}
	requestErr := errUnexpectedSuccess
	ctx := context.WithValue(context.Background(), utils.ErrorCaptureKey, func(err error) {
		requestErr = err
	})
	r, w := httptest.NewRequest("POST", "/", bytes.NewReader(buf)).WithContext(ctx), httptest.NewRecorder()
	fn(w, r)
	return requestErr
}
