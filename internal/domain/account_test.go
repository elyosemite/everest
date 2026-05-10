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

func TestSingleAccount(t *testing.T) {
	myAccount := myOwnAccount{
		name:        "fds",
		accountType: account.Savings,
		balance:     12,
		wantErr:     false,
	}

	t.Run(account.name, func(t *testing.T) {
		foo, err := account.NewAccount(myAccount.name)
		if account.wantErr {
			if err == nil {
				t.Fatalf("expected error, got nil")
				return
			}
		}
	})
}

func TestNewAccount(t *testing.T) {
	account_test_cases := []myOwnAccount{
		{"checking with zero balance", account.Checking, 0, false},
		{"savings with positive balance", account.Savings, 500.0, false},
		{"negative balance returns error", account.Checking, -1, true},
	}

	for _, testCase := range account_test_cases {
		t.Run(testCase.name, func(t *testing.T) {
			a, err := account.NewAccount("user-1", testCase.accountType, testCase.balance)
			if testCase.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("ünexpected error: %v", err)
			}
			if a.Type() != testCase.accountType {
				t.Errorf("got type %v, want %v", a.Type(), testCase.accountType)
			}
			if a.Balance() != testCase.balance {
				t.Errorf("got balance %v, want %v", a.Balance(), testCase.balance)
			}
		})
	}
}

type accountTypeStringCase struct {
	t    account.AccountType
	want string
}

func TestAccountType_String(t *testing.T) {
	tests := []accountTypeStringCase{
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

func TestNewAccount_NegativeBalanceReturnsError(t *testing.T) {
	_, err := account.NewAccount("user-1", account.Checking, -1)
	if err == nil {
		t.Fatalf("expected error for negative balance")
	}
}
