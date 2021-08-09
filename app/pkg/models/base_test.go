package models

import (
	"testing"
)

func TestInit(t *testing.T) {
}

func TestEncrypt(t *testing.T) {
	want := "9d4e1e23bd5b727046a9e3b4b7db57bd8d6ee684"
	got := Encrypt("pass")
	if got != want {
		t.Errorf("want %s, but %s", want, got)
	}
}

func TestCreateUUID(t *testing.T) {
}
