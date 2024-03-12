package enclave_encrypt_test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cloudflare/circl/hpke"
	"github.com/stretchr/testify/assert"

	"github.com/tkhq/go-sdk/pkg/enclave_encrypt"
)

const KemId hpke.KEM = hpke.KEM_P256_HKDF_SHA256

func randomSignature(key *ecdsa.PrivateKey) []byte {
	signature, err := enclave_encrypt.P256Sign(key, []byte("random"))
	if err != nil {
		panic(err)
	}
	return signature
}

func TestP256SignAndVerify(t *testing.T) {
	msg := []byte("test message")
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	sig, err := enclave_encrypt.P256Sign(privateKey, msg)
	assert.Nil(t, err)

	assert.Equal(
		t,
		enclave_encrypt.P256Verify(&privateKey.PublicKey, msg, sig),
		true,
	)
}

// Second half of notarizer public key fixture for integration fixtures
func getNotarizerPublic(t *testing.T) ecdsa.PublicKey {
	notarizerAuthPublicHex := "04d91ba02c67a962fd9369acb35e8d39de543dcedc774ed27701afb9159184df6e1d607c0a12e93e7ea9f6eff876779523c7afcf1cd49fb82c2b146a5207ff3212"
	notarizerPublicBytes, err := hex.DecodeString(notarizerAuthPublicHex)
	assert.Nil(t, err)

	// FIXME: `elliptic.Unmarshal` is deprecated, but scm does not know how to replace it.
	// nolint:staticcheck
	x, y := elliptic.Unmarshal(elliptic.P256(), notarizerPublicBytes)
	assert.NotNil(t, x)

	return ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
}

func TestP256Verify(t *testing.T) {
	// Values generated by rust code
	msgHex := "047432fea9cd18d99908b7b680d3ca951038371e21f005668836722a0ff486b02154c76ab53f8953ad42382cb54441c4b37cdc0a797609e36d7766578665eeda9a"
	sigHex := "304502206dd781e8bef50a336072fbc741aa78edd4f58ee42b4e96e859a11fd286c0cf7a022100f83e370559d8842c148cfb6e44cc6495938b1a00371313b09e483f0970c1a63a"

	msg, err := hex.DecodeString(msgHex)
	assert.Nil(t, err)

	sig, err := hex.DecodeString(sigHex)
	assert.Nil(t, err)

	notarizerPublic := getNotarizerPublic(t)
	assert.True(
		t,
		enclave_encrypt.P256Verify(&notarizerPublic, msg, sig),
	)
}

func TestHexSerializationWorks(t *testing.T) {
	enc := enclave_encrypt.Bytes([]byte{1, 42})
	ciph := enclave_encrypt.Bytes([]byte{2})
	original := enclave_encrypt.ClientSendMsg{
		EncappedPublic: &enc,
		Ciphertext:     &ciph,
	}

	encoded, err := json.Marshal(original)
	assert.Nil(t, err)

	var decoded enclave_encrypt.ClientSendMsg
	err = json.Unmarshal(encoded, &decoded)
	assert.Nil(t, err)

	assert.Equal(
		t,
		original.EncappedPublic,
		decoded.EncappedPublic,
	)
}

func TestClientToServerE2e(t *testing.T) {
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey)
	assert.Nil(t, err)

	serverTarget, err := server.PublishTarget()
	assert.Nil(t, err)
	serverRecv := server.IntoEnclaveServerRecv()

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)
	clientCiphertext, err := client.Encrypt([]byte("test message"), *serverTarget)
	assert.Nil(t, err)

	bytes, err := json.Marshal(clientCiphertext)
	assert.Nil(t, err)
	clientCipherTextUnmarshal := enclave_encrypt.ClientSendMsg{}

	err = json.Unmarshal(bytes, &clientCipherTextUnmarshal)
	assert.Nil(t, err)

	plaintext, err := serverRecv.Decrypt(
		clientCipherTextUnmarshal,
	)
	assert.Nil(t, err)

	assert.Equal(
		t,
		plaintext,
		[]byte("test message"),
	)
}

func TestClientToServerE2eExistingTargetKey(t *testing.T) {
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey)
	assert.Nil(t, err)

	serverTarget, err := server.PublishTarget()
	assert.Nil(t, err)
	serverRecv := server.IntoEnclaveServerRecv()

	_, targetPrivate, err := KemId.Scheme().GenerateKeyPair()
	assert.Nil(t, err)
	client, err := enclave_encrypt.NewEnclaveEncryptClientFromTargetKey(&authKey.PublicKey, targetPrivate)
	assert.Nil(t, err)
	clientCiphertext, err := client.Encrypt([]byte("test message"), *serverTarget)
	assert.Nil(t, err)

	bytes, err := json.Marshal(clientCiphertext)
	assert.Nil(t, err)
	clientCipherTextUnmarshal := enclave_encrypt.ClientSendMsg{}

	err = json.Unmarshal(bytes, &clientCipherTextUnmarshal)
	assert.Nil(t, err)

	plaintext, err := serverRecv.Decrypt(
		clientCipherTextUnmarshal,
	)
	assert.Nil(t, err)

	assert.Equal(
		t,
		plaintext,
		[]byte("test message"),
	)
}

func TestClientToServerRejectBadServerTargetSignature(t *testing.T) {
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey)
	assert.Nil(t, err)

	serverTarget, err := server.PublishTarget()
	sig := enclave_encrypt.Bytes(randomSignature(authKey))
	serverTarget.TargetPublicSignature = &sig
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	_, err = client.Encrypt([]byte("test message"), *serverTarget)
	assert.Equal(
		t,
		err.Error(),
		"invalid enclave auth key signature",
	)
}

func TestServerToClientE2e(t *testing.T) {
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey)
	assert.Nil(t, err)
	serverCiphertext, err := server.Encrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	bytes, err := json.Marshal(serverCiphertext)
	assert.Nil(t, err)
	var serverCiphertextUnmarshal enclave_encrypt.ServerSendMsg
	err = json.Unmarshal(bytes, &serverCiphertextUnmarshal)
	assert.Nil(t, err)

	plaintext, err := client.Decrypt(serverCiphertextUnmarshal)
	assert.Nil(t, err)
	assert.Equal(
		t,
		plaintext,
		[]byte("test message"),
	)
}

func TestServerToClientE2eExistingTargetKey(t *testing.T) {
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	_, targetPrivate, err := KemId.Scheme().GenerateKeyPair()
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServerFromTargetKey(authKey, &targetPrivate)
	assert.Nil(t, err)
	serverCiphertext, err := server.Encrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	bytes, err := json.Marshal(serverCiphertext)
	assert.Nil(t, err)
	var serverCiphertextUnmarshal enclave_encrypt.ServerSendMsg
	err = json.Unmarshal(bytes, &serverCiphertextUnmarshal)
	assert.Nil(t, err)

	plaintext, err := client.Decrypt(serverCiphertextUnmarshal)
	assert.Nil(t, err)
	assert.Equal(
		t,
		plaintext,
		[]byte("test message"),
	)
}

func TestServerToClientRejectBadEncappedPublicSignature(t *testing.T) {
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey)
	assert.Nil(t, err)
	serverCiphertext, err := server.Encrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	sig := enclave_encrypt.Bytes(randomSignature(authKey))
	serverCiphertext.EncappedPublicSignature = &sig

	_, err = client.Decrypt(*serverCiphertext)
	assert.Equal(
		t,
		err.Error(),
		"invalid enclave auth key signature",
	)
}

func TestEmailRecoveryServerToClientE2e(t *testing.T) {
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey)
	assert.Nil(t, err)
	serverCiphertext, err := server.AuthEncrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	plaintext, err := client.AuthDecrypt(serverCiphertext)
	assert.Nil(t, err)
	assert.Equal(
		t,
		plaintext,
		[]byte("test message"),
	)
}

func TestValidateChecksum(t *testing.T) {
	assert.Nil(t, enclave_encrypt.ValidateChecksum(base58.Decode("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")))
	assert.Equal(t,
		"checksum mismatch for payload 0062e907b15cbf27d5425399ebf6f0fb50ebb88f18c29b7d94: [194 155 125 147] (computed) != [194 155 125 148] (last four bytes)",
		enclave_encrypt.ValidateChecksum(base58.Decode("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNb")).Error(),
	)
	assert.Equal(t,
		"payload length is < 5 (length: 3)",
		enclave_encrypt.ValidateChecksum(base58.Decode("dang")).Error(),
	)
}
