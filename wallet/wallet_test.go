package wallet

import (
	"testing"
)

func TestNew(t *testing.T) {
	type testCase struct {
		name     string
		wantID   ID
		visitFns []Option
	}
	tests := []testCase{
		{name: "emptyID", wantID: "", visitFns: []Option{WalletID("")}},
		{name: "ID_HelloWorld", wantID: "HelloWorld", visitFns: []Option{WalletID("HelloWorld")}},
		{name: "ID_HelloWorld", wantID: "HelloWorld", visitFns: []Option{WalletID("HelloWorld")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); got.ID != "" {
				t.Errorf("New() = %v, want %v", got, tt.wantID)
			}
		})
	}
}
