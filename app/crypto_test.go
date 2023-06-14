package app

import (
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	plaintext := "a0a4f5f469624bfdade3f71410d0a6a5"
	encrypted, err := Crypto().Encrypt(plaintext)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	decrypted, err := Crypto().Decrypt(encrypted)
	if err != nil {
		t.Errorf("Error occurred [%v]", err)
	}
	if decrypted != plaintext {
		t.Errorf("Expected decrypted [%v], got [%v]", plaintext, decrypted)
	}
}
