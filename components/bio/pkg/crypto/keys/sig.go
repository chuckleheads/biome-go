package keys

import "github.com/GoKillers/libsodium-go/cryptosign"

func GenerateSigKeyPair(name string, rev string) KeyPair {
	sk, pk, _ := cryptosign.CryptoSignKeyPair()

	return NewFromBytes(name, rev, pk, sk, Sig)
}
