package ident

import "testing"

func TestCompare(t *testing.T) {
	data := []struct {
		mine     string
		other    string
		expected int
	}{
		{"org1/redis/1234/123456", "org1/redis/1234/123456", 0},
		{"org2/redis/1.2.3/1234", "org1/redis/1.2.4/1234", -1},
		{"org2/redis/1.2.4/1234", "org1/redis/1.2.4-dev/1234", 1},
		{"org2/redis/1.2.4/1234", "org1/redis/1.2.3/1234", 1},
		{"org2/redis/1.2.4/1235", "org1/redis/1.2.4/1234", 1},
		{"org2/redis/1.2.4/1233", "org1/redis/1.2.4/1234", -1},
		{"org2/redis/1.2.4/1233", "org1/redis/1.2.4", -1},
		{"org2/redis/1.2.4/1233", "org1/redis/1.2.4", -1},
		{"org2/redis/1.2.4/1233", "org1/redis", -1},
		{"org2/redis", "org1/redis", 0},
	}

	for _, testData := range data {
		myIdent, err := FromString(testData.mine)
		if err != nil {
			t.Errorf("Error parsing Ident String: %s", err)
		}
		otherIdent, err := FromString(testData.other)
		if err != nil {
			t.Errorf("Error parsing Ident String: %s", err)
		}
		actual, err := myIdent.Compare(otherIdent)
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if actual != testData.expected {
			t.Errorf("Expected: %d, Got: %d", testData.expected, actual)
		}
	}
}
