package common

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/nacl/box"
	"io"
)

const metaMaskNaclVersion = "x25519-xsalsa20-poly1305"

type metaMaskEncryptedMessageFormat struct {
	Version        string `json:"version"`
	Nonce          string `json:"nonce"`
	EphemPublicKey string `json:"ephemPublicKey"`
	Ciphertext     string `json:"ciphertext"`
}

func EncryptMetaMaskMessage(message, recipientPublicKey []byte) ([]byte, error) {
	senderPublicKey, senderPrivateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate nacl key pair")
	}
	var nonce [24]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(err)
	}
	var recipientPublicKey2 [32]byte
	copy(recipientPublicKey2[:], recipientPublicKey)
	encrypted := box.Seal(nonce[:], message, &nonce, &recipientPublicKey2, senderPrivateKey)
	jsonMessage, err := json.Marshal(metaMaskEncryptedMessageFormat{
		Version:        metaMaskNaclVersion,
		Nonce:          base64.StdEncoding.EncodeToString(nonce[:]),
		EphemPublicKey: base64.StdEncoding.EncodeToString(senderPublicKey[:]),
		Ciphertext:     base64.StdEncoding.EncodeToString(encrypted[24:]),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to serialize nacl ciphertext to json")
	}
	return jsonMessage, nil
}

func DecryptMetaMaskMessage(ciphertext, recipientPrivateKey []byte) ([]byte, error) {
	var messageBody metaMaskEncryptedMessageFormat
	if err := json.Unmarshal(ciphertext, &messageBody); err != nil {
		return nil, fmt.Errorf("not compatible ciphertext with nacl encryption")
	}
	rawCiphertext, err := base64.StdEncoding.DecodeString(messageBody.Ciphertext)
	if err != nil {
		return nil, fmt.Errorf("can't decode ciphertext from base64")
	}
	rawNonce, err := base64.StdEncoding.DecodeString(messageBody.Nonce)
	if err != nil {
		return nil, fmt.Errorf("can't decode nonce from base64")
	}
	var nonce24 [24]byte
	copy(nonce24[:], rawNonce)
	senderPublicKey, err := base64.StdEncoding.DecodeString(messageBody.EphemPublicKey)
	if err != nil {
		return nil, fmt.Errorf("can't decode nonce from base64")
	}
	var recipientPrivateKey32 [32]byte
	copy(recipientPrivateKey32[:], recipientPrivateKey)
	var senderPublicKey32 [32]byte
	copy(senderPublicKey32[:], senderPublicKey)
	decrypted, ok := box.Open(nil, rawCiphertext, &nonce24, &senderPublicKey32, &recipientPrivateKey32)
	if !ok {
		return nil, fmt.Errorf("failed to decrypt nacl message")
	}
	return decrypted, nil
}
