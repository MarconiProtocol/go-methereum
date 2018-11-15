package cryptonight

/*
#cgo CFLAGS: -maes
#cgo LDFLAGS:
#include "hash-ops.h"
*/
import "C"
import (
	"encoding/binary"
	"unsafe"
)

// Setting the following environment variable before running 'make
// geth' was required to avoid an aes-related compile error:
// CGO_CFLAGS_ALLOW=-maes
//
// Some other possibly useful CFLAGS:
// -I. -Ofast -fuse-linker-plugin -funroll-loops -fvariable-expansion-in-unroller -ftree-loop-if-convert-stores -fmerge-all-constants -fbranch-target-load-optimize2 -fsched2-use-superblocks -falign-loops=16 -falign-functions=16 -falign-jumps=16 -falign-labels=16 -Wno-pointer-sign -Wno-pointer-to-int-cast -maes -march=native -Wl,--stack,10485760

// Direct wrapper around cryptonight's cn_slow_hash. You should
// probably not use this function, and instead use
// HashVariant{1,2}ForEthereumHeader below.
//
// Takes input hash material with length at least 43 bytes, and
// returns the 32 byte hash. The number 43 is due to an invariant in
// the C implementation of cn_slow_hash. If you pass less than 43
// bytes, the C code will halt with a runtime error and a reasonably
// descriptive message.
//
// Variant must be 1 or 2.
func hashCryptonight(input []byte, variant int) []byte {
	result := make([]byte, 32)
	input_ptr := unsafe.Pointer(&input[0])
    output_ptr := unsafe.Pointer(&result[0])
	C.cn_slow_hash(input_ptr, C.size_t(len(input)), (*C.char)(output_ptr), (C.int)(variant), 0 /*prehashed*/)
	return result
}

// Uses the above hash function in a specific way for given ethereum
// header and ethereum nonce (which is 8 bytes, vs more typical 4
// bytes in other protocols).
func hashCryptonightForEthereumHeader(block_header_hash []byte, nonce uint64, variant int) ([]byte, []byte) {
	// Note: this blob format intentionally looks hacky. We're trying
	// to match the length and some of the byte offsets that monero
	// uses, e.g. its major/minor versions and nonce, so that existing
	// monero-like mining software implementations (both cpu and gpu)
	// remain compatible with fewer changes.
	blob := make([]byte, 76)
	for i := 0; i < len(blob); i++ {
		// Initialize to 0x77 for all bytes.
		blob[i] = 119
	}
	var blen int = 0
	// major version
  if variant == 1 {
    blob[blen] = 7;
  } else {
    blob[blen] = 8;
  }
	blen++
  // and minor version
	blob[blen] = 0;
	blen++
	// 5 byte timestamp
	for i := 0; i < 5; i++ {
		// Initialize to zeroes for timestamp.
		blob[blen] = 0
		blen++
	}
	copy(blob[blen:], block_header_hash)
	blen += 32
	binary.LittleEndian.PutUint64(blob[blen:], nonce)
	result := hashCryptonight(blob, variant)

	// The digest is always equal to the result, since unlike
	// hashimoto, we don't have separate full and light versions of
	// cryptonight hashing.
	return result, result
}

// Similar to hashimoto, HashVariant{1,2}ForEthereumHeader accepts 32
// byte block header hash and 8 byte nonce, then returns 32 byte
// digest and 32 byte result.
func HashVariant1ForEthereumHeader(block_header_hash []byte, nonce uint64) ([]byte, []byte) {
	return hashCryptonightForEthereumHeader(block_header_hash, nonce, 1)
}

func HashVariant2ForEthereumHeader(block_header_hash []byte, nonce uint64) ([]byte, []byte) {
	return hashCryptonightForEthereumHeader(block_header_hash, nonce, 2)
}
