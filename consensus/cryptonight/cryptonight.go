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

// Some other possibly useful CFLAGS:
// -I. -Ofast -fuse-linker-plugin -funroll-loops -fvariable-expansion-in-unroller -ftree-loop-if-convert-stores -fmerge-all-constants -fbranch-target-load-optimize2 -fsched2-use-superblocks -falign-loops=16 -falign-functions=16 -falign-jumps=16 -falign-labels=16 -Wno-pointer-sign -Wno-pointer-to-int-cast -maes -march=native -Wl,--stack,10485760

// Direct wrapper around cryptonight's cn_slow_hash. You should
// probably not use this function, and instead use
// HashVariant1ForEthereumHeader below.
//
// Takes input hash material with length at least 43 bytes, and
// returns the 32 byte hash. The number 43 is due to an invariant in
// the C implementation of cn_slow_hash. If you pass less than 43
// bytes, the C code will halt with a runtime error and a reasonably
// descriptive message.
func HashVariant1(input []byte) []byte {
	result := make([]byte, 32)
	input_ptr := unsafe.Pointer(&input[0])
    output_ptr := unsafe.Pointer(&result[0])
	C.cn_slow_hash(input_ptr, C.size_t(len(input)), (*C.char)(output_ptr), 1 /*variant*/, 0 /*prehashed*/)
	return result
}

// Similar to hashimoto, HashVariant1ForEthereumHeader accepts 32 byte
// block header hash and 8 byte nonce, then returns 32 byte digest and
// 32 byte result.
func HashVariant1ForEthereumHeader(block_header_hash []byte, nonce uint64) ([]byte, []byte) {
	/*
	// Note: we need to use at least 43 bytes of hash material, so we
	// repeat the 8 bytes of nonce.
	header_concat_nonce := make([]byte, 48)
	copy(header_concat_nonce, block_header_hash)
	binary.LittleEndian.PutUint64(header_concat_nonce[32:], nonce)
	binary.LittleEndian.PutUint64(header_concat_nonce[40:], nonce)
	result := HashVariant1(header_concat_nonce)
    */
	blob := make([]byte, 76)
	for i := 0; i < len(blob); i++ {
		// Initialize to 0x77 for all bytes.
		blob[i] = 119
	}
	var blen int = 0
	// major version and minor version
	blob[blen] = 7;
	blen++
	blob[blen] = 0;
	blen++
	// 5 byte timestamp
	for i := 0; i < 5; i++ {
		// Initialize to 0x77 for all bytes.
		blob[blen] = 0
		blen++
	}
	copy(blob[blen:], block_header_hash)
	blen += 32
	var actual_nonce uint32 = uint32(nonce & 0x00000000ffffffff)
	binary.LittleEndian.PutUint32(blob[blen:], actual_nonce)
	
	result := HashVariant1(blob)
	
	// The digest is always equal to the result, since unlike
	// hashimoto, we don't have separate full and light versions of
	// cryptonight hashing.
	return result, result
}
