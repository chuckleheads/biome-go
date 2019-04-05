package keys

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/GoKillers/libsodium-go/cryptobox"
	"github.com/GoKillers/libsodium-go/randombytes"
	"github.com/biome-sh/biome-go/components/bio/pkg/crypto"
)

type BoxSecret struct {
	Sender     string
	Ciphertext []byte
	Receiver   string
	Nonce      []byte
}

func GenerateBoxKeyPair(name string, rev string) KeyPair {
	sk, pk, _ := cryptobox.CryptoBoxKeyPair()

	return NewFromBytes(name, rev, pk, sk, Box)
}

func (kp *KeyPair) Encrypt(data []byte, reciever *KeyPair) (string, error) {
	if reciever == nil {
		return kp.encryptAnonymousBox(data)
	}
	return kp.encryptBox(data, *reciever)
}

func (kp *KeyPair) Decrypt(ciphertext []byte, nonce *[]byte, reciever *KeyPair) ([]byte, error) {
	if reciever == nil {
		return kp.decryptAnonymousBox(ciphertext)
	}
	return kp.decryptBox(ciphertext, *reciever, *nonce)
}
func SecretMetadata(payload string) (BoxSecret, error) {
	boxSecret := BoxSecret{}
	scanner := bufio.NewScanner(strings.NewReader(payload))
	var version string
	if scanner.Scan() {
		version = scanner.Text()
	} else {
		return boxSecret, fmt.Errorf("Malformed metadata: %s", payload)
	}

	if scanner.Scan() {
		boxSecret.Sender = scanner.Text()
	} else {
		return boxSecret, fmt.Errorf("Malformed metadata: %s", payload)
	}

	if isAnonymousBox(version) {
		boxSecret.Receiver = ""
	} else {
		if scanner.Scan() {
			boxSecret.Receiver = scanner.Text()
		} else {
			return boxSecret, fmt.Errorf("Malformed metadata: %s", payload)
		}
	}

	if isAnonymousBox(version) {
		boxSecret.Nonce = []byte{}
	} else {
		if scanner.Scan() {
			dec, err := base64.StdEncoding.DecodeString(scanner.Text())
			if err != nil {
				return boxSecret, fmt.Errorf("Malformed metadata: %s", payload)
			}
			boxSecret.Nonce = dec
		} else {
			return boxSecret, fmt.Errorf("Malformed metadata: %s", payload)
		}
	}

	if scanner.Scan() {
		dec, err := base64.StdEncoding.DecodeString(scanner.Text())
		if err != nil {
			return boxSecret, fmt.Errorf("Malformed metadata: %s", payload)
		}
		boxSecret.Ciphertext = dec
	} else {
		return boxSecret, fmt.Errorf("Malformed metadata: %s", payload)
	}
	return boxSecret, nil
}

func (kp *KeyPair) encryptBox(data []byte, reciever KeyPair) (string, error) {
	if reciever.PublicKey == nil {
		return "", fmt.Errorf("Public reciever key is required but not present")
	}
	if kp.PrivateKey == nil {
		return "", fmt.Errorf("Private sender key is required but not present")
	}
	nonce := randombytes.RandomBytes(cryptobox.CryptoBoxNonceBytes())
	ciphertext, _ := cryptobox.CryptoBoxEasy(data, nonce, reciever.PublicKey.Key, kp.PrivateKey.Key)

	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s",
		crypto.BoxFormatVersion,
		kp.PublicKey.nameWithRev(),
		reciever.PrivateKey.nameWithRev(),
		base64.StdEncoding.EncodeToString(nonce),
		base64.StdEncoding.EncodeToString(ciphertext),
	), nil

}

func (kp *KeyPair) decryptBox(data []byte, reciever KeyPair, nonce []byte) ([]byte, error) {
	if reciever.PrivateKey == nil {
		return []byte{}, fmt.Errorf("Private reciever key is required but not present")
	}
	if kp.PublicKey == nil {
		return []byte{}, fmt.Errorf("Public sender key is required but not present")
	}
	secret, exit := cryptobox.CryptoBoxOpenEasy(data, nonce, kp.PublicKey.Key, reciever.PrivateKey.Key)
	fmt.Println("Exit: ", exit)
	return secret, nil
}

func (kp *KeyPair) encryptAnonymousBox(data []byte) (string, error) {
	if kp.PublicKey == nil {
		return "", fmt.Errorf("Public key is required but not present")
	}

	ciphertext, _ := cryptobox.CryptoBoxSeal(data, kp.PublicKey.Key)
	return fmt.Sprintf("%s\n%s\n%s",
		crypto.AnonymousBoxFormatVersion,
		kp.PublicKey.nameWithRev(),
		base64.StdEncoding.EncodeToString(ciphertext),
	), nil
}

func (kp *KeyPair) decryptAnonymousBox(ciphertext []byte) ([]byte, error) {
	if kp.PrivateKey == nil || kp.PublicKey == nil {
		return []byte{}, fmt.Errorf("Public and Private keys are required but not present")
	}
	res, _ := cryptobox.CryptoBoxSealOpen(ciphertext, kp.PublicKey.Key, kp.PrivateKey.Key)
	return res, nil
}

func isAnonymousBox(version string) bool {
	return version == crypto.AnonymousBoxFormatVersion
}
