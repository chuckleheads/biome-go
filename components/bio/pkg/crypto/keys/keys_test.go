package keys

import (
	"encoding/base64"
	"testing"
)

func TestKeyFromString(t *testing.T) {
	keyStr := `SIG-PUB-1
edavis-20180223211128
	
zgH4jGMKwebgZtztsh3nPqZqyE6+2ENMCuyNEbHV5uc=
`
	key := KeyFromString(keyStr)
	if key.Type != Sig {
		t.Errorf("Expected %d. Got %d", Sig, key.Type)
	}
	if key.Name != "edavis" {
		t.Errorf("Expected edavis. Got %s", key.Name)
	}
	if key.Rev != "20180223211128" {
		t.Errorf("Expected 20180223211128. Got %s", key.Rev)
	}
	if key.Secret != false {
		t.Errorf("Expected false. Got %v", key.Secret)
	}
	encodedKey := base64.StdEncoding.EncodeToString(key.Key)
	if encodedKey != "zgH4jGMKwebgZtztsh3nPqZqyE6+2ENMCuyNEbHV5uc=" {
		t.Errorf("Expected zgH4jGMKwebgZtztsh3nPqZqyE6+2ENMCuyNEbHV5uc=. Got %s", encodedKey)
	}
}
