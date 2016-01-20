package security

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	mathrand "math/rand"
)

// SHA1 Computes the checksum of `s` using SHA-1 and returns the string version.
func SHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

// Returns a 63-bit random (or pseudo random if there's not enough entropy).
func RandInt63() int64 {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err != nil {
		return mathrand.Int63()
	}

	var n uint64 = 0
	shift := uint(0)
	for i := 0; i < len(b); i++ {
		n |= uint64(b[i]) << shift
		shift += 8
	}

	r := int64(n & 0x7fffffffffffffff)
	mathrand.Seed(r)
	return r
}
