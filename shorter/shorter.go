package shorter

import (
	"crypto/sha256"
	"hash/fnv"
	"math/rand"
)

const base63 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func getHash(original string) int64 {
	shaHash := sha256.Sum256([]byte(original))
	fnvHash := fnv.New64a()
	fnvHash.Write(shaHash[:])
	return int64(fnvHash.Sum64())
}
func RandString() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = base63[rand.Intn(len(base63))]
	}
	return string(b)
}

// GetShort
func GetShort(original string) string {
	hash := getHash(original)

	rand.Seed(hash)

	return RandString()
}
