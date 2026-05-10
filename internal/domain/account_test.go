package domain

import (
	"testing"

	"github.com/elyosemite/everest/internal/domain/account"
)

type myOwnAccount struct {
	name        string
	accountType account.AccountType
	balance     float64
	wantErr     bool
}

func TestNewAccount(t *testing.T) {
	tests := []myOwnAccount{
		{"checking with zero balance", account.Checking, 0, false},
		{"savings with positive balance", account.Savings, 500.0, false},
		{"negative balance returns error", account.Checking, -1, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a, err := account.NewAccount("user-1", tc.accountType, tc.balance)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("ünexpected error: %v", err)
			}
			if a.Type() != tc.accountType {
				t.Errorf("got type %v, want %v", a.Type(), tc.accountType)
			}
			if a.Balance() != tc.balance {
				t.Errorf("got balance %v, want %v", a.Balance(), tc.balance)
			}
		})
	}
}

func TestAccountType_String(t *testing.T) {
	tests := []struct {
		t    account.AccountType
		want string
	}{
		{account.Checking, "Conta Corrente"},
		{account.Savings, "Conta Poupança"},
		{account.AccountType(99), "Conta desconhecida"},
	}

	for _, tc := range tests {
		t.Run(tc.want, func(t *testing.T) {
			if got := tc.t.String(); got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
