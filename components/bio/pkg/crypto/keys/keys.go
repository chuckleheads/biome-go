package keys

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/biome-sh/biome-go/components/bio/pkg/crypto"
	"github.com/biome-sh/biome-go/components/bio/pkg/fs"
)

type KeyType int

const (
	Sig KeyType = iota
	Box
	Sym
)

const cacheKeyPathEnvVar = "CACHE_KEY_PATH"

type Key struct {
	Name   string
	Rev    string
	Type   KeyType
	Key    []byte
	Secret bool
}

type KeyPair struct {
	PublicKey  *Key
	PrivateKey *Key
}

func New(public *Key, private *Key) KeyPair {
	kp := KeyPair{}
	if public != nil {
		kp.PublicKey = public
	}
	if private != nil {
		kp.PrivateKey = private
	}
	return kp
}

func NewFromBytes(name string, rev string, public []byte, private []byte, kt KeyType) KeyPair {
	kp := KeyPair{}
	if len(public) > 0 {
		kp.PublicKey = &Key{
			Name:   name,
			Rev:    rev,
			Type:   kt,
			Secret: false,
			Key:    public,
		}
	}

	if len(private) > 0 {
		kp.PrivateKey = &Key{
			Name:   name,
			Rev:    rev,
			Type:   kt,
			Secret: true,
			Key:    private,
		}
	}

	return kp
}

func NewFromString(public string, private string) KeyPair {
	kp := KeyPair{}
	if public != "" {
		key := KeyFromString(public)
		kp.PublicKey = &key
	}
	if private != "" {
		key := KeyFromString(private)
		kp.PrivateKey = &key
	}
	return kp
}

func (k *Key) nameWithRev() string {
	return fmt.Sprintf("%s-%s", k.Name, k.Rev)
}

func KeyFromString(keystr string) Key {
	scanner := bufio.NewScanner(strings.NewReader(keystr))
	key := Key{}
	// Key Type
	if !scanner.Scan() {
		panic(fmt.Sprintf("Malformed key string:\n(%s)", keystr))
	}

	if scanner.Text() == crypto.PublicSigKeyVersion || scanner.Text() == crypto.PublicBoxKeyVersion {
		key.Secret = false
	} else if scanner.Text() == crypto.PrivateSigKeyVersion ||
		scanner.Text() == crypto.PrivateSymKeyVersion ||
		scanner.Text() == crypto.PrivateSigKeyVersion ||
		scanner.Text() == crypto.PrivateBoxKeyVersion {
		key.Secret = true
	} else {
		panic(fmt.Sprintf("Unsupported Key Version: %s", scanner.Text()))
	}

	// Name and Rev
	if scanner.Scan() {
		nameAndRev := strings.Split(scanner.Text(), "-")
		key.Name, key.Rev = nameAndRev[0], nameAndRev[1]
	} else {
		panic(fmt.Sprintf("Malformed key string:\n(%s)", keystr))
	}

	// Key Content
	scanner.Scan()
	if !scanner.Scan() {
		panic(fmt.Sprintf("Malformed key string:\n(%s)", keystr))
	}

	keyBytes, err := base64.StdEncoding.DecodeString(scanner.Text())
	if err != nil {
		panic(err.Error())
	}

	key.Key = keyBytes
	return key
}

// DefaultCachePeyPath returns the default cache key path or the environment variable override value
func DefaultCacheKeyPath() string {
	val, present := os.LookupEnv(cacheKeyPathEnvVar)
	if present {
		return val
	}
	return fs.GetCacheKeyPath()
}
