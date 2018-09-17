package cryptonight

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestHashVariant1(t *testing.T) {
	var input []byte
	var expected_hash []byte
	var actual_hash []byte

	// The following five test cases come from the Monero github
	// repository's test file, tests-slow-1.txt. They make sure we
	// correctly pulled in and compiled the cryptonight variant 1
	// implementation. The file is just some pairs of input and
	// expected output. For the curious, you can see how exactly
	// tests-slow-1.txt gets used by taking a look at Monero's
	// tests/hash/main.cpp file.
	input = hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	expected_hash = hexutil.MustDecode("0xb5a7f63abb94d07d1a6445c36c07c7e8327fe61b1647e391b4c7edae5de57a3d")
	actual_hash = HashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	expected_hash = hexutil.MustDecode("0x80563c40ed46575a9e44820d93ee095e2851aa22483fd67837118c6cd951ba61")
	actual_hash = HashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x8519e039172b0d70e5ca7b3383d6b3167315a422747b73f019cf9528f0fde341fd0f2a63030ba6450525cf6de31837669af6f1df8131faf50aaab8d3a7405589")
	expected_hash = hexutil.MustDecode("0x5bb40c5880cef2f739bdb6aaaf16161eaae55530e7b10d7ea996b751a299e949")
	actual_hash = HashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x37a636d7dafdf259b7287eddca2f58099e98619d2f99bdb8969d7b14498102cc065201c8be90bd777323f449848b215d2977c92c4c1c2da36ab46b2e389689ed97c18fec08cd3b03235c5e4c62a37ad88c7b67932495a71090e85dd4020a9300")
	expected_hash = hexutil.MustDecode("0x613e638505ba1fd05f428d5c9f8e08f8165614342dac419adc6a47dce257eb3e")
	actual_hash = HashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}

	input = hexutil.MustDecode("0x38274c97c45a172cfc97679870422e3a1ab0784960c60514d816271415c306ee3a3ed1a77e31f6a885c3cb")
	expected_hash = hexutil.MustDecode("0xed082e49dbd5bbe34a3726a0d1dad981146062b39d36d62c71eb1ed8ab49459b")
	actual_hash = HashVariant1(input)
	if !bytes.Equal(actual_hash, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(actual_hash), " versus ", hex.EncodeToString(expected_hash))
	}
}

func TestHashVariant1ForEthereum(t *testing.T) {
	block_header_bytes := []byte("deadbeefdeadbeefdeadbeefdeadbeef")
	var nonce uint64 = 123
	var digest []byte
	var result []byte
	digest, result = HashVariant1ForEthereumHeader(block_header_bytes, nonce)
	if !bytes.Equal(digest, result) {
		t.Error("Digest should match result: ", hex.EncodeToString(digest), " versus ", string(result))
	}
	expected_hash := hexutil.MustDecode("0x3c71949afc7c6e042d81f25c97f6da25b6ff6abcb039e170610adf94c409922d")
	if !bytes.Equal(result, expected_hash) {
		t.Error("Unexpected result: ", hex.EncodeToString(result), " versus ", hex.EncodeToString(expected_hash))
	}
}
