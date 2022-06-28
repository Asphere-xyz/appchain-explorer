package common

import (
	"crypto/rand"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/nacl/box"
	"testing"
)

func TestMetaMaskMessageEncryption(t *testing.T) {
	metaMaskPublicKey, metaMaskPrivateKey, err := box.GenerateKey(rand.Reader)
	require.NoError(t, err)
	originalMessage := []byte("Alas, poor Yorick! I knew him, Horatio")
	cipherText, err := EncryptMetaMaskMessage(originalMessage, metaMaskPublicKey[:])
	require.NoError(t, err)
	decryptedMessage, err := DecryptMetaMaskMessage(cipherText, metaMaskPrivateKey[:])
	require.NoError(t, err)
	require.Equal(t, decryptedMessage, originalMessage)
}
