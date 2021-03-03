package hashing

import (
	"golang.org/x/crypto/argon2"
)

// don't change, or we have to rehash things! 😂
// TODO: bring from env/encrypt with KMS instead of hardcoding (and then rehash lol)
var staticSalt = []byte("7m34KH7OxahHlGmPrwzs5DJnDaTfX0he")

// Hash hashes lookup keys in db
// must use static salt in order to make this work, but is as strong as can be without having to do the lookups in memory
func Hash(cleartext []byte) []byte {
	return argon2.IDKey(cleartext, staticSalt, 3, 32*1024, 1, 32)
}
