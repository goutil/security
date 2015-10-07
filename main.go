package security

import (
	"crypto/sha1"
	"fmt"
)

// SHA1 Computes the checksum of `s` using SHA-1 and returns the string version.
func SHA1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}
