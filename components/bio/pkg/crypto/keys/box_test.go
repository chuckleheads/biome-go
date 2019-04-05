package keys

import (
	"testing"
)

func TestEncryptWithNoReciever(t *testing.T) {
	testStr := "test string"
	kp := GenerateBoxKeyPair("test", "1234")
	enc, err := kp.Encrypt([]byte(testStr), nil)
	if err != nil {
		t.Fatal("Failed to Encrypt")
	}
	bs, err := SecretMetadata(enc)
	if err != nil {
		t.Fatal("Failed to get metadata")
	}
	dec, err := kp.Decrypt(bs.Ciphertext, nil, nil)
	if err != nil {
		t.Fatal("Failed to Decrypt")
	}
	if string(dec) != testStr {
		t.Fatalf("Expected %s. Got %s", testStr, string(dec))
	}
}

func TestEncryptWithReciever(t *testing.T) {
	testStr := "test string"
	kp := GenerateBoxKeyPair("test", "1234")
	recKp := GenerateBoxKeyPair("recv", "123456")
	enc, err := kp.Encrypt([]byte(testStr), &recKp)
	if err != nil {
		t.Fatal("Failed to Encrypt")
	}
	bs, err := SecretMetadata(enc)
	if err != nil {
		t.Fatal("Failed to get metadata")
	}
	dec, err := recKp.Decrypt(bs.Ciphertext, &bs.Nonce, &kp)
	if err != nil {
		t.Fatal("Failed to Decrypt")
	}
	if string(dec) != testStr {
		t.Fatalf("Expected %s. Got %s", testStr, string(dec))
	}
}
