package smrpg_test

import (
	"fmt"
	"savemyrpg/dal"
	"savemyrpg/smrpg"
	"testing"
)

func TestInit(t *testing.T) {
	if smrpg.Init() != true {
		t.Errorf("Output Failed to connect")
	}
}

func TestRegister(t *testing.T) {
	smrpg.Init()
	_, err := smrpg.Register("Carl", "carl@gmail.com")
	dal.RemoveUser("carl@gmail.com")
	if err != nil {
		fmt.Println(err)
		t.Errorf("Failed to Register User")
	}
}
