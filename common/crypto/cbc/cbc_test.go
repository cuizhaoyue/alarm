package cbc_test

import (
	"encoding/hex"
	"testing"

	"ezone.xxxxx.com/xxxxx/xxxxx/alarm/common/crypto/cbc"
	"github.com/stretchr/testify/require"
)

func TestAESCBC(t *testing.T) {
	data := []byte("hello world")
	key := []byte("abcdefghabcdefgh")

	should := require.New(t)

	cipherData, err := cbc.Encrypt(data, key)
	should.NoError(err)
	t.Logf("cipher data: %s", cipherData)

	rawData, err := cbc.Decrypt(cipherData, key)
	should.NoError(err)
	t.Logf("raw data: %s", rawData)

	should.Equal(data, rawData)
}

func TestAESDecode(t *testing.T) {
	s := "3685e251367f231b0226300e46738055"
	encryData, _ := hex.DecodeString(s)

	key := []byte("abcdefghabcdefgh")

	should := require.New(t)

	rawData, err := cbc.Decrypt(encryData, key)
	should.NoError(err)
	t.Logf("raw data: %s", rawData)
}
