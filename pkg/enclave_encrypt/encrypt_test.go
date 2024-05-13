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
		true,
		enclave_encrypt.P256Verify(&privateKey.PublicKey, msg, sig),
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
	quorumPubHex := "04d91ba02c67a962fd9369acb35e8d39de543dcedc774ed27701afb9159184df6e1d607c0a12e93e7ea9f6eff876779523c7afcf1cd49fb82c2b146a5207ff3212"

	msg, err := hex.DecodeString(msgHex)
	assert.Nil(t, err)

	sig, err := hex.DecodeString(sigHex)
	assert.Nil(t, err)

	quorumPubBytes, err := hex.DecodeString(quorumPubHex)
	assert.Nil(t, err)
	quorumPub, err := enclave_encrypt.ToEcdsaPublic(quorumPubBytes)
	assert.Nil(t, err)

	notarizerPublic := getNotarizerPublic(t)
	assert.Equal(
		t,
		&notarizerPublic,
		quorumPub,
	)

	assert.True(
		t,
		enclave_encrypt.P256Verify(quorumPub, msg, sig),
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
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	userId := "4eb08a82-00ee-4f29-b076-fca769209725"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, organizationId, &userId)
	assert.Nil(t, err)

	serverTarget, err := server.PublishTarget()
	assert.Nil(t, err)
	serverTargetBytes, err := json.Marshal(serverTarget)
	assert.Nil(t, err)

	serverRecv := server.IntoEnclaveServerRecv()

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)
	clientCiphertext, err := client.Encrypt([]byte("test message"), serverTargetBytes, organizationId, userId)
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
		[]byte("test message"),
		plaintext,
	)
}

func TestClientToServerE2eExistingTargetKey(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	userId := "4eb08a82-00ee-4f29-b076-fca769209725"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, organizationId, &userId)
	assert.Nil(t, err)

	serverTarget, err := server.PublishTarget()
	assert.Nil(t, err)
	serverTargetBytes, err := json.Marshal(serverTarget)
	assert.Nil(t, err)

	serverRecv := server.IntoEnclaveServerRecv()

	_, targetPrivate, err := KemId.Scheme().GenerateKeyPair()
	assert.Nil(t, err)
	client, err := enclave_encrypt.NewEnclaveEncryptClientFromTargetKey(&authKey.PublicKey, targetPrivate)
	assert.Nil(t, err)
	clientCiphertext, err := client.Encrypt([]byte("test message"), serverTargetBytes, organizationId, userId)
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
		[]byte("test message"),
		plaintext,
	)
}

func TestClientToServerRejectBadServerTargetSignature(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	userId := "4eb08a82-00ee-4f29-b076-fca769209725"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, organizationId, &userId)
	assert.Nil(t, err)

	serverTarget, err := server.PublishTarget()
	assert.Nil(t, err)

	sig := enclave_encrypt.Bytes(randomSignature(authKey))
	serverTarget.DataSignature = sig

	serverTargetBytes, err := json.Marshal(serverTarget)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)
	_, err = client.Encrypt([]byte("test message"), serverTargetBytes, organizationId, userId)
	assert.Equal(
		t,
		"invalid enclave auth key signature",
		err.Error(),
	)
}

func TestClientToServerRejectBadOrganizationId(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	userId := "4eb08a82-00ee-4f29-b076-fca769209725"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, "a9ee173c-e089-46a0-8cbc-c0d87a125b07", &userId)
	assert.Nil(t, err)

	serverTarget, err := server.PublishTarget()
	assert.Nil(t, err)
	serverTargetBytes, err := json.Marshal(serverTarget)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)
	_, err = client.Encrypt([]byte("test message"), serverTargetBytes, organizationId, userId)
	assert.Equal(
		t,
		"organization id does not match expected value",
		err.Error(),
	)
}

func TestClientToServerRejectMissingUserId(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	userId := "4eb08a82-00ee-4f29-b076-fca769209725"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, organizationId, &userId)
	assert.Nil(t, err)

	serverTarget, err := server.PublishTarget()
	assert.Nil(t, err)
	serverTargetBytes, err := json.Marshal(serverTarget)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	userId = "fccd371a-a417-447e-89d0-5af8501cd8d9"
	_, err = client.Encrypt([]byte("test message"), serverTargetBytes, organizationId, userId)
	assert.Equal(
		t,
		"user id does not match expected value",
		err.Error(),
	)
}

func TestServerToClientE2e(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, organizationId, nil)
	assert.Nil(t, err)
	serverCiphertext, err := server.Encrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	bytes, err := json.Marshal(serverCiphertext)
	assert.Nil(t, err)
	plaintext, err := client.Decrypt(bytes, organizationId)
	assert.Nil(t, err)
	assert.Equal(
		t,
		[]byte("test message"),
		plaintext,
	)
}

func TestServerToClientE2eExistingTargetKey(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	_, targetPrivate, err := KemId.Scheme().GenerateKeyPair()
	assert.Nil(t, err)
	server, err := enclave_encrypt.NewEnclaveEncryptServerFromTargetKey(authKey, &targetPrivate, organizationId, nil)
	assert.Nil(t, err)
	serverCiphertext, err := server.Encrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	bytes, err := json.Marshal(serverCiphertext)
	assert.Nil(t, err)
	plaintext, err := client.Decrypt(bytes, organizationId)
	assert.Nil(t, err)
	assert.Equal(
		t,
		[]byte("test message"),
		plaintext,
	)
}

func TestServerToClientRejectBadEncappedPublicSignature(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, organizationId, nil)
	assert.Nil(t, err)
	serverCiphertext, err := server.Encrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	sig := enclave_encrypt.Bytes(randomSignature(authKey))
	serverCiphertext.DataSignature = sig

	serverSendtBytes, err := json.Marshal(serverCiphertext)
	assert.Nil(t, err)
	_, err = client.Decrypt(serverSendtBytes, organizationId)
	assert.Equal(
		t,
		"invalid enclave auth key signature",
		err.Error(),
	)
}

func TestServerToClientRejectBadOrganizationId(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, "a9ee173c-e089-46a0-8cbc-c0d87a125b07", nil)
	assert.Nil(t, err)
	serverCiphertext, err := server.Encrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	bytes, err := json.Marshal(serverCiphertext)
	assert.Nil(t, err)
	_, err = client.Decrypt(bytes, organizationId)
	assert.Equal(
		t,
		"organization id does not match expected value",
		err.Error(),
	)
}

func TestEmailRecoveryServerToClientE2e(t *testing.T) {
	organizationId := "f412ea93-998b-45a5-9df8-d2797c7f1a67"
	authKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.Nil(t, err)

	client, err := enclave_encrypt.NewEnclaveEncryptClient(&authKey.PublicKey)
	assert.Nil(t, err)

	clientTarget, err := client.TargetPublic()
	assert.Nil(t, err)

	server, err := enclave_encrypt.NewEnclaveEncryptServer(authKey, organizationId, nil)
	assert.Nil(t, err)
	serverCiphertext, err := server.AuthEncrypt(clientTarget, []byte("test message"))
	assert.Nil(t, err)

	plaintext, err := client.AuthDecrypt(serverCiphertext)
	assert.Nil(t, err)
	assert.Equal(
		t,
		[]byte("test message"),
		plaintext,
	)
}

// This test is here to prove that bundles are deterministic: given the same bundle and the same receiver private key,
// decrypting a bundle should always yield the same result.
// It is also a nice demonstration of the internals of our flow. Below we do not go through the abstractions we've built.
// This makes this test a good guide for anyone who wants to implement email auth bundle decryption in other languages.
func TestDeterministicDecryptEmailBundle(t *testing.T) {
	// Value grabbed from running "TestEmailRecoveryServerToClientE2e" with:
	//    _, targetPrivate, err := KemId.Scheme().GenerateKeyPair()
	//    b, _ := targetPrivate.MarshalBinary()
	//    fmt.Printf("generated key: %02x\n", b)
	privateKeyBytes, err := hex.DecodeString("6727b16dbe8e137c5755b839b2b2a96adae1187eaf7f08bc9938afe5ed8a8a6e")
	assert.Nil(t, err)
	receiverPrivate, err := KemId.Scheme().UnmarshalBinaryPrivateKey(privateKeyBytes)
	assert.Nil(t, err)

	suite := hpke.NewSuite(KemId, enclave_encrypt.KdfId, enclave_encrypt.AeadId)
	receiver, err := suite.NewReceiver(receiverPrivate, []byte(enclave_encrypt.TurnkeyHpkeInfo))
	assert.Nil(t, err)

	// Value captured from "TestEmailRecoveryServerToClientE2e"
	payload := "BrcAL1P9Hb77ojs4MUG8svAnrHG717xwdoSnFWsTpA5XbbuoHPUtaZuRrUW8ZQGo11Z7Y5orfTUb7rr2mELPP8y5"
	payloadBytes := base58.Decode(payload)
	assert.Nil(t, enclave_encrypt.ValidateChecksum(payloadBytes))

	// Trim the checksum now that we've validated it
	payloadBytes = payloadBytes[:len(payloadBytes)-4]

	// Get the sender public key from the payload
	compressedKey := payloadBytes[0:33]
	ciphertext := payloadBytes[33:]
	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), compressedKey)
	// nolint:staticcheck
	encappedPublic := elliptic.Marshal(elliptic.P256(), x, y)

	opener, err := receiver.Setup(encappedPublic)
	assert.Nil(t, err)

	receiverPublicBytes, err := receiverPrivate.Public().MarshalBinary()
	assert.Nil(t, err)

	aad := []byte{}
	aad = append(aad, encappedPublic...)
	aad = append(aad, receiverPublicBytes...)

	plaintext, err := opener.Open(ciphertext, aad)
	assert.Nil(t, err)
	assert.Equal(
		t,
		[]byte("test message"),
		plaintext,
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
