package builder

import (
	"fmt"
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"reflect"
)

type ProbeAccount struct {
	AccountValueEntryList []AccountFieldMeta
}

type ProbeAccount2 struct {
	Username *AccountField
	Password *AccountField
	Address  *AccountField
}

type AccountValuesConverter struct{}

func NewAccountValuesConverter2(account *ProbeAccount2) *AccountValuesConverter {
	t := reflect.TypeOf(account)
	// Check if the input is a pointer and dereference it if yes
	if t.Kind() == reflect.Ptr {
		fmt.Printf("Input param is a pointer\n")
		t = t.Elem()
	}
	// Check if the input is a struct
	if t.Kind() != reflect.Struct {
		fmt.Printf("Input param is not a struct\n")
	}
	fmt.Printf("****** Input Type: %v::%v\n", t, t.Name())
	fmt.Printf("Num of fields = %d\n", t.NumField())
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("%d %s = %v\n", i, f.Type, f.Name)

		s := reflect.TypeOf(f)
		// Check if the input is a struct
		if s.Kind() != reflect.Struct {
			fmt.Printf("\tInput param is not a struct\n")
		}
		fmt.Printf("\tInput Type: %v::%v\n", s, s.Name())

		value_of_s := reflect.ValueOf(f)
		fmt.Printf("\tValue type: %v\n", value_of_s.Kind())
		fmt.Printf("\tNum of fields in the value = %d\n", value_of_s.NumField())
	}

	return nil
}

func NewAccountValuesConverter(account *ProbeAccount) *AccountValuesConverter {
	for _, accountValueEntry := range account.AccountValueEntryList {
		t := reflect.TypeOf(accountValueEntry)
		// Check if the input is a pointer and dereference it if yes
		if t.Kind() == reflect.Ptr {
			//fmt.Printf("Input param is a pointer\n")
			t = t.Elem()
		}
		// Check if the input is a struct
		if t.Kind() != reflect.Struct {
			fmt.Printf("Input param is not a struct\n")
		}
		fmt.Printf("****** Input Type: %v::%v\n", t, t.Name())
		fmt.Printf("Num of fields = %d\n", t.NumField())
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Printf("%d %s = %v\n", i, f.Type, f.Name)
		}

		s := reflect.ValueOf(accountValueEntry)
		fmt.Printf("Value %v\n", s)

		//s := reflect.ValueOf(&t)
		//fmt.Println("type of p:", s.Type())
		//fmt.Printf("Value Type %v\n", s)
		//fmt.Printf("Value %v:\n", s.Elem())
		//
		//typeOfT := s.Elem().Kind()
		//if typeOfT != reflect.Struct {
		//	fmt.Printf("Input param is not a struct\n")
		//}
		//fmt.Printf("Value Type %v\n", typeOfT)

	}

	return nil
}

func (converter *AccountValuesConverter) getAccountDefinitions() []*proto.AccountDefEntry {
	return nil
}

func (converter *AccountValuesConverter) getTargetIdFields() []string {
	return nil
}

func (converter *AccountValuesConverter) createAccountValue(accountValues []*proto.AccountValue) {
	return
}
