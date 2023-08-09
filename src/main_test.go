package main

import (
	"fmt"
	"savemyrpg/dal"
	"testing"
)

func TestInit(t *testing.T) {
	if Init() != true {
		t.Errorf("Output Failed to connect")
	}
}

func TestRegister(t *testing.T) {
	Init()
	_, err := Register("Carl", "carl@gmail.com")
	dal.RemoveUser("carl@gmail.com")
	if err != nil {
		fmt.Println(err)
		t.Errorf("Failed to Register User")
	}
}
