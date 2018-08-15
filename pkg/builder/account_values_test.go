package builder

import (
	"fmt"
	"testing"
)

func TestAccount(t *testing.T) {

	var account *ProbeAccount
	account = &ProbeAccount{
		AccountValueEntryList: []AccountFieldMeta{},
	}
	acctFields := []*AccountField{
		{displayName: "User", targetId: false, description: "User name", secret: false},
		{displayName: "Password", targetId: false, description: "Password", secret: true},
		{displayName: "Address", targetId: true, description: "IP Address", secret: false},
	}
	for _, acctField := range acctFields {
		fmt.Printf("acctField: %++v\n", acctField)
		account.AccountValueEntryList = append(account.AccountValueEntryList, acctField)
	}
	NewAccountValuesConverter(account)
}

func TestAccount2(t *testing.T) {

	var account *ProbeAccount2
	account = &ProbeAccount2{}
	account.Username = &AccountField{displayName: "User", targetId: false, description: "User name", secret: false}
	account.Password = &AccountField{displayName: "Password", targetId: false, description: "Password", secret: true}
	account.Address = &AccountField{displayName: "Address", targetId: true, description: "IP Address", secret: false}
	NewAccountValuesConverter2(account)
}
